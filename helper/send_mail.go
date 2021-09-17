package helper

import (
	"log"
	"net/smtp"
	"sync"

	"github.com/gopher5889/interview/config"
)

func DoSendMail(from string, to string, msg string, config *config.Config, wg *sync.WaitGroup) {
	defer wg.Done()
	// Authentication
	auth := smtp.PlainAuth("", from, config.Password, config.SmtpHost)

	//Sending email

	err := smtp.SendMail(config.SmtpHost+":"+config.SmtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}

}
