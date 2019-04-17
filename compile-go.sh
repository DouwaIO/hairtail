#!/bin/sh
# set -x

# env
CURPATH=$(pwd)

# create go src link
GOSRC=$GOPATH/src/github.com/DouwaIO
mkdir -p $GOSRC
# ln -s $CURPATH $GOSRC > /dev/null 2>&1
ln -s $CURPATH $GOSRC

SRCPATH="${GOSRC}/hairtail/src"

echo "go path: ${GOPATH}"
echo "current path: ${CURPATH}"
echo "src path: ${SRCPATH}"
echo

CGO_ENABLED=0 GOOS=linux GOARCH=amd64

echo "compile server starting"
cd $SRCPATH/cmd
go build -o $SRCPATH/dist/hairtail
echo "compile server ending"
echo
