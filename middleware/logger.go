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
			//time logging
			start := time.Now()
			c.Next()
			latency := time.Since(start)

			fmt.Println("took ", latency, " for end point ", path)
			//more logggers

		} else {
			c.Next()
		}
	}
}
