package authentication

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/not-raghu/go-app/db"
	"github.com/not-raghu/go-app/models"
	"gorm.io/gorm"
)

type verifyOtp struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

var ctx = context.Background()

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

		user, err := gorm.G[models.User](db.Db).Take(ctx)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(400, gin.H{
					"error": "user not found in database, please register with email first",
				})
				return
			}
		}

		if otp == json.Otp {
			user.Is_Verified = true
		}

		db.Db.Save(&user)

		c.JSON(200, gin.H{
			"message": "valid otp",
		})

	}
}
