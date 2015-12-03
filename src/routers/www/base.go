package www

import (
	"github.com/gin-gonic/gin"

	"modules/auth"
)

func H(c *gin.Context, datas gin.H) gin.H {
	a := auth.Default(c)
	userId := a.User.UniqueId().(int64)

	h := make(map[string]interface{})

	// Common Data
	h["UserId"] = userId
	h["UserName"] = "用户名"

	// Request URL
	h["requestUrl"] = c.Request.URL.String()

	for key, value := range datas {
		h[key] = value
	}

	return h
}
