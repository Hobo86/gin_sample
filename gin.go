package main

import (
	"time"

	gin_cache "github.com/gin-gonic/contrib/cache"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/olebedev/staticbin"

	"assets"
	"conf"
	"models"
	"modules/auth"
	"modules/cache"
	"modules/render"
	"routers/api"
	"routers/www"
)

const (
	BINDATA   = true
	MEMCACHED = true
	REDIS     = true
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// 静态资源
	if BINDATA {
		r.Use(staticbin.Static(assets.Asset, staticbin.Options{
			Dir: "/",
		}))
	} else {
		r.Static("/assets", "./assets")
	}

	// 模板
	if BINDATA {
		r.HTMLRender = render.LoadBindataTemplates("templates")
	} else {
		r.HTMLRender = render.LoadTemplates("templates")
	}

	// 模型
	model := models.Model()
	r.Use(model)

	// Session
	if REDIS {
		store, err := sessions.NewRedisStore(10, "tcp", conf.REDIS_SERVER, conf.REDIS_PWD, []byte("secret"))
		if err != nil {
			panic(err)
		}
		r.Use(sessions.Sessions("mysession", store))
	} else {
		store := sessions.NewCookieStore([]byte("secret"))
		r.Use(sessions.Sessions("mysession", store))
	}

	// Cache
	if MEMCACHED {
		cacheStore := gin_cache.NewMemcachedStore([]string{conf.MEMCACHED_SERVER}, time.Hour)
		r.Use(cache.Cache(cacheStore))
	} else {
		cacheStore := gin_cache.NewInMemoryStore(time.Hour)
		r.Use(cache.Cache(cacheStore))
	}

	r.Use(auth.Auth(models.GenerateAnonymousUser))

	r.GET("", www.HomeHandler)
	r.GET("/login", www.LoginHandler)
	r.GET("/register", www.RegisterHandler)
	r.GET("/logout", www.LogoutHandler)
	r.POST("/login", www.LoginPostHandler)
	r.POST("/register", www.RegisterPostHandler)

	demo := r.Group("/demo")
	{
		demo.GET("", www.DemoHandler)
	}

	user := r.Group("/user")
	user.Use(auth.LoginRequired)
	{
		user.GET("/:id", www.UserHandler)
	}

	about := r.Group("/about")
	{
		about.GET("", www.AboutHandler)
	}

	gApi := r.Group("/api")
	{
		gApi.GET("/user/:id", api.UserHandler)
		gApi.GET("/login", api.UserLoginHandler)
		gApi.GET("/register", api.UserRegisterHandler)
	}

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
