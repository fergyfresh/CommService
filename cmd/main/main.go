package main

import (
	"github.com/DMEvanCT/CommService/MailGun"
	"github.com/DMEvanCT/GoBase/Middleware"
	"github.com/DMEvanCT/CommService/Slack"
	"github.com/gorilla/mux"
	ghandlers "github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	// IMFunc.go has the function for this
	r.HandleFunc("/api/comm/slack/message", Slack.Communication).Methods("POST")
	// EmailFunc.go has the function for this
	r.HandleFunc("/api/comm/mailgun/message", MailGunz.MailGunComm).Methods("POST")
	// Multi Person mailgun
	r.HandleFunc("/api/comm/mailgun/multimessage", MailGunz.MailGunMulti).Methods("POST")
	// writes all requests in apache format to standard out and recovers from panic
	http.Handle("/", middleware.PanicRecoveryHandler((ghandlers.LoggingHandler(os.Stdout, r))))
	// listens on localhost:8080
	log.Fatal(http.ListenAndServe(":8080", nil))


}


