#!/bin/bash

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$OLDGOPATH:$CURDIR"

LogPrefix=">>>>"
TIME="`date +"%H:%M:%S"`"

# 打包前检测Bindata是否开启

echo -e "$LogPrefix ${TIME} \033[42;37m start \033[0m"

echo "$LogPrefix ${TIME} assets bindata"
go-bindata -ignore=\\.DS_Store -pkg="assets" -o src/assets/assets.go assets/...

echo "$LogPrefix ${TIME} templates bindata"
go-bindata -ignore=\\.DS_Store -pkg="templates" -o src/templates/templates.go templates/...

echo "$LogPrefix ${TIME} src package"
gofmt -w src/

echo "$LogPrefix ${TIME} project bin"
go install gin_sample

echo -e "$LogPrefix ${TIME} \033[42;37m finished \033[0m"
