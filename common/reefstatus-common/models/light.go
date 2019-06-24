package models

type Light struct {
	BaseInfo
	Value          float64
	Channel        int
	IsDimmable     bool
	OperationHours int
	IsLightOn      bool
}

const LightType = "Light"

func NewLight(index int) Light {
	var light Light
	light.Channel = index
	light.Type = LightType
	light.Units = "%"
	light.ID = GetID(LightType, index)
	return light
}
