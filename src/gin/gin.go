package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/olebedev/staticbin"

	"assets"
	"conf"
	"models"
	"modules/auth"
	"modules/cache"
	"modules/render"
	"modules/sessions"
	"routers/api"
	"routers/www"
)

func Run() {
	if conf.GIN_RELEASE_MODE {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// 静态资源
	switch conf.STATIC_TYPE {
	case conf.BINDATA:
		r.Use(staticbin.Static(assets.Asset, staticbin.Options{
			Dir: "/",
		}))
	default:
		r.Static("/assets", "./assets")
	}

	// 模板
	r.HTMLRender = render.LoadTemplates()
	r.Use(render.Render())

	// 模型
	model := models.Model()
	r.Use(model)

	// Session
	r.Use(sessions.Sessions())

	// Cache
	r.Use(cache.Cache())

	// Auth
	r.Use(auth.Auth(models.GenerateAnonymousUser))

	// Routers
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

		gApi.GET("/post/save", api.PostSaveHandler)
		gApi.GET("/post/id/:id", api.PostHandler)
		gApi.GET("/posts/:userId/p/:p/s/:s", api.PostsHandler)
	}

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
