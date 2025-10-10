package authentication

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/not-raghu/go-app/db"
	"github.com/not-raghu/go-app/helpers/authhelpers"
)

type forgotPass struct {
	Email string `json:"email"`
}

func ForgotPass() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json forgotPass
		// check for the fields
		if err := c.ShouldBindJSON(&json); err != nil {
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect data fields sent"})
				return
			}
		}
		ok := authhelpers.ValidEmail(json.Email)
		if !ok {
			c.JSON(200, gin.H{
				"error": "incorrect email",
			})
			return
		}

		user, ok := authhelpers.CheckUserInDb(json.Email)

		//idiot never registered in the first place || maybe a hecker
		if !ok || user == nil {
			c.JSON(200, gin.H{
				"error": "user not found" + json.Email,
			})
			return
		}

		if !user.IsVerified {
			c.JSON(400, gin.H{
				"error": "user not found",
			})
			return
		}

		otp, err := authhelpers.GenerateOtp(6)
		if err != nil {
			log.Printf("generating otp error")
			c.JSON(500, gin.H{
				"error": "could not generate otp",
			})
			return
		}

		//send mail
		go authhelpers.SendOtpMail(json.Email, otp, "forgotpassword")
		//token
		tokenstr, err := authhelpers.JWTToken(user.Email)

		if err != nil {
			c.JSON(500, gin.H{
				"error": "couldn't generate token, try again later",
			})
			return
		}
		//set in red
		key := json.Email + ":forgotpassword"
		otpstatus := db.RedisClient.Set(db.Ctx, key, otp, 10*time.Minute).Err()

		if otpstatus != nil {
			c.JSON(500, gin.H{
				"error": "internal server error",
			})
			log.Printf("redis failed %s", otpstatus)
			return
		}

		c.JSON(200, gin.H{
			"message": "otp send to the mail " + user.Email,
			"token":   tokenstr,
		})

	}
}

//when he forgets his password, he sends his mail , we send a otp to his mail & token and then
//he sends his otp and token - verify them from the db
