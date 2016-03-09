package www

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"models"
)

func HomeHandler(c *gin.Context) {
	id, err := strconv.ParseUint("1", 10, 64)
	if err != nil {
		panic(err)
	}
	model := models.Default(c)
	u := model.GetUserById(id)

	c.Set("tmpl", "www/home")
	c.Set("data", map[string]interface{}{
		"title": "Home",
		"user":  u,
	})
}
