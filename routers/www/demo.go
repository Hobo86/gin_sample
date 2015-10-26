package www

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DemoHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "demo.tmpl", gin.H{
		"title": "Demo",
	})
}
