package cache

import (
	"time"

	gin_cache "github.com/gin-gonic/contrib/cache"
	"github.com/gin-gonic/gin"

	"conf"
)

const (
	DEFAULT    = time.Duration(0)
	FOREVER    = time.Duration(-1)
	DefaultKey = "modules/cache"
)

func Cache() gin.HandlerFunc {
	var store gin_cache.CacheStore

	switch conf.CACHE_STORE {
	case conf.MEMCACHED:
		store = gin_cache.NewMemcachedStore([]string{conf.MEMCACHED_SERVER}, time.Hour)
	default:
		store = gin_cache.NewInMemoryStore(time.Hour)
	}

	return func(c *gin.Context) {
		c.Set(DefaultKey, store)
		c.Next()
	}
}

// shortcut to get Cache
func Default(c *gin.Context) gin_cache.CacheStore {
	return c.MustGet(DefaultKey).(gin_cache.CacheStore)
}
