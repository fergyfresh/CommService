[![Build Status](https://travis-ci.com/DMEvanCT/CommService.svg?branch=master)](https://travis-ci.com/DMEvanCT/CommService)

# DM Communication Microservice
Used for communications between Application and users





Tool can be moved to /usr/local/bin by sudo cp claritytools /usr/local/bin.  After that you can run claritytools -h from the command line. 

FYI the service queries an auth db for token and username. They are set with the headers.
x-auth-token 
x-auth-user

If you have an issue setting this up reach out and I will post the demo SQL.

####Credentials Files

1. Create the directory /etc/commservice/comconfig.yaml
2. Inside of that directory creat the following file 

comconfig.yaml

```yaml
slack:
  token: (slack token)
  channel: (channel you want to post to by default)
mailgun:
  apiKey: (mailgun api key)
  domain: (mailgun domain)
  sender: (sender ex tech@darkamtterct.com)
  SenderName: (sender name ex: Dark Matter IT Tech)
authdb:
  username: (DB Username)
  password: (DB Password)
  dbhost:   (db ip)
emaildb:
  username: (Email DB Username)
  password: (Email DB password)
  dbhost: (Email DB Host)

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

:8080/api/comm/mailgun/multimessage
```json 
{
	"MailGunData": {
		"Subject": "Testing Multi", 
		"Message": "Here is the message!",
		"To": null, 
		"Service": "MailGunMulti"
	},
	"QueryField": "DatabaseColumn",
	"Condition": "ValueofColumn",
	"City": "null", 
	"State": "null",
	"Database": {
		"Database": "EmailDB.tbl_user_info"
	}
}
```

