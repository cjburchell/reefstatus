package alert

import (
	"github.com/cjburchell/go-uatu"
	"github.com/cjburchell/reefstatus/server/alert/check"
	"github.com/cjburchell/reefstatus/server/alert/state"
	"github.com/cjburchell/reefstatus/server/data/repo"
	"github.com/cjburchell/reefstatus/server/settings"
)

func Check(controller repo.Controller, stateRepo state.StateData) {
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
