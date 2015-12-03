package www

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DemoHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "www/demo", H(c, gin.H{
		"title": "Demo",
	}))
}
