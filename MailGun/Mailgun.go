package MailGunz

import (
	"context"
	"time"
	"github.com/mailgun/mailgun-go"
)





func SendMailGunEmail(domain, apiKey, subject, message, to, sender, senderName string) (string, error) {
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

	_, id, err := mg.Send(ctx, m)
	return id, err
}
