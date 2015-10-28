package main

import (
	// "log"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/olebedev/staticbin"

	"assets"
	"models"
	"modules/auth"
	"modules/render"
	"routers/api"
	"routers/www"
)

const (
	BINDATA = false
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
		r.HTMLRender = render.LoadTemplates("./templates")
	}

	// 模型
	model := models.Model()
	r.Use(model)

	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

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
