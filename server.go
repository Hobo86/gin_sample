package main

import (
	"net/http"

	"github.com/go-martini/martini"
	// "github.com/martini-contrib/auth"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"

	"./config"
	"./routers/api"
	"./routers/www"

	"./models"
)

func main() {

	m := martini.Classic()

	// 静态资源
	m.Use(martini.Static("assets"))

	m.Map(config.DB())

	// m.Use(Auth)

	// martini-contrib/auth
	// m.Use(auth.Basic("steven", "123456"))

	// martini-contrib/render
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	// sessions
	store := sessions.NewCookieStore([]byte("123456"))
	m.Use(sessions.Sessions("my_session", store))

	// sessionauth
	m.Use(sessionauth.SessionUser(model.GenerateAnonymousUser))
	sessionauth.RedirectUrl = "/login"
	sessionauth.RedirectParam = "redirect-url"

	// Routes
	m.Get("/", www.HomeHandler)

	m.Group("/demo", func(r martini.Router) {
		m.Get("", www.DemoHandler)
	})

	m.Group("", func(r martini.Router) {

		m.Get("/regist", www.RegistHandler)
		m.Post("/regist", binding.Bind(model.User{}), www.RegistPostHandler)

		m.Get("/login", www.LoginHandler)
		m.Post("/login", binding.Bind(model.User{}), www.LoginPostHandler)

		m.Get("/logout", sessionauth.LoginRequired, www.LogoutHandler)

		m.Get("/user/(?P<id>[0-9]*)", sessionauth.LoginRequired, www.UserHandler)

		m.Get("/about", www.AboutHandler)
	})

	m.Group("/api", func(r martini.Router) {

		// 在authorize方法没有输出时继续执行，用作接口验证
		// m.Get("/user/:id", authorize, api.UserHandler)
		m.Get("/user/(?P<id>[0-9]*)", api.UserHandler)

		m.Get("/login", api.UserLoginHandler)

		m.Get("/regist", api.UserRegistHandler)
	})

	m.NotFound(func(r render.Render) {
		// 404
		r.HTML(404, "40x", "404-Not found!")
	})

	m.Run()
}

func InitDB() {

}

func LoginAuth(r render.Render, session sessions.Session) {
	v := session.Get("userId")
	if v == nil {
		r.HTML(200, "www/login", "")
	}
	return
}

func Auth(res http.ResponseWriter, req *http.Request, r render.Render) {
	if req.Header.Get("API-KEY") != "secret123" {
		res.WriteHeader(http.StatusUnauthorized)

		// 403
		r.HTML(403, "40x", "403-Access Forbidden!")
	}
}

func authorize(res http.ResponseWriter, req *http.Request, r render.Render) {
	if req.Header.Get("API-KEY") != "secret123" {
		res.WriteHeader(http.StatusUnauthorized)

		// 403
		r.HTML(403, "40x", "403-Access Forbidden!")
	} else {
		return
	}
}
