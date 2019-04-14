# DM Communication Microservice
Used for communications between Application and users





Tool can be moved to /usr/local/bin by sudo cp claritytools /usr/local/bin.  After that you can run claritytools -h from the command line. 


####Credentials Files

1. Create the directory /etc/commservice/comconfig.yaml
2. Inside of that directory creat the following file 

credentials.yaml

```yaml
slack:
  token: (slack token)
  channel: (channel you want to post to by default)
mailgun:
  apiKey: (mailgun api key)
  domain: (mailgun domain)
  sender: (sender ex tech@darkamtterct.com)
  SenderName: (sender name ex: Dark Matter IT Tech)
```

The following are the current post to URL's and the json to go with it. 

:8080/api/comm/slack/message

```json 
{
"SlackMessage": "Hello All",
"Service": "slack"
}
```

:8080/api/comm/mailgun/message
```json
{
"Subject": "Sale at Moxie!",
"Message": "Hello All",
"To": "evan.haston@darkmatterct.com",
"Service": "MailGun"
}
```

