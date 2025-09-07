package authentication

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Print("reached")
		c.String(200, "recahed register route")
	}
}
