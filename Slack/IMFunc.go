package Slack

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"net/http"

)

type SlackData struct {
	SlackMessage string `json:"SlackMessage"  bson:"SlackMessage"`
	Service string `json:"Service"  bson:"Service"`
}


var service string;

func Communication(w http.ResponseWriter, r *http.Request) {


	var slackData  SlackData
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewDecoder(r.Body).Decode(&slackData)


	// Viper configuration this reads from the config file
	viper.AddConfigPath("/etc/commservice/")
	viper.SetConfigName("comconfig")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Not able to read in config", err)
	}

	// set the token and channel name to the token and channel name inside of the slack block of the yaml
	token := viper.GetString("slack.token")
	channelName := viper.GetString("slack.channel")
	// Selecting service to
	service = slackData.Service
	switch service {
	case "slack":
		SlackMessage(slackData.SlackMessage, token, channelName)

	default:
		log.Println("Service not yet defined." + service)
	}

}
