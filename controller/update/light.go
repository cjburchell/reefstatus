package update

import (
	"github.com/cjburchell/reefstatus/common/models"
	"github.com/cjburchell/reefstatus/controller/data"
	"github.com/cjburchell/reefstatus/controller/profilux"
)

func lightUpdate(light *models.Light, controller *profilux.Controller) error {
	var err error
	light.IsDimmable, err = controller.GetIsLightDimmable(light.Channel)
	if err != nil {
		return err
	}

	light.OperationHours, err = controller.GetLightOperationHours(light.Channel)
	if err != nil {
		return err
	}

	light.Value, err = controller.GetLightValue(light.Channel)
	if err != nil {
		return err
	}

	light.IsLightOn = light.Value != 0
	light.DisplayName, err = controller.GetLightName(light.Channel)
	return err
}

func lightUpdateState(light *models.Light, controller *profilux.Controller) error {
	var err error
	light.OperationHours, err = controller.GetLightOperationHours(light.Channel)
	if err != nil {
		return err
	}

	light.Value, err = controller.GetLightValue(light.Channel)
	if err != nil {
		return err
	}

	light.IsLightOn = light.Value != 0

	return nil
}

func lights(profiluxController *profilux.Controller, repo data.ControllerService) error {
	count, err := profiluxController.GetLightCount()
	if err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		light, err := repo.GetLight(models.GetID(models.LightType, i))

		if err != nil && err != data.ErrNotFound {
			return err
		}
		found := err != data.ErrNotFound
		isActive, err := profiluxController.GetIsLightActive(i)
		if err != nil {
			return err
		}

		if isActive {
			if !found {
				light = models.NewLight(i)
			}

			err = lightUpdate(&light, profiluxController)
			if err != nil {
				return err
			}

			err = repo.SetLight(light, !found)
			if err != nil {
				return err
			}

		} else {
			if found {
				err = repo.DeleteLight(light)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
