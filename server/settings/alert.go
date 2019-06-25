package settings

import (
	"github.com/cjburchell/tools-go/env"
)

// Alert Settings
var SendOnReminder = env.GetBool("ALERT_REMINDER_ENABLE", false)

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
