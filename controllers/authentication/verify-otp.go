package authentication

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/not-raghu/go-app/db"
	"github.com/not-raghu/go-app/models"
)

type verifyOtp struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

func VerifyOtp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json verifyOtp

		if err := c.ShouldBindBodyWithJSON(&json); err != nil {
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect fields"})
			}
			return
		}

		otp := db.RedisClient.Get(db.Ctx, json.Email).Val()

		if otp != json.Otp {
			c.JSON(400, gin.H{
				"error": "incorrect otp sent,please enter right otp",
			})
			return
		}

		var user models.User
		db.Db.Where(&models.User{Email: json.Email}).First(&user)
		user.IsVerified = true
		db.Db.Save(&user)

		val := db.RedisClient.Del(db.Ctx, json.Email)
		fmt.Println(val)

		c.JSON(200, gin.H{
			"message": "user created and verified",
		})
	}
}
