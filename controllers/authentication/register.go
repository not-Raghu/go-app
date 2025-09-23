package authentication

import (
	"crypto/rand"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/not-raghu/goober/db"
)

type register struct {
	Email string `json:"email"`
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

		if ok := validEmail(json.Email); !ok {
			c.JSON(400, gin.H{
				"error": "incorrect email format",
			})
			return
		}

		otp, err := generateOtp(6)

		if err != nil {
			c.JSON(400, gin.H{
				"error": "failed to generate otp, try again later",
			})
			return
		}

		status := db.RedisClient.Set(db.Ctx, json.Email, otp, 10*time.Minute)

		if status.Err() != nil {
			c.JSON(400, gin.H{
				"error": "failed to set otp in redis",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "otp sent to your mail, check it and enter that",
		})
	}

}

func validEmail(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func generateOtp(length int) (string, error) {
	const otpChars = "1234567890"
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

func sendMail() {

}

func validPass(p string) bool {
	//for now a simple password validation of length
	//regex:  ^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\$%\^&\*]).{8,}$    , Should contain at least a capital letter

	// Should contain at least a small letter
	// Should contain at least a number
	// Should contain at least a special character
	// And minimum length

	if p == "" || len(p) <= 6 {
		return false
	}
	return true
}
