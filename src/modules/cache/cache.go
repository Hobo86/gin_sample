package cache

import (
	"time"

	gin_cache "github.com/gin-gonic/contrib/cache"
	"github.com/gin-gonic/gin"
)

const (
	DEFAULT    = time.Duration(0)
	FOREVER    = time.Duration(-1)
	DefaultKey = "modules/cache"
)

func Cache(store gin_cache.CacheStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(DefaultKey, store)
		c.Next()
	}
}

// shortcut to get Cache
func Default(c *gin.Context) gin_cache.CacheStore {
	return c.MustGet(DefaultKey).(gin_cache.CacheStore)
}
