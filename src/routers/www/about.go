package www

import (
	"github.com/gin-gonic/gin"
)

func AboutHandler(c *gin.Context) {
	c.Set("tmpl", "www/about")
	c.Set("data", map[string]interface{}{
		"title": "About",
	})
}
