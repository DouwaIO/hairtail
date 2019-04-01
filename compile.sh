#!/bin/sh
# set -x

echo "compile server starting"
MOUNTPATH="/go/src/gitlab.com/douwa/registry"
SRCPATH="${MOUNTPATH}/dougo/src/dougo"
docker run --rm -v ${PWD}:${MOUNTPATH}/dougo -w ${MOUNTPATH} golang:1.9.7 sh -c "cd ${SRCPATH}/cmd/drone-server; go build -o ${SRCPATH}/dist/drone-server"
echo "compile server ending"
echo

echo "compile agent starting"
MOUNTPATH="/go/src/gitlab.com/douwa/registry"
SRCPATH="${MOUNTPATH}/dougo/src/dougo"
docker run --rm -v ${PWD}:${MOUNTPATH}/dougo -w ${MOUNTPATH} golang:1.9.7 sh -c "cd ${SRCPATH}/cmd/agent; go build -o ${SRCPATH}/dist/dougo-agent"
echo "compile agent ending"
echo

echo "compile public agent starting"
MOUNTPATH="/go/src/gitlab.com/douwa/registry"
SRCPATH="${MOUNTPATH}/dougo/src/dougo"
docker run --rm -v ${PWD}:${MOUNTPATH}/dougo -w ${MOUNTPATH} golang:1.9.7 sh -c "cd ${SRCPATH}/cmd/agent; go build -ldflags \"-X main.agentType=public\" -o ${SRCPATH}/dist/dougo-agent-public"
echo "compile public agent ending"
echo
