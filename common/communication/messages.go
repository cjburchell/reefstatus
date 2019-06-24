package communication

import (
	"encoding/json"
	"strconv"
)

// UpdateMessage name
const UpdateHistoryMessage = "UpdateHistory"
const UpdateAlertsMessage = "UpdateAlerts"

// FeedPauseMessage name
const FeedPauseMessage = "FeedPause"

// ThunderstormMessage name
const ThunderstormMessage = "Thunderstorm"

// ResetReminderMessage name
const ResetReminderMessage = "ResetReminder"

// MaintenanceMessage name
const MaintenanceMessage = "Maintenance"

// ClearLevelAlarmMessage name
const ClearLevelAlarmMessage = "ClearLevelAlarm"

// WaterChangeMessage name
const WaterChangeMessage = "WaterChange"

// Update command
func Update(session Session, isInitial bool) error {
	err := session.Publish(UpdateHistoryMessage, strconv.FormatBool(isInitial))
	if err != nil {
		return err
	}

	return session.Publish(UpdateAlertsMessage, strconv.FormatBool(isInitial))
}

// FeedPause command
func FeedPause(session Session, enable bool) error {
	return session.Publish(FeedPauseMessage, strconv.FormatBool(enable))
}

// Thunderstorm command
func Thunderstorm(session Session, duration int) error {
	return session.Publish(ThunderstormMessage, strconv.Itoa(duration))
}

// ResetReminder command
func ResetReminder(session Session, index int) error {
	return session.Publish(ResetReminderMessage, strconv.Itoa(index))
}

// Maintenance command
func Maintenance(session Session, index int, enable bool) error {
	data, _ := json.Marshal(struct {
		index  int
		enable bool
	}{index: index, enable: enable})

	return session.PublishData(MaintenanceMessage, data)
}

// ClearLevelAlarm command
func ClearLevelAlarm(session Session, id string) error {
	return session.Publish(ClearLevelAlarmMessage, id)
}

// WaterChange command
func WaterChange(session Session, id string) error {
	return session.Publish(WaterChangeMessage, id)
}
