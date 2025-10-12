package authentication

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/not-raghu/go-app/db"
	"github.com/not-raghu/go-app/models"
)

type verifyOtp struct {
	Token string `json:"token"`
	Otp   string `json:"otp"`
}

func VerifyOtp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json verifyOtp

		if err := c.ShouldBindJSON(&json); err != nil {
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect fields"})
				c.AbortWithError(400, err)
			}
			return
		}

		emailRaw, ok := c.Get("email")
		if !ok {
			c.JSON(400, gin.H{
				"error": "no token provided",
			})
			c.AbortWithError(400, errors.New("no token provided"))
			return
		}

		email, ok := emailRaw.(string)
		if !ok || email == "" {
			c.JSON(400, gin.H{
				"error": "invalid token format",
			})
			c.AbortWithError(400, errors.New("invalid token format"))
			return
		}

		otp := db.RedisClient.Get(db.Ctx, email).Val()

		if otp == "" || json.Otp == "" {
			if otp == "" {
				c.JSON(404, gin.H{
					"error": "otp expired, try again",
				})
				c.AbortWithError(400, errors.New("otp expired"))
				return
			} else {
				c.JSON(400, gin.H{
					"error": "enter a valid otp",
				})
				c.AbortWithError(400, errors.New("user didn't enter otp"))
				return
			}
		}

		if otp != json.Otp {
			c.JSON(400, gin.H{
				"error": "incorrect otp sent,please enter right otp",
			})
			c.AbortWithError(400, errors.New("incorrect otp"))
			return
		}

		var user models.User
		result := db.Db.Where("email = ?", email).First(&user)
		if result.Error != nil {
			c.JSON(400, gin.H{
				"error": "couldn't find user, register before verifying",
			})
			c.AbortWithError(400, result.Error)
			return
		}

		user.IsVerified = true
		result = db.Db.Save(&user)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "internal server error",
			})
			c.AbortWithError(500, result.Error)
		}
		val := db.RedisClient.Del(db.Ctx, email)
		fmt.Println(val)

		c.JSON(200, gin.H{
			"message": "user created and verified",
		})
	}
}
