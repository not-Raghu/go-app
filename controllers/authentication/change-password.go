package authentication

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ChangePassword struct {
	token    string
	password string
}

func ChangePass() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json ChangePassword

		if c.Header("token") == nil {
			c.JSON(400,gin.H{
				"error" : "token not provided",
			})
			return
		}

		token, err := jwt.Parse(c.Header["token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return sampleSecretKey, nil
	})

		if err := c.ShouldBindBodyWithJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": "incorrect data fields sent"})
			return
		}


		jwt.
		

	}
}
