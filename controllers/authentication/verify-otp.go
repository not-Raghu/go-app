package authentication

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/not-raghu/goober/db"
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

		if otp == json.Otp {
			c.JSON(400, gin.H{
				"error": "incorrect otp sent,please enter right otp",
			})
			return
		}

	}
}
