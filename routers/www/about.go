package www

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AboutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "about.tmpl", gin.H{
		"title": "About",
	})
}
