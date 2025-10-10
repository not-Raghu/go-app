package authhelpers

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/not-raghu/go-app/db"
	"github.com/not-raghu/go-app/models"
	"gorm.io/gorm"
)

func GenerateOtp(length int) (string, error) {
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
func ValidEmail(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func ValidPass(p string) bool {

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(p)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(p)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(p)
	hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString(p)

	if !hasLower || !hasUpper || !hasNumber || !hasSpecial || len(p) < 8 {
		return false
	}

	return true
}

func CheckUserInDb(email string) (*models.User, bool) {
	var user models.User
	result := db.Db.Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, false
	} else if result.Error != nil {
		log.Println("DB error:", result.Error)
		return nil, false
	}

	return &user, true

}

func SendOtpMail(to string, otp string, purpose string) {
	from := os.Getenv("MAIL")
	password := os.Getenv("MAILPASS")

	host := os.Getenv("SMTPHOST")
	port := os.Getenv("SMTPPORT")
	var msg []byte

	switch purpose {
	case "register":
		{
			msg = []byte("Subject: OTP FOR GOBBER!\r\n" + "\r\n" + "YOUR OTP TO REGISTER INTO GOBBER IS " + otp + ".\r\n" + "note that this otp expires in 10 mins" + ".\r\n")
		}
	case "forgotpassword":
		{
			msg = []byte("Subject: PASSWORD RESET FOR GOBBER!\r\n" + "\r\n" + "OTP TO APPLY FOR CHANGING PASSWORD IS " + otp + ".\r\n" + "note this otp expires in 10 mins" + ".\r\n")
		}

	}

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, []string{to}, []byte(msg))

	if err != nil {
		fmt.Println("error sending mail", err)
	}
}

func JWTToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["exp"] = time.Now().Add(1 * time.Second).Unix()

	tokenstr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Printf("couldn't not create jwt %s", err.Error())
		return "", err
	}

	return tokenstr, nil
}
