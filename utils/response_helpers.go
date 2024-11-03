package utils

import (
	"github.com/gin-gonic/gin"
)

func CreateResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}
