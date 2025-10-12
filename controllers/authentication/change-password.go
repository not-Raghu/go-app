package authentication

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/not-raghu/go-app/db"
	"github.com/not-raghu/go-app/models"
	"golang.org/x/crypto/bcrypt"
)

type ChangePassword struct {
	NewPassword string `json:"newpassword"`
	Otp         string `json:"otp"`
}

func ChangePass() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json ChangePassword
		if err := c.ShouldBindBodyWithJSON(&json); err != nil {
			c.JSON(400, gin.H{
				"error": "invalid request,send all fields",
			})
			c.AbortWithError(400, err)
			return
		}
		anyemail, ok := c.Get("email")

		if !ok {
			c.JSON(400, gin.H{
				"error": "incorrect token, please try again",
			})
			c.AbortWithError(400, errors.New("incorrect token"))
			return
		}
		email := fmt.Sprintf("%s", anyemail)
		key := email + ":forgotpassword"
		val, err := db.RedisClient.Get(db.Ctx, key).Result()

		if err != nil {
			c.JSON(400, gin.H{
				"error": "otp expired,try again",
			})
			c.AbortWithError(400, err)
			return
		}

		if val != json.Otp {
			c.JSON(400, gin.H{
				"error": "incorrect otp",
			})
			c.AbortWithError(400, errors.New("incorrect otp"))
			return
		}
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(json.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "error generating hash",
			})
			c.AbortWithError(400, err)
			return
		}
		var user models.User
		queryResult := db.Db.Where(&models.User{Email: email}).First(&user)

		if queryResult.Error != nil {
			c.JSON(400, gin.H{
				"error": "user not found",
			})
			c.AbortWithError(400, queryResult.Error)
			return
		}

		user.Password = string(hashedPass)
		db.Db.Save(&user)

		c.JSON(200, gin.H{
			"message": "changed password successfully",
		})
		db.RedisClient.Del(db.Ctx, key)
	}
}
