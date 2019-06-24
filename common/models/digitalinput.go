package models

import (
	"github.com/cjburchell/reefstatus/common/profilux/types"
)

// DigitalInput model
type DigitalInput struct {
	BaseInfo
	Index    int
	Value    types.CurrentState
	Function types.DigitalInputFunction
}

// DigitalInputType name of the type
const DigitalInputType = "DigitalInput"

// NewDigitalInput creates a new input
func NewDigitalInput(index int) DigitalInput {
	var digitalInput DigitalInput
	digitalInput.Index = index
	digitalInput.Type = DigitalInputType
	digitalInput.Units = "State"
	digitalInput.ID = GetID(DigitalInputType, index)

	return digitalInput
}
