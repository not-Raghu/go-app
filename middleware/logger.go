package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		if path != "" {

			s := time.Now()
			c.Next()
			latency := time.Since(s)
			fmt.Println("took " + latency.String() + " - " + path)

		} else {
			c.Next()
		}
	}
}
