package check

import (
	"fmt"
	"time"

	"github.com/cjburchell/reefstatus/server/data/repo"

	"github.com/cjburchell/reefstatus/common/profilux/types"
	"github.com/cjburchell/reefstatus/server/alert/slack"

	"github.com/cjburchell/reefstatus/server/alert/state"
)

// Alarm check
func Alarm(controller repo.Controller) error {
	info, err := controller.GetInfo()
	if err != nil {
		return err
	}

	if info.Alarm == types.CurrentStateOn {

		updated, err := state.UpdateAlarmSent(true)
		if err != nil {
			return err
		}

		if updated {
			return sendAlarmEmail(controller, false)
		}
	} else {
		updated, err := state.UpdateAlarmSent(false)
		if err != nil {
			return err
		}
		if updated {
			return sendAlarmEmail(controller, true)
		}
	}

	return nil
}

func sendAlarmEmail(controller repo.Controller, cleared bool) error {
	statusTable, err := createStatusTable(controller)
	if err != nil {
		return err
	}

	if !cleared {
		reason, err := findAlarmReason(controller)
		if err != nil {
			return err
		}

		message := fmt.Sprintf("Alarm Detected %s\nReasons\n%s\n%s", time.Now().Format("2006-01-02 15:04:05 MST"), reason, statusTable)
		err = slack.PrintMessage(message)
		if err != nil {
			return err
		}

	} else {
		err = slack.PrintMessage(fmt.Sprintf("Alarm Cleared %s\n%s", time.Now().Format("2006-01-02 15:04:05 MST"), statusTable))
		if err != nil {
			return err
		}
	}

	return nil
}

func findAlarmReason(controller repo.Controller) (reason string, err error) {
	probes, err := controller.GetProbes()
	if err != nil {
		return
	}

	for _, probe := range probes {
		if probe.AlarmState == types.CurrentStateOn && probe.AlarmEnable {
			if probe.Value > probe.NominalValue+probe.AlarmDeviation {
				reason += fmt.Sprintf("%s is too high\n", probe.DisplayName)
			} else if probe.Value < probe.NominalValue-probe.AlarmDeviation {
				reason += fmt.Sprintf("%s is too low\n", probe.DisplayName)
			} else {
				reason += fmt.Sprintf("Alarm on %s\n", probe.DisplayName)
			}
		}
	}

	sensors, err := controller.GetLevelSensors()
	if err != nil {
		return
	}

	for _, sensor := range sensors {
		if sensor.AlarmState != types.CurrentStateOn {
			continue
		}

		reason += fmt.Sprintf("Level Timeout %s\n", sensor.DisplayName)
	}

	if len(reason) == 0 {
		reason = "Unknown"
	}

	return
}

func createStatusTable(controller repo.Controller) (table string, err error) {
	probes, err := controller.GetProbes()
	if err != nil {
		return
	}

	for _, probe := range probes {
		table += fmt.Sprintf("%s\t%f%s\n", probe.DisplayName, probe.ConvertedValue, probe.Units)
	}

	sensors, err := controller.GetLevelSensors()
	if err != nil {
		return
	}

	for _, sensor := range sensors {
		table += fmt.Sprintf("%s\t%s\t%s\n", sensor.DisplayName, sensor.OperationMode, sensor.Value)
	}

	return
}
