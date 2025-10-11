package helpers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func ErrorLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				errstr := fmt.Sprintf("[Error] Method: %s | Path: %s | Error: %v\n",
					c.Request.Method, c.Request.URL.Path, err.Error())
				gin.DefaultErrorWriter.Write([]byte(errstr))
				log.Println(errstr)
			}
		}

	}
}
