package update

import (
	"github.com/cjburchell/reefstatus/common/data"
	"github.com/cjburchell/reefstatus/common/models"
	"github.com/cjburchell/reefstatus/common/profilux/types"
	"github.com/cjburchell/reefstatus/controller/profilux"
)

func programmableLogicUpdate(logic *models.ProgrammableLogic, controller *profilux.Controller) error {
	var err error
	logic.Input1, err = controller.GetProgramLogicInput(0, logic.Index)
	if err != nil {
		return err
	}
	logic.Input2, err = controller.GetProgramLogicInput(1, logic.Index)
	if err != nil {
		return err
	}
	logic.Function, err = controller.GetProgramLogicFunction(logic.Index)
	return err
}

func programmableLogic(profiluxController *profilux.Controller, repo data.ControllerService) error {
	count, err := profiluxController.GetProgrammableLogicCount()
	if err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		logic, err := repo.GetProgrammableLogic(models.GetID(models.ProgrammableLogicType, i))
		if err != nil && err != data.ErrNotFound {
			return err
		}

		found := err != data.ErrNotFound

		input1, err := profiluxController.GetProgramLogicInput(0, i)
		if err != nil {
			return err
		}

		input2, err := profiluxController.GetProgramLogicInput(1, i)
		if err != nil {
			return err
		}

		if input1.DeviceMode != types.DeviceModeAlwaysOff && input2.DeviceMode != types.DeviceModeAlwaysOff {
			if !found {
				logic = models.NewProgrammableLogic(i)
			}

			err = programmableLogicUpdate(&logic, profiluxController)
			if err != nil {
				return err
			}

			err = repo.SetProgrammableLogic(logic, !found)
			if err != nil {
				return err
			}

		} else {
			if found {
				err = repo.DeleteProgrammableLogic(logic)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
