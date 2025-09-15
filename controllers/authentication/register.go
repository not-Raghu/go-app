package authentication

import (
	"io"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type register struct {
	Email    string `json:"email"`
	Passowrd string `json:"password"`
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

		if ok := validPass(json.Passowrd); !ok {
			c.JSON(400, gin.H{
				"error": "min password length is 6",
			})
			return
		}

		//save it in db

		c.JSON(200, json)
	}

}

func validEmail(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
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
