package authentication

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type register struct {
	email string `json:email`
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json register

		if err := c.ShouldBindJSON(&json); err != nil {
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, gin.H{"error": "poor request or cna't parse json data"})
				return
			}
		}

		c.JSON(200, gin.H{
			"message": "receieved",
		})

	}

}
