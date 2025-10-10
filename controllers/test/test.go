package test

import (
	"github.com/gin-gonic/gin"
)

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(500, gin.H{
			"test request": "Hallo mein Freund",
		})

	}
}
