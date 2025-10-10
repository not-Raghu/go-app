package authentication

import (
	"github.com/gin-gonic/gin"
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
				"error": "invalid request",
			})
			return
		}

	}
}
