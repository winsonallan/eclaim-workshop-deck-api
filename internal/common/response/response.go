package response

import (
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"success": true,
		"status_code": statusCode,
		"message": message,
		"data":    data,
	})
}

func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"status_code": statusCode,

		"error":   message,
	})
}