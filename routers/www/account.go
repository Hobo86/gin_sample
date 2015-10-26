package www

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"

	"../../models"
	"../../modules/auth"
)

type LoginForm struct {
	Nickname string `form:"nickname" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	redirect := c.DefaultQuery(auth.RedirectParam, "/")
	a := auth.Default(c)
	if a.User.IsAuthenticated() {
		c.Redirect(http.StatusMovedPermanently, redirect)
		return
	}

	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title":         "Login",
		"redirectParam": auth.RedirectParam,
		"redirect":      redirect,
	})
}

func LoginPostHandler(c *gin.Context) {
	redirect := c.DefaultQuery(auth.RedirectParam, "/")
	a := auth.Default(c)
	if a.User.IsAuthenticated() {
		c.Redirect(http.StatusMovedPermanently, redirect)
		return
	}

	loginURL := fmt.Sprintf("/login?%s=%s", auth.RedirectParam, redirect)

	var form LoginForm
	if c.Bind(&form) == nil {
		model := models.Default(c)
		u := model.GetUserByNicknamePwd(form.Nickname, form.Password)
		if u != nil {
			session := sessions.Default(c)
			err := auth.AuthenticateSession(session, u)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			c.Redirect(http.StatusMovedPermanently, redirect)
			return
		} else {
			c.Redirect(http.StatusMovedPermanently, loginURL)
			return
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, loginURL)
		return
	}
}

func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	a := auth.Default(c)
	auth.Logout(session, a.User)

	c.Redirect(http.StatusMovedPermanently, "/")
}

func RegisterHandler(c *gin.Context) {
	redirect := c.DefaultQuery(auth.RedirectParam, "/")
	a := auth.Default(c)
	if a.User.IsAuthenticated() {
		log.Print("Register IsAuthenticated!")
		c.Redirect(http.StatusMovedPermanently, redirect)
		return
	}

	c.HTML(http.StatusOK, "register.tmpl", gin.H{
		"title":         "Register",
		"redirectParam": auth.RedirectParam,
		"redirect":      redirect,
	})
}

func RegisterPostHandler(c *gin.Context) {
	redirect := c.DefaultQuery(auth.RedirectParam, "/")
	a := auth.Default(c)
	if a.User.IsAuthenticated() {
		c.Redirect(http.StatusMovedPermanently, redirect)
		return
	}

	registerURL := fmt.Sprintf("/register?%s=%s", auth.RedirectParam, redirect)

	var form LoginForm
	if c.Bind(&form) == nil {
		model := models.Default(c)
		u := model.AddUserWithNicknamePwd(form.Nickname, form.Password)
		if u != nil {
			session := sessions.Default(c)
			err := auth.AuthenticateSession(session, u)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			c.Redirect(http.StatusMovedPermanently, redirect)
			return
		} else {
			log.Print("Register user add error")
			c.Redirect(http.StatusMovedPermanently, registerURL)
			return
		}
	} else {
		log.Print("Register form bind error")
		c.Redirect(http.StatusMovedPermanently, registerURL)
		return
	}
}
