package www

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DemoHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "www/demo", gin.H{
		"title": "Demo",
	})
}
