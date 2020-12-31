package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mohit810/outlook-actionable-message/controller"
	"github.com/mohit810/outlook-actionable-message/prodtesting"
)

func init() {
	prodtesting.GenPem()
}

func main() {
	port := flag.Int("port", 8080, "http server port")
	flag.Parse()
	r := httprouter.New()
	uc := controller.NewUserController()
	r.POST("/sendmail", uc.SendingNormalActionableMessage)
	r.POST("/outlookresponse", uc.MailResponse)
	fmt.Println(http.ListenAndServeTLS(":"+strconv.Itoa(*port), "cert.pem", "key.pem", r))
}
