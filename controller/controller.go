package controller

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	auth2 "github.com/mohit810/outlook-actionable-message/auth"
	"net"
	"net/http"
	"net/smtp"
	"text/template"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) SendingNormalActionableMessage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cardMessage := r.FormValue("cardMessage")
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
		CardMessage: cardMessage,
		Value:       "{{answer.value}}",
	})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest) // 400
		err = json.NewEncoder(w).Encode(struct {
			Status string `json:"status"`
		}{
			Status: err.Error(),
		})
		if err != nil {
			fmt.Println(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 201
		err = json.NewEncoder(w).Encode(struct {
			Status string `json:"status"`
		}{
			Status: "Success",
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}
