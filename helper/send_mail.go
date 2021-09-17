package helper

import (
	"log"
	"net/smtp"

	"github.com/gopher5889/interview/models"
)

func SendMail(customer *models.Customer, host, port, from, passwd, msg string) {
	// Authentication
	auth := smtp.PlainAuth("", from, passwd, host)

	//Sending email
	err := smtp.SendMail(host+":"+port, auth, from, []string{customer.Email}, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Send email successfully")
}
