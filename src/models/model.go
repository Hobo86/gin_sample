package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"conf"
)

const (
	DefaultKey  = "models/model"
	errorFormat = "[models] ERROR! %s\n"
)

type model struct {
	db gorm.DB
}

func Model() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := conf.DB()
		model := model{db}
		c.Set(DefaultKey, model)
		c.Next()
	}
}

// shortcut to get model
func Default(c *gin.Context) model {
	return c.MustGet(DefaultKey).(model)
}
