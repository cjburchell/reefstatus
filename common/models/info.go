package models

import (
	"time"

	"github.com/cjburchell/reefstatus/common/profilux/types"
)

// Info model
type Info struct {
	Maintenance     []Maintenance
	OperationMode   types.OperationMode
	Model           types.Model
	SoftwareDate    time.Time
	DeviceAddress   int
	Latitude        float64
	Longitude       float64
	MoonPhase       float64
	Alarm           types.CurrentState
	SoftwareVersion float64
	SerialNumber    int
	LastUpdate      time.Time
	Reminders       []Reminder
}

// NewInfo creates new object
func NewInfo() Info {
	var info Info
	info.Maintenance = make([]Maintenance, 0)
	info.Reminders = make([]Reminder, 0)

	return info
}

// IsP3 checks to see if the controller is a p3
func (info Info) IsP3() bool {
	return info.Model == types.ProfiLuxIII || info.Model == types.ProfiLuxIIIEx
}
