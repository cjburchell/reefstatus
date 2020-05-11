package commands

import (
	"encoding/json"
	"strconv"

	"github.com/cjburchell/reefstatus/common/models"
	"github.com/cjburchell/reefstatus/controller/profilux"

	"github.com/cjburchell/reefstatus/common/communication"
	"github.com/cjburchell/reefstatus/controller/service"

	"github.com/cjburchell/reefstatus/controller/settings"

	"github.com/cjburchell/reefstatus/controller/update"
	logger "github.com/cjburchell/uatu-go"
)

const queueName = "Controller"

// Handle setup and routing of commands
func Handle(session communication.SubscribeSession, repo service.Controller, log logger.ILog) {
	feedPauseChannel, err := session.QueueSubscribe(communication.FeedPauseMessage, queueName)
	if err != nil {
		log.Errorf(err, "subscibe to %s", communication.FeedPauseMessage)
		return
	}

	thunderstormChannel, err := session.QueueSubscribe(communication.ThunderstormMessage, queueName)
	if err != nil {
		log.Errorf(err, "subscibe to %s", communication.ThunderstormMessage)
		return
	}

	resetReminderChannel, err := session.QueueSubscribe(communication.ResetReminderMessage, queueName)
	if err != nil {
		log.Errorf(err, "subscibe to %s", communication.ResetReminderMessage)
		return
	}

	maintenanceChannel, err := session.QueueSubscribe(communication.MaintenanceMessage, queueName)
	if err != nil {
		log.Errorf(err, "subscibe to %s", communication.MaintenanceMessage)
		return
	}

	clearLevelChannel, err := session.QueueSubscribe(communication.ClearLevelAlarmMessage, queueName)
	if err != nil {
		log.Errorf(err, "subscibe to %s", communication.ClearLevelAlarmMessage)
		return
	}

	waterChangeChannel, err := session.QueueSubscribe(communication.WaterChangeMessage, queueName)
	if err != nil {
		log.Errorf(err, "subscibe to %s", communication.WaterChangeMessage)
		return
	}

	for {
		var result string
		select {
		case result = <-feedPauseChannel:
			enabled, _ := strconv.ParseBool(result)
			feedPause(enabled, repo, log)
		case result = <-thunderstormChannel:
			duration, _ := strconv.Atoi(result)
			thunderstorm(duration, repo, log)
		case result = <-resetReminderChannel:
			index, _ := strconv.Atoi(result)
			resetReminder(index, repo, log)
		case result = <-maintenanceChannel:
			message := struct {
				index  int
				enable bool
			}{}
			_ = json.Unmarshal([]byte(result), &message)
			maintenance(message.index, message.enable, repo, log)
		case result = <-clearLevelChannel:
			clearLevelAlarm(result, repo, log)
		case result = <-waterChangeChannel:
			waterChange(result, repo, log)
		}
	}

}
func feedPause(bool bool, repo service.Controller, log logger.ILog) {
	profiluxController, err := profilux.NewController(settings.Connection)
	if err != nil {
		log.Errorf(err, "unable to connect")
		return
	}

	defer profiluxController.Disconnect()

	err = profiluxController.FeedPause(0, bool)
	if err != nil {
		log.Errorf(err, "Unable to send feed pause")
		return
	}

	err = update.InfoState(profiluxController, repo)
	if err != nil {
		log.Errorf(err, "Unable update state")
		return
	}
}
func thunderstorm(duration int, repo service.Controller, log logger.ILog) {
	profiluxController, err := profilux.NewController(settings.Connection)
	if err != nil {
		log.Errorf(err, "Unable to connect")
		return
	}

	defer profiluxController.Disconnect()

	err = profiluxController.Thunderstorm(duration)
	if err != nil {
		log.Errorf(err, "Unable to send start thunderstorm")
		return
	}

	err = update.InfoState(profiluxController, repo)
	if err != nil {
		log.Errorf(err, "Unable update state")
		return
	}
}
func resetReminder(index int, repo service.Controller, log logger.ILog) {

	var reminder *models.Reminder

	info, err := repo.GetInfo()
	if err != nil {
		log.Errorf(err, "Unable to get info")
		return
	}

	for _, item := range info.Reminders {
		if item.Index == index {
			reminder = &item
			break
		}
	}

	if reminder == nil {
		log.Warnf("unable to find reminder")
		return
	}

	profiluxController, err := profilux.NewController(settings.Connection)
	if err != nil {
		log.Errorf(err, "unable to connect")
		return
	}

	defer profiluxController.Disconnect()

	if reminder.IsRepeating {
		err = profiluxController.ResetReminder(index, reminder.Period)
		if err != nil {
			log.Errorf(err, "Unable Reset Reminder")
			return
		}

	} else {
		err = profiluxController.ClearReminder(index)
		if err != nil {
			log.Errorf(err, "Unable Clear Reminder")
			return
		}
	}

	err = update.InfoState(profiluxController, repo)
	if err != nil {
		log.Errorf(err, "Unable update state")
		return
	}
}
func maintenance(index int, enable bool, repo service.Controller, log logger.ILog) {
	profiluxController, err := profilux.NewController(settings.Connection)
	if err != nil {
		log.Errorf(err, "unable to connect")
		return
	}

	defer profiluxController.Disconnect()

	var maintenance *models.Maintenance
	info, err := repo.GetInfo()
	if err != nil {
		log.Errorf(err, "Unable to get info")
		return
	}

	for _, item := range info.Maintenance {
		if item.Index == index {
			maintenance = &item
			break
		}
	}

	if maintenance == nil {
		log.Warnf("unable to find reminder")
		return
	}

	err = profiluxController.Maintenance(enable, index)
	if err != nil {
		log.Errorf(err, "Unable to set Maintenance")
		return
	}

	err = update.InfoState(profiluxController, repo)
	if err != nil {
		log.Errorf(err, "Unable update state")
		return
	}
}
func clearLevelAlarm(id string, repo service.Controller, log logger.ILog) {

	var sensor *models.LevelSensor
	items, err := repo.GetLevelSensors()
	if err != nil {
		log.Warnf("unable to find level sensor %s", id)
		return
	}

	for _, level := range items {
		if level.ID == id {
			sensor = &level
			break
		}
	}

	if sensor == nil {
		log.Warnf("unable to find level sensor %s", id)
		return
	}

	profiluxController, err := profilux.NewController(settings.Connection)
	if err != nil {
		log.Errorf(err, "unable to connect")
	}

	defer profiluxController.Disconnect()

	err = profiluxController.ClearLevelAlarm(sensor.Index)
	if err != nil {
		log.Errorf(err, "Unable Clear Level Alarm")
		return
	}

	err = update.LevelSensors(profiluxController, repo)
	if err != nil {
		log.Errorf(err, "Unable to update Level Sensors")
		return
	}

	err = update.InfoState(profiluxController, repo)
	if err != nil {
		log.Errorf(err, "Unable update state")
		return
	}
}
func waterChange(id string, repo service.Controller, log logger.ILog) {
	var sensor *models.LevelSensor

	sensors, err := repo.GetLevelSensors()
	if err != nil {
		log.Errorf(err, "Unable to get level sensors")
		return
	}

	for _, level := range sensors {
		if level.ID == id {
			sensor = &level
			break
		}
	}

	if sensor == nil {
		log.Warnf("unable to find level snsoro %s", id)
		return
	}

	profiluxController, err := profilux.NewController(settings.Connection)
	if err != nil {
		log.Errorf(err, "unable to connect")
	}

	defer profiluxController.Disconnect()

	err = profiluxController.WaterChange(sensor.Index)
	if err != nil {
		log.Errorf(err, "Unable to do Water Change")
		return
	}

	err = update.LevelSensors(profiluxController, repo)
	if err != nil {
		log.Errorf(err, "Unable to update Level Sensors")
		return
	}

	err = update.InfoState(profiluxController, repo)
	if err != nil {
		log.Errorf(err, "Unable update state")
		return
	}
}
