package sessions

import (
	gin_sessions "github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"conf"
)

func Sessions() gin.HandlerFunc {
	switch conf.SESSION_STORE {
	case conf.REDIS:
		store, err := gin_sessions.NewRedisStore(10, "tcp", conf.REDIS_SERVER, conf.REDIS_PWD, []byte("secret"))
		if err != nil {
			panic(err)
		}
		return gin_sessions.Sessions("mysession", store)
	default:
		store := gin_sessions.NewCookieStore([]byte("secret"))
		return gin_sessions.Sessions("mysession", store)
	}
}
