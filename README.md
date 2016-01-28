#Gin Sample

##测试
```
./run.sh
```

##打包
```
./build.sh
```

##Bindata
https://github.com/olebedev/staticbin

##Bindata打包
https://github.com/jteeuwen/go-bindata
```
go-bindata -ignore=\\.DS_Store -pkg="assets" -o src/assets/assets.go assets/...
go-bindata -ignore=\\.DS_Store -pkg="templates" -o src/templates/templates.go templates/...
```

##交叉编译
```
cd $GOROOT/src
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./make.bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install gin_sample
```

##依赖
###Gin框架
```
github.com/gin-gonic/gin
github.com/gin-gonic/gin/render
github.com/gin-gonic/contrib/sessions
github.com/gin-gonic/contrib/cache
```
###ORM
```
github.com/jinzhu/gorm
```

###模板
```
github.com/flosch/pongo2
github.com/robvdl/pongo2gin
```
###其他
```
github.com/olebedev/staticbin
```

##框架功能
###配置
	conf

###缓存
- [x] Memcached<br>
- [x] Memory<br>
	查询缓存<br>
	SQL解析缓存<br>

###Session
- [x] Cookie<br>
- [x] Redis<br>

###FLASH
	
###权限
	Auth<br>
	AuthApi<br>
###模型
####DB
	多连接及切换
####ORM
- [x] gorm<br>
	CURD<br>
	关联<br>
	事务<br>
	form绑定<br>
	自动验证<br>
	分页<br>

###模板
- [x] PONGO2<br>
	模板布局<br>
	公共参数<br>
	模板替换<br>

###日志
- [x] 调试日志<br>
	日志分级

###安全
	SQL注入<br>
	XSS<br>
	表单令牌<br>
	验证码<br>
	
###部署
	模块部署<br>
	域名部署<br>
	
###多语言



