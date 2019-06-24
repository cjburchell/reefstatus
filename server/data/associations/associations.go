package associations

import (
	"github.com/cjburchell/reefstatus-common/models"
	"github.com/cjburchell/reefstatus-data/repo"

	"github.com/cjburchell/profilux-go/types"
)

func getAssociatedModeItem(mode types.PortMode, repo repo.Controller) string {
	index := mode.Port
	if mode.IsProbe {
		probe, err := repo.GetProbe(models.GetID(models.ProbeType, index))
		if err == nil {
			return probe.ID
		}
	}

	switch mode.DeviceMode {
	case types.DeviceModeLights:
		light, err := repo.GetLight(models.GetID(models.LightType, index))
		if err == nil {
			return light.ID
		}

	case types.DeviceModeTimer:
		timer, err := repo.GetDosingPump(models.GetID(models.DosingPumpType, index))
		if err == nil {
			return timer.ID
		}
	case types.DeviceModeWater:
		level, err := repo.GetLevelSensor(models.GetID(models.LevelSensorType, index))
		if err == nil {
			return level.ID
		}

	case types.DeviceModeDrainWater:
		level, err := repo.GetLevelSensor(models.GetID(models.LevelSensorType, index))
		if err == nil {
			return level.ID
		}

	case types.DeviceModeWaterChange:
		level, err := repo.GetLevelSensor(models.GetID(models.LevelSensorType, index))
		if err == nil {
			return level.ID
		}

	case types.DeviceModeCurrentPump:
		pump, err := repo.GetCurrentPump(models.GetID(models.CurrentPumpType, index))
		if err == nil {
			return pump.ID
		}

	case types.DeviceModeProgrammableLogic:
		logic, err := repo.GetProgrammableLogic(models.GetID(models.ProgrammableLogicType, index))
		if err == nil {
			return logic.Id
		}
	}

	return ""
}

// Update the associations
func Update(repo repo.Controller) {

	logicItems, _ := repo.GetProgrammableLogics()
	for _, logic := range logicItems {
		logic.Input1.Id = getAssociatedModeItem(logic.Input1, repo)
		logic.Input1.Id = getAssociatedModeItem(logic.Input2, repo)
		repo.SetProgrammableLogic(logic)
	}

	sPorts, _ := repo.GetSPorts()
	for _, port := range sPorts {
		port.Mode.Id = getAssociatedModeItem(port.Mode, repo)
		repo.SetSPort(port)
	}

	lPorts, _ := repo.GetLPorts()
	for _, port := range lPorts {
		port.Mode.Id = getAssociatedModeItem(port.Mode, repo)
		repo.SetLPort(port)
	}
}
