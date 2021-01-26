package alert

import (
	"github.com/cjburchell/reefstatus/server/alert/check"
	"github.com/cjburchell/reefstatus/server/alert/state"
	"github.com/cjburchell/reefstatus/server/data/repo"
	logger "github.com/cjburchell/uatu-go"
)

func Check(controller repo.Controller, stateRepo state.StateData, log logger.ILog, sendOnReminder bool, slackDestination string) {
	if sendOnReminder {
		log.Debug("Checking reminders")
		err := check.Reminders(controller, stateRepo, slackDestination)
		if err != nil {
			log.Error(err, "Error sending reminders")
		}
	}

	log.Debug("Checking alarms")
	err := check.Alarm(controller, stateRepo, slackDestination)
	if err != nil {
		log.Error(err, "Error checking alarms")
	}

}
