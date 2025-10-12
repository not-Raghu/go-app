package authentication

import (
	"errors"
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
			c.AbortWithError(400, errors.New("not authorized"))
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
			c.AbortWithError(400, err)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(400, gin.H{
				"error": "incorrect token",
			})
			c.AbortWithError(400, errors.New("incorrect token"))
			return
		}

		c.Set("email", claims["email"])
		c.Next()
	}
}
