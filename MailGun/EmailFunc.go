package MailGunz

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"net/http"
)



type MailGunData struct {
	Message string `json:"Message"  bson:"Message"`
	Subject string `json:"Subject"  bson:"Subject"`
	To string `json:"To"  bson:"To"`
	Service string `json:"Service"  bson:"Service"`
}




func MailGunComm(w http.ResponseWriter, r *http.Request) {


	var   MailGunData MailGunData
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewDecoder(r.Body).Decode(&MailGunData)


	// Viper configuration this reads from the config file
	viper.AddConfigPath("/etc/commservice/")
	viper.SetConfigName("comconfig")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Not able to read in config", err)
	}

	// set the token and channel name to the token and channel name inside of the mailgun block of the yaml
	token := viper.GetString("mailgun.apiKey")
	domain := viper.GetString("mailgun.domain")
	sender := viper.GetString("mailgun.Sender")
	senderName := viper.GetString("mailgun.SenderName")
	// Selecting service to
	service := MailGunData.Service
	switch service {
	case "MailGun":
		SendMailGunEmail(domain, token, MailGunData.Subject, MailGunData.Message, MailGunData.To, sender, senderName)


	}


}

