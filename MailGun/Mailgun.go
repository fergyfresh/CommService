package MailGunz

import (
	"context"
	"database/sql"
	"github.com/mailgun/mailgun-go/v3"
	"github.com/spf13/viper"
	"log"
	"time"
)



type QueryDat struct {
	Condition string
}



func SendMailGunEmail(domain, apiKey, subject, message, to, sender, senderName string) (string,  error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		 senderName + " " +  "<" + sender + ">",
		subject,
		message,
		to,
	)
	m.SetHtml( "<html><h1>" + subject +"</h1><p>" +  message + "<p></html>")
	//m.AddAttachment("files/test.jpg")
	//m.AddAttachment("files/test.txt")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err  := mg.Send(ctx, m)
	return id, err
}

func SendMaulGunEmailTemplate(domain, apiKey, subject, message, to, sender, senderName string) (string,  error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		senderName + " " +  "<" + sender + ">",
		subject,
		message,
		to,

	)
	m.SetHtml( "<html><h1>" + subject +"</h1><p>" +  message + "<p></html>")
	//m.AddAttachment("files/test.jpg")
	//m.AddAttachment("files/test.txt")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err  := mg.Send(ctx, m)
	return id, err
}



func SendMailGunEmailByFieldNonTemplate(Condition, DB, QueryField,  domain, apiKey, subject, message, sender, senderName, City, State  string ) {

	viper.AddConfigPath("/etc/commservice//")
	viper.SetConfigName("comconfig")
	viper.ReadInConfig()
	username := viper.GetString("emaildb.username")
	dbpass := viper.GetString("emaildb.password")
	serverip := viper.GetString("emaildb.dbhost")
	mg := mailgun.NewMailgun(domain, apiKey)
	var Email string
	var FirstName string
	var LastName string


	db, err := sql.Open("mysql", username+":"+dbpass+"@tcp("+serverip+")"+"/")
	if err != nil {
		log.Fatal("Can't connect to mysql DATABASE!")
	}
	// I left this pretty open ended. You can query on any field for the first one. I added the city and state as an or meaning you could use these if you want
	rows, err := db.Query("SELECT Email, FirstName, LastName from " + 	DB 	+ " where " + QueryField + " = ? or Town = ? or State = ?", Condition, City, State )
	if err != nil {
		log.Println("SELECT Email, FirstName, LastName from ? where ? = ? or Town = ? or State = ?", DB, QueryField, Condition, City, State)
		log.Fatal("There was an error with your query", err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Email, &FirstName, &LastName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(Email)
		m := mg.NewMessage(
			senderName + " " +  "<" + sender + ">",
			subject,
			message,
			Email,
		)
		m.SetHtml( "<html><h1>" + subject + " Hello " + FirstName + " " + LastName + "</h1><p>" +  message + "<p></html>")
		m.SetTemplate("")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		_, id, err  := mg.Send(ctx, m)
		log.Println(id)



	}


}
