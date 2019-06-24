package update

import (
	"github.com/cjburchell/reefstatus/common/models"
	"github.com/cjburchell/reefstatus/common/profilux/types"
	"github.com/cjburchell/reefstatus/controller/data"
	"github.com/cjburchell/reefstatus/controller/profilux"
)

func digitalInputs(controller *profilux.Controller, repo data.ControllerService) error {
	count, err := controller.GetDigitalInputCount()
	if err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		mode, err := controller.GetDigitalInputFunction(i)
		if err != nil {
			return err
		}

		sensor, err := repo.GetDigitalInput(models.GetID(models.DigitalInputType, i))
		if err != nil && err != data.ErrNotFound {
			return err
		}

		found := err != data.ErrNotFound
		if mode != types.DigitalInputFunctionNotUsed {
			if !found {
				sensor = models.NewDigitalInput(i)
			}

			sensor.Function, err = controller.GetDigitalInputFunction(sensor.Index)
			if err != nil {
				return err
			}

			sensor.Value, err = controller.GetDigitalInputState(sensor.Index)
			if err != nil {
				return err
			}

			err = repo.SetDigitalInput(sensor, !found)
			if err != nil {
				return err
			}

		} else {
			if found {
				err = repo.DeleteDigitalInput(sensor)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
