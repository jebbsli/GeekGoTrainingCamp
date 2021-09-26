package api

import (
	"github.com/gin-gonic/gin"
	"week04/biz"
)

var ServerEngine *gin.Engine

func ApiRegistry() {
	ServerEngine.GET("/hello", biz.Hello)
}

func init() {
	ServerEngine = gin.Default()
}
