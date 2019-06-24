package update

import (
	"github.com/cjburchell/profilux-go"
	"github.com/cjburchell/profilux-go/types"
	"github.com/cjburchell/reefstatus-common/data"
	"github.com/cjburchell/reefstatus-common/models"
)

func lPortUpdate(port *models.LPort, controller *profilux.Controller) error {
	var err error
	port.Mode, err = controller.GetLPortFunction(port.PortNumber)
	if err != nil {
		return err
	}

	port.Value, err = controller.GetLPortValue(port.PortNumber)
	return err
}

func lPortUpdateState(port *models.LPort, controller *profilux.Controller) error {
	var err error
	port.Value, err = controller.GetLPortValue(port.PortNumber)
	return err
}

func lPorts(profiluxController *profilux.Controller, repo data.ControllerService) error {
	count, err := profiluxController.GetLPortCount()
	if err != nil {
		return err
	}

	for portNumber := 0; portNumber < count; portNumber++ {
		port, err := repo.GetLPort(models.GetID(models.LPortType, portNumber))
		if err != nil && err != data.ErrNotFound {
			return err
		}

		found := err != data.ErrNotFound

		mode, err := profiluxController.GetLPortFunction(portNumber)
		if err != nil {
			return err
		}

		if mode.DeviceMode != types.DeviceModeAlwaysOff {
			if !found {
				port = models.NewLPort(portNumber)
			}

			err = lPortUpdate(&port, profiluxController)
			if err != nil {
				return err
			}

			err = repo.SetLPort(port, !found)
			if err != nil {
				return err
			}

		} else {
			if found {
				err = repo.DeleteLPort(port)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
