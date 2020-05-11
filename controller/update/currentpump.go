package update

import (
	"github.com/cjburchell/reefstatus/common/models"
	"github.com/cjburchell/reefstatus/controller/service"
	"github.com/cjburchell/reefstatus/controller/profilux"
)

func currentPumps(controller *profilux.Controller, repo service.Controller) error {
	for i := 0; i < controller.GetCurrentPumpCount(); i++ {
		pump, err := repo.GetCurrentPump(models.GetID(models.CurrentPumpType, i))
		if err != nil && err != service.ErrNotFound {
			return err
		}

		found := err != service.ErrNotFound
		isAssigned, err := controller.GetIsCurrentPumpAssigned(i)

		if isAssigned {
			if !found {
				pump = models.NewCurrentPump(i)
			}

			pump.Value, err = controller.GetCurrentPumpValue(pump.Index)
			if err != nil {
				return err
			}

			err = repo.SetCurrentPump(pump, !found)
			if err != nil {
				return err
			}

		} else {
			if found {
				err = repo.DeleteCurrentPump(pump)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
