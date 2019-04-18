#!/bin/sh
# set -x


echo "compile starting"
APPPATH="/go/src/github.com/DouwaIO/hairtail"
docker run --rm -v ${PWD}:${APPPATH} -w ${APPPATH} golang:1.9.7 sh -c "cd ${APPPATH}/src/cmd; go build -o ${APPPATH}/dist/htail"
echo "compile ending"
echo
