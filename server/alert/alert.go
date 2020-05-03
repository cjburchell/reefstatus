package alert

import (
	"github.com/cjburchell/reefstatus/server/alert/check"
	"github.com/cjburchell/reefstatus/server/alert/state"
	"github.com/cjburchell/reefstatus/server/data/repo"
	"github.com/cjburchell/reefstatus/server/settings"
	logger "github.com/cjburchell/uatu-go"
)

func Check(controller repo.Controller, stateRepo state.StateData, log logger.ILog) {
	if settings.SendOnReminder {
		log.Debug("Checking reminders")
		err := check.Reminders(controller, stateRepo)
		if err != nil {
			log.Error(err, "Error sending reminders")
		}
	}

	log.Debug("Checking alarms")
	err := check.Alarm(controller, stateRepo)
	if err != nil {
		log.Error(err, "Error checking alarms")
	}

}
