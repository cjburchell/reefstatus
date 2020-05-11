package update

import (
	"github.com/cjburchell/reefstatus/controller/service"
	"github.com/cjburchell/reefstatus/controller/profilux"
	"github.com/cjburchell/reefstatus/controller/settings"
	logger "github.com/cjburchell/uatu-go"
)

// All update all the data
func All(repo service.Controller, log logger.ILog) error {

	log.Debug("RefreshSettings - Start")
	profiluxController, err := profilux.NewController(settings.Connection)
	if err != nil {
		return err
	}

	defer profiluxController.Disconnect()

	profiluxController.ResetStats()
	err = info(profiluxController, repo)
	if err != nil {
		return err
	}

	err = probes(profiluxController, repo)
	if err != nil {
		return err
	}

	err = LevelSensors(profiluxController, repo)
	if err != nil {
		return err
	}

	err = digitalInputs(profiluxController, repo)
	if err != nil {
		return err
	}

	err = dosingPumps(profiluxController, repo)
	if err != nil {
		return err
	}

	err = lights(profiluxController, repo)
	if err != nil {
		return err
	}

	err = currentPumps(profiluxController, repo)
	if err != nil {
		return err
	}

	err = programmableLogic(profiluxController, repo)
	if err != nil {
		return err
	}

	err = sPorts(profiluxController, repo)
	if err != nil {
		return err
	}

	err = lPorts(profiluxController, repo)
	if err != nil {
		return err
	}

	err = repo.UpdateAssociations()
	log.Debugf("Call Count %d", profiluxController.GetCallCount())
	log.Debug("RefreshSettings - End")
	return err
}

// State update only the state data
func State(repo service.Controller, log logger.ILog) error {
	log.Debug("RefreshState - Start")
	profiluxController, err := profilux.NewController(settings.Connection)
	if err != nil {
		log.Errorf(err, "unable to connect")
	}

	defer profiluxController.Disconnect()

	profiluxController.ResetStats()
	err = InfoState(profiluxController, repo)
	if err != nil {
		return err
	}

	probes, err := repo.GetProbes()
	if err != nil {
		return err
	}

	for _, item := range probes {
		err = probeUpdateState(&item, profiluxController)
		if err != nil {
			return err
		}

		err = repo.SetProbe(item, false)
		if err != nil {
			return err
		}
	}

	levelSensors, err := repo.GetLevelSensors()
	if err != nil {
		return err
	}

	for _, item := range levelSensors {
		err = levelSensorUpdateState(&item, profiluxController)
		if err != nil {
			return err
		}

		err = repo.SetLevelSensor(item, false)
		if err != nil {
			return err
		}
	}

	digitalInputs, err := repo.GetDigitalInputs()
	if err != nil {
		return err
	}

	for _, item := range digitalInputs {
		item.Value, err = profiluxController.GetDigitalInputState(item.Index)
		if err != nil {
			return err
		}

		err = repo.SetDigitalInput(item, false)
		if err != nil {
			return err
		}
	}

	lights, err := repo.GetLights()
	if err != nil {
		return err
	}

	for _, item := range lights {
		err = lightUpdateState(&item, profiluxController)
		if err != nil {
			return err
		}

		err = repo.SetLight(item, false)
		if err != nil {
			return err
		}
	}

	pumps, err := repo.GetCurrentPumps()
	if err != nil {
		return err
	}

	for _, item := range pumps {
		item.Value, err = profiluxController.GetCurrentPumpValue(item.Index)
		if err != nil {
			return err
		}

		err = repo.SetCurrentPump(item, false)
		if err != nil {
			return err
		}
	}

	sPorts, err := repo.GetSPorts()
	if err != nil {
		return err
	}

	for _, item := range sPorts {
		err = sPortUpdateState(&item, profiluxController)
		if err != nil {
			return err
		}

		err = repo.SetSPort(item, false)
		if err != nil {
			return err
		}
	}

	lPorts, err := repo.GetLPorts()
	if err != nil {
		return err
	}

	for _, item := range lPorts {
		err = lPortUpdateState(&item, profiluxController)
		if err != nil {
			return err
		}

		err = repo.SetLPort(item, false)
		if err != nil {
			return err
		}
	}

	log.Debugf("Call Count %d", profiluxController.GetCallCount())
	log.Debug("RefreshState - End")

	return nil
}
