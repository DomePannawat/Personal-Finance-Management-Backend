package utils

import (
	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}
