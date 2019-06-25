package commands

import (
	"encoding/json"
	"strconv"

	"github.com/cjburchell/reefstatus/common/communication"
)

// FeedPause command
func FeedPause(session communication.PublishSession, enable bool) error {
	return session.Publish(communication.FeedPauseMessage, strconv.FormatBool(enable))
}

// Thunderstorm command
func Thunderstorm(session communication.PublishSession, duration int) error {
	return session.Publish(communication.ThunderstormMessage, strconv.Itoa(duration))
}

// ResetReminder command
func ResetReminder(session communication.PublishSession, index int) error {
	return session.Publish(communication.ResetReminderMessage, strconv.Itoa(index))
}

// Maintenance command
func Maintenance(session communication.PublishSession, index int, enable bool) error {
	data, _ := json.Marshal(struct {
		index  int
		enable bool
	}{index: index, enable: enable})

	return session.PublishData(communication.MaintenanceMessage, data)
}

// ClearLevelAlarm command
func ClearLevelAlarm(session communication.PublishSession, id string) error {
	return session.Publish(communication.ClearLevelAlarmMessage, id)
}

// WaterChange command
func WaterChange(session communication.PublishSession, id string) error {
	return session.Publish(communication.WaterChangeMessage, id)
}
