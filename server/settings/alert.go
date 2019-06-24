package settings

import (
	"github.com/cjburchell/tools-go/env"
)

// Alert Settings
type Alert struct {
	SendOnReminder bool
}

// NewAlertSettings initializes the alert settings
func NewAlertSettings() (settings Alert) {
	settings.SendOnReminder = env.GetBool("ALERT_REMINDER_ENABLE", false)
	return
}

// Mail Settings
type MailSettings struct {
	UserName string
	Password string
	Server   string
	To       string
	From     string
}

// NewMailSettings initializes the mail settings
func newMailSettings() (settings MailSettings) {
	settings.UserName = env.Get("ALERT_MAIL_USERNAME", "reefstatusalert")
	settings.Password = env.Get("ALERT_MAIL_PASSWORD", "")
	settings.Server = env.Get("ALERT_MAIL_SERVER", "smtp.gmail.com:587")
	settings.To = env.Get("ALERT_MAIL_To", "cjburchell@yahoo.com")
	settings.From = env.Get("ALERT_MAIL_From", "reefstatusalert@gmail.com")
	return
}

var Mail = newMailSettings()
var PubSubAddress = env.Get("PUB_SUB_ADDRESS", "tcp://localhost:4222")
var PubSubToken = env.Get("PUB_SUB_TOKEN", "token")
var SlackDestination = env.Get("SLACK_DESTINATION", "")
var MongoUrl = env.Get("MONGO_URL", "localhost")
var DataServiceAddress = env.Get("DATA_SERVICE", "http://localhost")
var DataServiceToken = env.Get("DATA_SERVICE_TOKEN", "token")
