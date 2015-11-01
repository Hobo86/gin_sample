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
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		panic(err)
	}

	model := models.Default(c)
	u := model.GetUserById(id)

	value := -1
	cacheStore := cache.Default(c)
	if id == 1 {
		value = 0
		// cacheStore.Set("userId", 1, cache.FOREVER)
	} else {
		cacheStore.Get("userId", &value)
	}

	c.JSON(http.StatusOK, gin.H{
		"title": "User",
		"user":  u,
		"value": value,
	})
}

func UserLoginHandler(c *gin.Context) {

	c.JSON(200, map[string]interface{}{"URI": "api user login"})
}

func UserRegisterHandler(c *gin.Context) {

	c.JSON(200, map[string]interface{}{"URI": "api user regist"})
}
