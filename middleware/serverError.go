package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors.ByType(500)) > 0 {
			c.JSON(http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
			c.Abort()
		}

	}
}
