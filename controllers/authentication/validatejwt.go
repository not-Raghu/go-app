package authentication

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Validatejwt() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(400, gin.H{
				"error": "not authorized",
			})
			c.Abort()
		}

		sentToken := strings.Split(authHeader, " ")[1]
		token, err := jwt.Parse(sentToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(400, gin.H{
				"error": "invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
