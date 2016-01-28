package render

import (
	"net/http"

	p "github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"

	"conf"
	"modules/log"
)

func pongo2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		tmpl, context, err := getContext(c)
		if err == nil {
			c.HTML(http.StatusOK, tmpl+conf.TMPL_SUFFIX, p.Context(context))
		} else {
			log.DebugPrint("Render Error: %v", err)
		}
	}
}
