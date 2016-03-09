package api

import (
	"net/http"
	"strconv"
	. "time"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"

	"models"
	"modules/cache"
)

func UserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		panic(err)
	}

	model := models.Default(c)
	u := model.GetUserById(id)

	value := -1
	cacheStore := cache.Default(c)
	if id == 1 {
		value = 0
		cacheStore.Set("userId", 1, Minute)
	} else {
		cacheStore.Get("userId", &value)
	}

	c.JSON(http.StatusOK, gin.H{
		"title":      "User",
		"user":       u,
		"value":      value,
		"host":       c.Request.Host,
		"referer":    c.Request.Referer(),
		"method":     c.Request.Method,
		"RequestURI": c.Request.RequestURI,
		"RemoteAddr": c.Request.RemoteAddr,
		"url":        c.Request.URL.String(),
		"path":       c.Request.URL.Path,
		"query":      c.Request.URL.Query().Encode(),
		"uri":        c.Request.URL.RequestURI(),
		"rawquery":   c.Request.URL.RawQuery,
	})
}

func UserLoginHandler(c *gin.Context) {

	c.JSON(200, map[string]interface{}{"URI": "api user login"})
}

func UserRegisterHandler(c *gin.Context) {

	c.JSON(200, map[string]interface{}{"URI": "api user regist"})
}
