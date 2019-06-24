package models

import (
	"github.com/cjburchell/reefstatus/common/profilux/types"
)

// SPort model
type SPort struct {
	BaseInfo
	PortNumber         int
	Mode               types.PortMode
	Value              types.CurrentState
	CurrentColourValue int
	IsActive           bool
	Current            float64
}

// SPortType name
const SPortType = "SPort"

// NewSPort creates new object
func NewSPort(index int) SPort {
	var probe SPort
	probe.Type = SPortType
	probe.Units = "State"
	probe.PortNumber = index
	probe.ID = GetID(SPortType, index)
	return probe
}
