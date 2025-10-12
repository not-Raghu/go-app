package authentication

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/not-raghu/go-app/db"
	"github.com/not-raghu/go-app/helpers"
	"github.com/not-raghu/go-app/helpers/authhelpers"
	"github.com/not-raghu/go-app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type registerRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json registerRequest
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": "email and password are required"})
			c.AbortWithError(400, err)
			return
		}

		if err := validateRegistrationInput(&json); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			c.AbortWithError(400, err)
			return
		}

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "internal server error"})
			c.AbortWithError(500, err)
			return
		}

		err = handleUserCreationTransaction(db.Db, &json, string(hashedPass))
		if err != nil {
			if err.Error() == "email already in use" {
				c.JSON(409, gin.H{"error": err.Error()})
				c.AbortWithError(409, err)
			} else {
				c.JSON(500, gin.H{"error": "failed to create user"})
				c.AbortWithError(500, err)
			}
			return
		}

		if err := sendVerificationOTP(json.Email); err != nil {
			c.JSON(500, gin.H{
				"message": "failed to register.",
			})
			c.AbortWithError(500, err)
			return
		}

		token, err := authhelpers.JWTToken(json.Email)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "internal server error",
			})
			c.AbortWithError(500, err)
			return
		}

		c.JSON(200, gin.H{
			"message": "registration successful.otp has been sent to your account for verification.",
			"token":   token,
		})
	}
}

func validateRegistrationInput(req *registerRequest) error {
	if !authhelpers.ValidEmail(req.Email) {
		return errors.New("invalid email format")
	}
	if !authhelpers.ValidPass(req.Password) {
		return errors.New("bad password")
	}
	return nil
}

func handleUserCreationTransaction(gormDB *gorm.DB, req *registerRequest, hashedPass string) error {
	return gormDB.Transaction(func(tx *gorm.DB) error {
		var user models.User
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("email = ?", req.Email).First(&user).Error

		if err == nil {
			if user.IsVerified {
				return errors.New("email already in use")
			}
			return tx.Model(&user).Updates(models.User{
				Name:     helpers.GenerateNames(),
				Password: hashedPass,
			}).Error
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			newUser := models.User{
				Name:       helpers.GenerateNames(),
				Email:      req.Email,
				Password:   hashedPass,
				IsVerified: false,
			}
			if err := tx.Create(&newUser).Error; err != nil {
				return errors.New("could not create user record")
			}
			return nil
		}

		return errors.New("db query failed")
	})
}

func sendVerificationOTP(email string) error {
	otp, err := authhelpers.GenerateOtp(6)
	if err != nil {
		return fmt.Errorf("failed to generate otp: %w", err)
	}

	if err := db.RedisClient.Set(db.Ctx, email, otp, 10*time.Minute).Err(); err != nil {
		return fmt.Errorf("failed to store otp in redis: %w", err)
	}

	go func() {
		err := authhelpers.SendOtpMail(email, otp, "register")
		if err != nil {
			log.Printf("ERROR: Sending OTP email to %s failed: %v", email, err)
		}
	}()

	return nil
}
