package main

import (
	"html/template"
	// "net/http"
	// "log"
	"path/filepath"

	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	// "./config"
	"./models"
	// "./routers/api"
	"./modules/auth"
	"./routers/api"
	"./routers/www"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// 静态资源
	r.Static("/assets", "./assets")

	r.HTMLRender = loadTemplates("./templates")

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

func loadTemplates(templatesDir string) multitemplate.Render {
	r := multitemplate.New()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		tmpl := template.Must(template.ParseFiles(files...))
		// log.Printf(filepath.Base(layout))
		r.Add(filepath.Base(layout), tmpl)
	}
	return r
}
