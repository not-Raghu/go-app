package authentication

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/not-raghu/go-app/db"
	"github.com/not-raghu/go-app/helpers"
	"github.com/not-raghu/go-app/helpers/authhelpers"
	"github.com/not-raghu/go-app/models"
	"golang.org/x/crypto/bcrypt"
)

type register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {

		var json register

		if err := c.ShouldBindJSON(&json); err != nil {
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, gin.H{"error": "poor request or cna't parse json data"})
				return
			}
		}

		if ok := authhelpers.ValidEmail(json.Email); !ok {
			c.JSON(400, gin.H{
				"error": "incorrect email format",
			})
			return
		}

		if ok := authhelpers.ValidPass(json.Password); !ok {
			c.JSON(400, gin.H{
				"error": "bad password",
			})
			return
		}

		user, ok := authhelpers.CheckUserInDb(json.Email)

		if user != nil && user.IsVerified {
			c.JSON(400, gin.H{
				"error": "email already in use",
			})
			return
		}

		if !ok || !user.IsVerified {

			otp, err := authhelpers.GenerateOtp(6)

			if err != nil {
				c.JSON(400, gin.H{
					"error": "failed to generate otp, try again later",
				})
				return
			}

			status := db.RedisClient.Set(db.Ctx, json.Email, otp, 10*time.Minute)
			if status.Err() != nil {
				c.JSON(400, gin.H{
					"error": "failed setting up otp",
				})
				return
			}

			hashedPass, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
			if err != nil {
				c.JSON(400, gin.H{
					"error": "error generating hash",
				})
				return
			}

			go authhelpers.SendOtpMail(json.Email, otp, "register")
			db.Db.Create(&models.User{
				Name:     helpers.GeneateNames(),
				Email:    json.Email,
				Password: string(hashedPass),
			})

			c.JSON(200, gin.H{
				"message": "otp sent to your mail, check it and enter that",
			})
		}
	}
}
