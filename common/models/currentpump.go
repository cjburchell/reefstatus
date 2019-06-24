package models

// CurrentPump Model
type CurrentPump struct {
	BaseInfo
	Value int
	Index int
}

// CurrentPumpType is the type name
const CurrentPumpType = "CurrentPump"

// NewCurrentPump creates a new pump
func NewCurrentPump(index int) CurrentPump {
	var pump CurrentPump
	pump.Type = CurrentPumpType
	pump.Units = "%"
	pump.Index = index
	pump.ID = GetID(CurrentPumpType, index)

	return pump
}
