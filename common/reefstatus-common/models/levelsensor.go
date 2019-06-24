package models

import (
	"github.com/cjburchell/profilux-go/types"
)

// LevelSensor model
type LevelSensor struct {
	BaseInfo
	Index             int
	AlarmState        types.CurrentState
	OperationMode     types.LevelSensorOperationMode
	Value             types.CurrentState
	SensorIndex       int
	WaterMode         types.WaterMode
	SecondSensor      types.CurrentState
	SecondSensorIndex int
	HasTwoInputs      bool
	HasWaterChange    bool
}

// LevelSensorType type name
const LevelSensorType = "LevelSensor"

// NewLevelSensor creates a new level sensor
func NewLevelSensor(index int) LevelSensor {
	var sensor LevelSensor
	sensor.Index = index
	sensor.Type = LevelSensorType
	sensor.Units = "State"
	sensor.ID = GetID(LevelSensorType, index)
	return sensor
}
