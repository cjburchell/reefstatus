package main

import (
	"github.com/cjburchell/go-uatu"
	"github.com/cjburchell/reefstatus/common/communication"
	"github.com/cjburchell/reefstatus/server/alert/check"
	"github.com/cjburchell/reefstatus/server/alert/slack"
	"github.com/cjburchell/reefstatus/server/alert/state"
	"github.com/cjburchell/reefstatus/server/data/repo"
	"github.com/cjburchell/reefstatus/server/settings"
)

func RunAlerts(controller repo.Controller) {
	err := slack.Setup(settings.SlackDestination)
	if err != nil {
		log.Fatal(err, "Unable to connect to slack")
	}

	err = state.Setup(settings.MongoUrl)
	if err != nil {
		log.Fatal(err, "Unable to setup state repo")
	}

	session, err := communication.NewSession(settings.PubSubAddress, settings.PubSubToken)
	if err != nil {
		log.Fatal(err, "Unable to connect to Pub Sub")
	}

	defer session.Close()
	ch, err := session.QueueSubscribe(communication.UpdateAlertsMessage, "alert")
	if err != nil {
		log.Fatal(err, "Unable to connect subscribe to UpdateMessage")
	}

	alertSettings := settings.NewAlertSettings()

	for {
		<-ch

		if alertSettings.SendOnReminder {
			log.Debug("Checking reminders")
			err = check.Reminders(controller)
			if err != nil {
				log.Error(err, "Error sending reminders")
			}
		}

		log.Debug("Checking alarms")
		err = check.Alarm(controller)
		if err != nil {
			log.Error(err, "Error checking alarms")
		}
	}
}
