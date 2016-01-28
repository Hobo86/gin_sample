package log

import (
	l "log"

	"github.com/gin-gonic/gin"
)

func DebugPrint(format string, values ...interface{}) {
	if gin.IsDebugging() {
		l.Printf("[Debug] "+format, values...)
	}
}

func DebugPrinError(err error) {
	if err != nil {
		DebugPrint("[ERROR] %v\n", err)
	}
}
