package main

import (
	"bytes"
	"log"
	"net/smtp"
)

func mail(name, email, message string) bool {
	sent := false
	// Connect to the remote SMTP server.
	c, err := smtp.Dial("localhost:25")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	// Set the sender and recipient.
	c.Mail("admin@pcilookup.com")
	c.Rcpt("joshuagreen118@gmail.com")
	c.Rcpt("iamwill.knoll@gmail.com")

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString("To: joshuagreen118@gmail.com\r\n" + "Subject: Contact Form Submission!\r\n" + "\r\n" +"Message from " + name + " (" + email + ").\n" + message)
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	} else {
		sent = true
	}

	return sent
}