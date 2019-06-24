package models

import (
	"github.com/cjburchell/profilux-go/types"
	"github.com/cjburchell/tools-go/math"
)

type Probe struct {
	BaseInfo
	Format         int
	SensorType     types.SensorType
	Index          int
	AlarmState     types.CurrentState
	NominalValue   float64
	SensorMode     types.SensorMode
	AlarmEnable    bool
	AlarmDeviation float64
	Value          float64
	OperationHours int
	ConvertedValue float64
	CenterValue    float64
	MaxRange       float64
	MinRange       float64
}

func (probe *Probe) SetValue(value float64) {
	probe.Value = value
	probe.ConvertedValue = probe.convertValue(value)
}

func (probe *Probe) SetNominalValue(value float64) {
	probe.NominalValue = value
	probe.CenterValue = probe.convertValue(probe.NominalValue)
	probe.MaxRange = probe.convertValue(probe.NominalValue + probe.AlarmDeviation)
	probe.MinRange = probe.convertValue(probe.NominalValue - probe.AlarmDeviation)
}

func (probe *Probe) SetAlarmDeviation(value float64) {
	probe.AlarmDeviation = value
	probe.MaxRange = probe.convertValue(probe.NominalValue + probe.AlarmDeviation)
	probe.MinRange = probe.convertValue(probe.NominalValue - probe.AlarmDeviation)
}

func (probe Probe) GetUnits() string {
	switch probe.SensorType {
	case types.SensorTypePH:
		return "PH"
	case types.SensorTypeAirTemperature:
		if probe.Format == 1 {
			return "°F"
		} else {
			return "°C"
		}
	case types.SensorTypeTemperature:
		if probe.Format == 1 {
			return "°F"
		} else {
			return "°C"
		}
	case types.SensorTypeConductivityF:
		return "μS"
	case types.SensorTypeConductivity:
		if probe.Format == 1 {
			return "ppt/PSU"
		} else if probe.Format == 2 {
			return "SG"
		} else {
			return "mS"
		}
	case types.SensorTypeRedox:
		return "mV"
	case types.SensorTypeOxygen:
		return "%"
	case types.SensorTypeHumidity:
		return "%"
	case types.SensorTypeVoltage:
		return "V"
	}

	return ""
}

func (probe Probe) convertValue(value float64) float64 {
	digits := probe.getDigits()
	switch probe.SensorType {
	case types.SensorTypeAirTemperature:
		if probe.Format == 1 {
			// convert temperature to Fahrenheit
			return math.Round(1.8*value+32.0, digits)
		}
	case types.SensorTypeTemperature:
		if probe.Format == 1 {
			// convert temperature to Fahrenheit
			return math.Round(1.8*value+32.0, digits)
		}
	case types.SensorTypeConductivity:
		if probe.Format == 1 {
			return math.Round(convertToSalinity(value), digits)
		}

		if probe.Format == 2 {
			return math.Round(convertToSg(value, false), digits)
		}
	}

	return math.Round(value, digits)
}

func (probe Probe) GetValueMultiplier() float64 {
	switch probe.SensorType {
	case types.SensorTypePH:
		return 0.01
	case types.SensorTypeAirTemperature:
		return 0.1
	case types.SensorTypeTemperature:
		return 0.1
	case types.SensorTypeConductivityF:
		return 1
	case types.SensorTypeConductivity:
		return 0.1
	case types.SensorTypeOxygen:
		return 0.1
	case types.SensorTypeHumidity:
		return 0.1
	}
	return 1
}

func (probe Probe) getDigits() int {
	switch probe.SensorType {
	case types.SensorTypePH:
		return 2
	case types.SensorTypeAirTemperature:
		return 2
	case types.SensorTypeTemperature:
		return 2
	case types.SensorTypeConductivityF:
		return 2
	case types.SensorTypeOxygen:
		return 2
	case types.SensorTypeHumidity:
		return 2
	case types.SensorTypeConductivity:
		if probe.Format == 1 {
			return 2
		}

		if probe.Format == 2 {
			return 4
		}

		return 2
	}

	return 0
}

func convertToSalinity(cond float64) float64 {
	conversionTable := map[float64]float64{
		40:   25.5,
		40.5: 25.9,
		41:   26.2,
		41.5: 26.6,
		42:   26.9,
		42.5: 27.3,
		43:   27.7,
		43.5: 28,
		44:   28.4,
		44.5: 28.7,
		45:   29.1,
		45.5: 29.5,
		46:   29.8,
		46.5: 30.2,
		47:   30.5,
		47.5: 30.9,
		48:   31.3,
		48.5: 31.6,
		49:   32,
		49.5: 32.4,
		50:   32.7,
		50.5: 33.1,
		51:   33.5,
		51.5: 33.8,
		52:   34.2,
		52.5: 34.6,
		53:   34.9,
		53.5: 35.3,
		54:   35.7,
		54.5: 36.1,
		55:   36.4,
		55.5: 36.8,
		56:   37.2,
		56.5: 37.6,
		57:   37.9,
		57.5: 38.3,
		58:   38.7,
		58.5: 39.1,
		59:   39.6,
		59.5: 39.8,
		60:   40.2,
	}

	var salinity float64
	for key := range conversionTable {
		if key >= cond {
			salinity = cond * (conversionTable[key] / key)
			break
		}
	}

	if salinity == 0 {
		salinity = cond * (conversionTable[60.0] / 60.0)
	}

	return math.Round(salinity, 1)
}

func convertToSg(cond float64, offset bool) float64 {
	conversionTable := map[float64]float64{
		40:   1.0187,
		40.5: 1.019,
		41:   1.0193,
		41.5: 1.0195,
		42:   1.0198,
		42.5: 1.0201,
		43:   1.0204,
		43.5: 1.0206,
		44:   1.0209,
		44.5: 1.0212,
		45:   1.0214,
		45.5: 1.0217,
		46:   1.022,
		46.5: 1.0223,
		47:   1.0225,
		47.5: 1.0228,
		48:   1.0231,
		48.5: 1.0234,
		49:   1.0236,
		49.5: 1.0239,
		50:   1.0242,
		50.5: 1.0245,
		51:   1.0248,
		51.5: 1.025,
		52:   1.0253,
		52.5: 1.0256,
		53:   1.0259,
		53.5: 1.0262,
		54:   1.0264,
		54.5: 1.0267,
		55:   1.027,
		55.5: 1.0273,
		56:   1.0276,
		56.5: 1.0278,
		57:   1.0281,
		57.5: 1.0284,
		58:   1.0287,
		58.5: 1.029,
		59:   1.0293,
		59.5: 1.0296,
		60:   1.0299,
	}

	var sg float64
	for key := range conversionTable {
		if key >= cond {
			sg = cond * ((conversionTable[key] - 1.0) / key)
			break
		}
	}

	if sg == 0 {
		sg = cond * ((conversionTable[60] - 1.0) / 60.0)
	}

	if offset {
		return math.Round(sg, 4)
	}
	return math.Round(sg, 4) + 1.0

}

const ProbeType = "Probe"

func NewProbe(index int) Probe {
	var probe Probe
	probe.Type = ProbeType
	probe.Index = index
	probe.ID = GetID(ProbeType, index)
	return probe
}
