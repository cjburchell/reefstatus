package update

import (
	"github.com/cjburchell/reefstatus/common/models"
	"github.com/cjburchell/reefstatus/common/profilux/types"
	"github.com/cjburchell/reefstatus/controller/data"
	"github.com/cjburchell/reefstatus/controller/profilux"
)

func sPortUpdate(port *models.SPort, controller *profilux.Controller) error {
	var err error
	port.Mode, err = controller.GetSPortFunction(port.PortNumber)
	if err != nil {
		return err
	}

	port.Value, err = controller.GetSPortValue(port.PortNumber)
	if err != nil {
		return err
	}

	port.IsActive = port.Value == types.CurrentStateOn
	port.DisplayName, err = controller.GetSPortName(port.PortNumber)
	return err
}

func sPortUpdateState(port *models.SPort, controller *profilux.Controller) error {
	var err error
	port.Value, err = controller.GetSPortValue(port.PortNumber)
	if err != nil {
		return err
	}

	port.IsActive = port.Value == types.CurrentStateOn
	return nil
}

func sPorts(profiluxController *profilux.Controller, repo data.ControllerService) error {
	count, err := profiluxController.GetSPortCount()
	if err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		port, err := repo.GetSPort(models.GetID(models.SPortType, i))
		if err != nil && err != data.ErrNotFound {
			return err
		}

		found := err != data.ErrNotFound
		mode, err := profiluxController.GetSPortFunction(i)
		if err != nil {
			return err
		}

		if mode.DeviceMode != types.DeviceModeAlwaysOff {
			if !found {
				port = models.NewSPort(i)
			}

			err = sPortUpdate(&port, profiluxController)
			if err != nil {
				return err
			}

			err = repo.SetSPort(port, !found)
			if err != nil {
				return err
			}

		} else {
			if found {
				err = repo.DeleteSPort(port)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
