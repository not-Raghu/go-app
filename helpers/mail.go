package helpers

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMail() error {
	from := os.Getenv("MAIL")
	password := os.Getenv("MAILPASS")

	toList := []string{"example@gmail.com"}
	host := "smtp.gmail.com"
	port := "587"
	msg := "Hello geeks!!!"
	body := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, toList, body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully sent mail to all user in toList")
	return nil
}
