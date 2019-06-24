package models

import (
	"time"
)

// Reminder model
type Reminder struct {
	IsOverdue   bool
	Next        time.Time
	Text        string
	Index       int
	Period      int
	IsRepeating bool
}

// NewReminder creates new object
func NewReminder(index int) *Reminder {
	var reminder Reminder
	reminder.Index = index
	return &reminder
}
