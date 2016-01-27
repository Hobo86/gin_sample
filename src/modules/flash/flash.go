package flash

const (
	DefaultKey = "modules/flash"
)

func Flash() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
