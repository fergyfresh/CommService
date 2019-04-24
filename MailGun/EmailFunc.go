package MailGunz

import (
	"encoding/json"
	"github.com/DMEvanCT/Auth"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type Database struct {
	Database string `json:"Database"  bson:"Database"`
}

type MailGunData struct {
	Message string `json:"Message"  bson:"Message"`
	Subject string `json:"Subject"  bson:"Subject"`
	To string `json:"To"  bson:"To"`
	Service string `json:"Service"  bson:"Service"`
}


type MailGunReturnData struct {
	Sent string  `json:"Sent"  bson:"Sent"`
	Subject string `json:"Subject"  bson:"Subject"`
	Message string `json:"Message"  bson:"Message"`
	To string `json:"To"  bson:"To"`
}

type MailGunMultiSend struct {

	MailGunData MailGunData
	QueryField string `json:"QueryField"  bson:"QueryField"`
	Condition string `json:"Condition"  bson:"Condition"`
	State string `json:"State"  bson:"State"`
	City string `json:"City"  bson:"City"`
	Database Database

}



func MailGunComm(w http.ResponseWriter, r *http.Request) {
	var   MailGunData MailGunData
	w.Header().Set("Content-Type", "application/json")
	akey := r.Header.Get("x-auth-token")
	username := r.Header.Get("x-auth-user")
	//auth := Auth.AuthenticatedUser("c4a3acd4-2ef6-4a5c-b97a-0aa5578503cf", "Clarity")
	auth := Auth.AuthenticatedUser(akey, username)
	if auth == false {
		w.WriteHeader(http.StatusForbidden)

		authenticated := Auth.Authenticated{
			Authenticated: "False",
		}
		json.NewEncoder(w).Encode(authenticated)


		return


	}





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
		MailGunInfo := MailGunReturnData{
			Sent: "Yes",
			Subject: MailGunData.Subject,
			Message: MailGunData.Message,
			To: MailGunData.To,


		}
		json.NewEncoder(w).Encode(MailGunInfo)




	}

	}

func MailGunMulti(w http.ResponseWriter, r *http.Request) {
	var   MultiSendData MailGunMultiSend
	//var MailGunData MailGunData
	// var Database Database
	w.Header().Set("Content-Type", "application/json")
	akey := r.Header.Get("x-auth-token")
	username := r.Header.Get("x-auth-user")

	auth := Auth.AuthenticatedUser(akey, username)
	if auth == false {
		w.WriteHeader(http.StatusForbidden)

		authenticated := Auth.Authenticated{
			Authenticated: "False",
		}
		json.NewEncoder(w).Encode(authenticated)


		return


	}
	_ = json.NewDecoder(r.Body).Decode(&MultiSendData)


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

	service := MultiSendData.MailGunData.Service

	switch service {
	case "MailGunMulti":
		SendMailGunEmailByFieldNonTemplate(MultiSendData.Condition, MultiSendData.Database.Database, MultiSendData.QueryField, domain, token, MultiSendData.MailGunData.Subject, MultiSendData.MailGunData.Message,
			sender, senderName, MultiSendData.City, MultiSendData.State)
			json.NewEncoder(w).Encode(&MultiSendData)



	}
	}











