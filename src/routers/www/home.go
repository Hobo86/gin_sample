package www

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"models"
)

func HomeHandler(c *gin.Context) {
	id, err := strconv.ParseInt("1", 10, 64)
	if err != nil {
		panic(err)
	}
	model := models.Default(c)
	u := model.GetUserById(id)

	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": "Home",
		"user":  u,
	})
}
