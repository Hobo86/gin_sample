#!/bin/bash

OS="local"
while getopts "lh" arg
do
    case $arg in
        l)
            OS="linux"
            ;;
        h)
            echo "-l build linux bin, default local"
            echo "-h [a] help"
            exit
            ;;
        ?)
            echo "unkonw argument"
            echo "-l build linux bin, default local"
            echo "-h help"
            exit 1
            ;;
    esac
done

# Go Path
CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$OLDGOPATH:$CURDIR"

LogPrefix=">>>>"

# 打包前检测Bindata是否开启

echo -e "$LogPrefix `date +"%H:%M:%S"` \033[42;37m start \033[0m"

echo "$LogPrefix `date +"%H:%M:%S"` assets bindata"
go-bindata -ignore=\\.DS_Store -pkg="assets" -o src/assets/assets.go assets/...

echo "$LogPrefix `date +"%H:%M:%S"` templates bindata"
go-bindata -ignore=\\.DS_Store -pkg="templates" -o src/templates/templates.go templates/...

echo "$LogPrefix `date +"%H:%M:%S"` src package"
gofmt -w src/

# 交叉编译
case  $OS  in   
    linux)  
        # Linux
        echo "$LogPrefix `date +"%H:%M:%S"` build linux bin"
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install gin_sample
        ;;  
    *) 
        # 本机
        echo "$LogPrefix `date +"%H:%M:%S"` build local bin"
        go install gin_sample
        ;;
esac 

echo -e "$LogPrefix `date +"%H:%M:%S"` \033[42;37m finished \033[0m"