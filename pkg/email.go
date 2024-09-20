package pkg

import (
	"net/smtp"
	"yugioh-decks/config"
)

func send (body string  , to string ) error {
	
	from := config.Email
	pass := config.Pass
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}	


	return nil
}