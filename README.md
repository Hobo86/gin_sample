#Gin Sample
==============

#测试
./run.sh

#打包
./build.sh

#Bindata
https://github.com/olebedev/staticbin

#Bindata打包
https://github.com/jteeuwen/go-bindata
go-bindata -ignore=\\.DS_Store -pkg="assets" -o src/assets/assets.go assets/...
go-bindata -ignore=\\.DS_Store -pkg="templates" -o src/templates/templates.go templates/...

#交叉编译
cd $GOROOT/src
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./make.bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install gin_sample