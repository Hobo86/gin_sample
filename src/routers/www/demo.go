package www

import (
	"github.com/gin-gonic/gin"
)

func DemoHandler(c *gin.Context) {
	c.Set("tmpl", "www/demo")
	c.Set("data", map[string]interface{}{
		"title": "Demo",
	})
}
