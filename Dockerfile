FROM golang:1.10-alpine3.8 AS compiler

WORKDIR /go/src/github.com/DouwaIO/hairtail
COPY . ./

RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    export GOARCH=amd64 && \
    cd src/cmd && \
    go vet && \
    go build -o /htail && \
    chmod +x /htail

FROM alpine:3.8

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --update bash

WORKDIR /
COPY --from=compiler /htail /bin/

# Metadata params
ARG VERSION
ARG BUILD_DATE
ARG VCS_URL
ARG VCS_REF
ARG NAME
ARG VENDOR

# Metadata
LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name=$NAME \
      org.label-schema.description="kubectl plugin for CodeRun" \
      org.label-schema.url="https://coderun.top" \
      org.label-schema.vcs-url=https://gitlab.com/douwa/dougo-plugins/$VCS_URL \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vendor=$VENDOR \
      org.label-schema.version=$VERSION \
      org.label-schema.docker.schema-version="1.0" \
      org.label-schema.docker.cmd="docker run -d crun/kubectl"

EXPOSE 8080
CMD ["htail"]
