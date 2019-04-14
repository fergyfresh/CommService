package main

import (
	"DMCommunicationService/src/github.com/DMEvanCT/MailGun"
	"DMCommunicationService/src/github.com/DMEvanCT/Slack"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main() {
	r := mux.NewRouter()
	// IMFunc.go has the function for this
	r.HandleFunc("/api/comm/slack/message", Slack.Communication).Methods("POST")
	// EmailFunc.go has the function for this
	r.HandleFunc("/api/comm/mailgun/message", MailGunz.MailGunComm).Methods("POST")

	// listens on localhost:8080
	log.Fatal(http.ListenAndServe(":8080", r))



}


