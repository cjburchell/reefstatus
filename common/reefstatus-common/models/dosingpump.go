package models

import (
	"github.com/cjburchell/profilux-go/types"
)

// DosingPump creates a new item
type DosingPump struct {
	BaseInfo
	Channel  int
	Rate     int
	PerDay   int
	Settings types.TimerSettings
}

// DosingPumpType name
const DosingPumpType = "Dosing"

// NewDosingPump creates a new pump
func NewDosingPump(index int) DosingPump {
	var pump DosingPump
	pump.Channel = index
	pump.Type = DosingPumpType
	pump.Units = "ml/day"
	pump.ID = GetID(DosingPumpType, index)
	return pump
}
