package test

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func Test() gin.HandlerFunc {
	err := errors.New("OMFG VERY BAD ERROR , PLS DIE")
	return func(c *gin.Context) {
		c.AbortWithError(500, err)
		c.JSON(500, gin.H{
			"test request": "Hallo mein Freund",
		})
	}
}
