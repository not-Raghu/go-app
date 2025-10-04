package authentication

import (
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"net/smtp"
	"os"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/not-raghu/go-app/db"
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

		if ok := validEmail(json.Email); !ok {
			c.JSON(400, gin.H{
				"error": "incorrect email format",
			})
			return
		}

		if ok := validPass(json.Password); !ok {
			c.JSON(400, gin.H{
				"error": "bad password",
			})
			return
		}

		//check in db for email if already taken.

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

		go sendOtpMail(json.Email, otp)

		c.JSON(200, gin.H{
			"message": "otp sent to your mail, check it and enter that",
		})
	}
}

func sendOtpMail(to string, otp string) {
	from := os.Getenv("MAIL")
	password := os.Getenv("MAILPASS")

	host := os.Getenv("SMTPHOST")
	port := os.Getenv("SMTPPORT")

	msg := []byte("Subject: OTP FOR GOBBER!\r\n" + "\r\n" + "YOUR OTP TO REGISTER INTO GOBBER IS " + otp + ".\r\n" + "note that this otp expires in 10 mins" + ".\r\n")

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, []string{to}, []byte(msg))

	if err != nil {
		fmt.Println("error sending mail", err)
	}
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

func validEmail(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func validPass(p string) bool {

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(p)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(p)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(p)
	hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString(p)

	if !hasLower || !hasUpper || !hasNumber || !hasSpecial || len(p) < 8 {
		return false
	}

	return true
}
