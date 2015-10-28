package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"

	"models"
)

func UserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		panic(err)
	}

	model := models.Default(c)
	u := model.GetUserById(id)

	c.JSON(http.StatusOK, gin.H{
		"title": "User",
		"user":  u,
	})
}

func UserLoginHandler(c *gin.Context) {

	c.JSON(200, map[string]interface{}{"URI": "api user login"})
}

func UserRegisterHandler(c *gin.Context) {

	c.JSON(200, map[string]interface{}{"URI": "api user regist"})
}
