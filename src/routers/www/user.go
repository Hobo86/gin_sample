package www

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

	c.HTML(http.StatusOK, "www/user", H(c, gin.H{
		"title": "User",
		"user":  u,
	}))
}
