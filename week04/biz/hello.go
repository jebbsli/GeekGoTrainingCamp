package biz

import "github.com/gin-gonic/gin"

func Hello(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello Gin!",
	})
}
