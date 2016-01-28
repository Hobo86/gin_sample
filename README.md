#Gin Sample
==============

#测试
```
./run.sh
```

#打包
```
./build.sh
```

#Bindata
https://github.com/olebedev/staticbin

#Bindata打包
https://github.com/jteeuwen/go-bindata
```
go-bindata -ignore=\\.DS_Store -pkg="assets" -o src/assets/assets.go assets/...
go-bindata -ignore=\\.DS_Store -pkg="templates" -o src/templates/templates.go templates/...
```

#交叉编译
```
cd $GOROOT/src
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./make.bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install gin_sample
```

#依赖
##Gin框架
```
github.com/gin-gonic/gin
github.com/gin-gonic/gin/render
github.com/gin-gonic/contrib/sessions
github.com/gin-gonic/contrib/cache
```
##ORM
```
github.com/jinzhu/gorm
```

##模板
```
github.com/flosch/pongo2
github.com/robvdl/pongo2gin
```
## 其他
```
github.com/olebedev/staticbin
```

#框架功能
##配置
- [x] conf

##缓存
- [x] Memcached
- [x] Memory
	查询缓存
	SQL解析缓存
##Session
- [x] Cookie
- [x] Redis

##FLASH
	
##权限
	Auth
	AuthApi
##模型
### DB
	多连接及切换
### ORM
	CURD
	关联
	事务
	form绑定
	自动验证
	分页
##模板
- [x] PONGO2

	模板布局
	公共参数
	模板替换
##日志

##安全
	SQL注入
	XSS
	表单令牌
	验证码
##部署
	模块部署
	域名部署
##多语言



