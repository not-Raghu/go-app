package authentication

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ChangePassword struct {
	Password string `json:"password"`
	Otp      string `json:"otp"`
}

func ChangePass() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json ChangePassword
		c.ShouldBindBodyWithJSON(&json)

		if c.GetHeader("Authorization") == "" {
			c.JSON(400, gin.H{
				"error": "no token provideed",
			})
			return
		}

		fmt.Println(c.GetHeader("Authorization"))
		tokenbytes := string(c.GetHeader("Authorization")[1])
		token, err := jwt.Parse(string(tokenbytes), func(token *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(token)
	}
}
