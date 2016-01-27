package conf

import ()

const (
	// Session
	SESSION_STORE = "REDIS"

	// Cache
	CACHE_STORE = "MEMCACHED"

	// Tmpl
	TMPL_TYPE = "BINDATA"
	TMPL_DIR  = "templates"

	// Static
	STATIC_TYPE = "BINDATA"

	// MySQL
	DB_NAME      = "gin_db"
	DB_USER_NAME = "gin_dba"
	DB_PASSWORD  = "123456"
	DB_HOST      = "127.0.0.1"
	DB_PORT      = "3306"

	// Redis
	REDIS        = "REDIS"
	REDIS_SERVER = "127.0.0.1:6379"
	REDIS_PWD    = "123456"

	// Memcached
	MEMCACHED        = "MEMCACHED"
	MEMCACHED_SERVER = "localhost:11211"

	// Bindata
	BINDATA = "BINDATA"
)
