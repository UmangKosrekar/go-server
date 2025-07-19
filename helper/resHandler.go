package helper

import (
	"github.com/gin-gonic/gin"
)

func ResponseHandler(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
	ctx.Done()
}
