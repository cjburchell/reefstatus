package email

import (
	"fmt"
	"net"
	"net/mail"
	"net/smtp"

	"github.com/pkg/errors"

	"github.com/cjburchell/reefstatus-alert/settings"
)

// Send a email
func Send(subject, body string, settings settings.MailSettings) error {
	from := mail.Address{Name: "Reef Status", Address: settings.From}
	to := mail.Address{Name: "", Address: settings.To}
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	serverName := settings.Server
	host, _, err := net.SplitHostPort(serverName)
	if err != nil {
		return errors.WithStack(err)
	}

	auth := smtp.PlainAuth("", settings.UserName, settings.Password, host)

	err = smtp.SendMail(serverName, auth, from.Address, []string{to.Address}, []byte(message))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
