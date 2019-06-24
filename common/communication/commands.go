package communication

import (
	"encoding/json"
	"strconv"
)

// Update command
func Update(session PublishSession, isInitial bool) error {
	err := session.Publish(UpdateHistoryMessage, strconv.FormatBool(isInitial))
	if err != nil {
		return err
	}

	return session.Publish(UpdateAlertsMessage, strconv.FormatBool(isInitial))
}

// FeedPause command
func FeedPause(session PublishSession, enable bool) error {
	return session.Publish(FeedPauseMessage, strconv.FormatBool(enable))
}

// Thunderstorm command
func Thunderstorm(session PublishSession, duration int) error {
	return session.Publish(ThunderstormMessage, strconv.Itoa(duration))
}

// ResetReminder command
func ResetReminder(session PublishSession, index int) error {
	return session.Publish(ResetReminderMessage, strconv.Itoa(index))
}

// Maintenance command
func Maintenance(session PublishSession, index int, enable bool) error {
	data, _ := json.Marshal(struct {
		index  int
		enable bool
	}{index: index, enable: enable})

	return session.PublishData(MaintenanceMessage, data)
}

// ClearLevelAlarm command
func ClearLevelAlarm(session PublishSession, id string) error {
	return session.Publish(ClearLevelAlarmMessage, id)
}

// WaterChange command
func WaterChange(session PublishSession, id string) error {
	return session.Publish(WaterChangeMessage, id)
}
