package authentication

import (
	"errors"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/not-raghu/go-app/db"
	"github.com/not-raghu/go-app/helpers/authhelpers"
)

type resendOtp struct {
	Token string `json:"token"`
}

func ResendOtp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json resendOtp

		if err := c.ShouldBindBodyWithJSON(&json); err != nil {
			if err == io.EOF {
				c.JSON(400, gin.H{
					"error": "incorrect fields sent",
				})
				c.AbortWithError(400, err)
				return
			}
		}

		rawKey, ok := c.Get("email")

		if !ok {
			c.JSON(400, gin.H{
				"error": "corrupted token,retry!",
			})
			c.AbortWithError(400, errors.New("corrrupted jwt token in resend otp"))
			return

		}

		key, ok := rawKey.(string)
		if !ok || key == "" {
			c.JSON(400, gin.H{
				"error": "invlaid email,try again",
			})
			c.AbortWithError(400, errors.New("invalid token"))
			return
		}
		ttl := db.RedisClient.TTL(db.Ctx, key).Val()
		if ttl == -2 {
			c.JSON(400, gin.H{
				"error": "otp expired, request a new one",
			})
			c.AbortWithError(400, errors.New("otp doens't exist in db"))
			return
		}
		if ttl > 0 && ttl < 9*time.Minute {
			otp, _ := authhelpers.GenerateOtp(6)
			err := authhelpers.SendOtpMail(key, otp, "") //add a nice parameter later
			if err != nil {
				c.JSON(500, gin.H{
					"error": "could not send mail",
				})
				c.AbortWithError(500, err)
				return
			}
			c.JSON(200, gin.H{
				"message": "send mail successfully, check again",
			})
			return
		} else {
			c.JSON(400, gin.H{
				"message": "wait until the timer runs out before requesting another otp",
			})
			c.AbortWithError(400, errors.New("WAIT BEFORE ASKING THAT MANY MAILS SIR"))
		}

	}
}
