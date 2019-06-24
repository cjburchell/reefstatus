package models

import (
	"github.com/cjburchell/profilux-go/types"
)

type LPort struct {
	BaseInfo
	PortNumber int
	Mode       types.PortMode
	Value      float64
}

const LPortType = "LPort"

func NewLPort(portNumber int) LPort {
	var lPort LPort
	lPort.Type = "LPort"
	lPort.Units = "%"
	lPort.PortNumber = portNumber
	lPort.ID = GetID(LPortType, portNumber)
	return lPort
}
