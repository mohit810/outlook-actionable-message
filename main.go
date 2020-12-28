package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	auth2 "github.com/mohit810/outlook-actionable-message/auth"
	"net"
	"net/smtp"
	"text/template"
)

func main() {

	// Sender data.
	from := "email-Id"
	password := "pwd"

	// Receiver email address.
	to := []string{
		"email-Id",
	}

	// smtp server configuration.
	smtpHost := "smtp.office365.com"
	smtpPort := "587"

	conn, err := net.Dial("tcp", "smtp.office365.com:587")
	if err != nil {
		println(err)
	}

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		println(err)
	}

	tlsconfig := &tls.Config{
		ServerName: smtpHost,
	}

	if err = c.StartTLS(tlsconfig); err != nil {
		println(err)
	}

	auth := auth2.LoginAuth(from, password)

	if err = c.Auth(auth); err != nil {
		println(err)
	}

	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		CardMessage string
		Value       string
	}{
		CardMessage: "Please take a minute and share your thoughts on 2020.",
		Value:       "{{answer.value}}",
	})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
