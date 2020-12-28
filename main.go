package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mohit810/outlook-actionable-message/controller"
	"net/http"
)

func main() {

	r := httprouter.New()
	uc := controller.NewUserController()
	r.POST("/sendmail", uc.SendingNormalActionableMessage)
	http.ListenAndServe(":8080", r)
}
