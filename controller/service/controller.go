package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/cjburchell/reefstatus/common/models"
)

// ErrNotFound when resource is not found
var ErrNotFound = fmt.Errorf("not found")

// Controller interface
type Controller interface {
	GetInfo() (models.Info, error)
	GetDigitalInputs() ([]models.DigitalInput, error)
	GetDosingPumps() ([]models.DosingPump, error)
	GetLevelSensors() ([]models.LevelSensor, error)
	GetLights() ([]models.Light, error)
	GetLPorts() ([]models.LPort, error)
	GetProbes() ([]models.Probe, error)
	GetProgrammableLogics() ([]models.ProgrammableLogic, error)
	GetCurrentPumps() ([]models.CurrentPump, error)
	GetSPorts() ([]models.SPort, error)

	GetCurrentPump(id string) (models.CurrentPump, error)
	GetDigitalInput(id string) (models.DigitalInput, error)
	GetDosingPump(id string) (models.DosingPump, error)
	GetLevelSensor(id string) (models.LevelSensor, error)
	GetLight(id string) (models.Light, error)
	GetLPort(id string) (models.LPort, error)
	GetProbe(id string) (models.Probe, error)
	GetProgrammableLogic(id string) (models.ProgrammableLogic, error)
	GetSPort(id string) (models.SPort, error)

	SetInfo(info models.Info, create bool) error
	SetDigitalInput(input models.DigitalInput, create bool) error
	SetDosingPump(pump models.DosingPump, create bool) error
	SetLevelSensor(sensor models.LevelSensor, create bool) error
	SetLight(light models.Light, create bool) error
	SetLPort(port models.LPort, create bool) error
	SetProbe(probe models.Probe, create bool) error
	SetProgrammableLogic(item models.ProgrammableLogic, create bool) error
	SetCurrentPump(pump models.CurrentPump, create bool) error
	SetSPort(port models.SPort, create bool) error

	DeleteDigitalInput(input models.DigitalInput) error
	DeleteDosingPump(pump models.DosingPump) error
	DeleteLevelSensor(models.LevelSensor) error
	DeleteLight(models.Light) error
	DeleteLPort(models.LPort) error
	DeleteProbe(models.Probe) error
	DeleteProgrammableLogic(models.ProgrammableLogic) error
	DeleteCurrentPump(models.CurrentPump) error
	DeleteSPort(models.SPort) error

	UpdateAssociations() error
}

type controller struct {
	client  *http.Client
	address string
	token   string
}

func (c controller) GetInfo() (models.Info, error) {
	var result models.Info
	err := c.get("info", &result)
	return result, err
}

const controllerPath = "/api/v1/controller"

func (c controller) get(path string, target interface{}) error {
	fullPath := c.address + controllerPath + path
	//log.Debugf("GET %s", fullPath)
	req, err := http.NewRequest("GET", fullPath, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("APIKEY %s", c.token))

	resp, err := c.client.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return ErrNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("controler GET %s error: %d", fullPath, resp.StatusCode)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c controller) getID(path, id string, target interface{}) error {
	fullPath := c.address + controllerPath + path + "/" + id
	//log.Debugf("GET %s", fullPath)
	req, err := http.NewRequest("GET", fullPath, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("APIKEY %s", c.token))

	resp, err := c.client.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return ErrNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return errors.WithStack(fmt.Errorf("controler GET %s error: %d", fullPath, resp.StatusCode))
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c controller) send(method string, path string, target interface{}) error {
	//log.Debugf("%s %s", method, path)

	jsonValue, err := json.Marshal(target)
	if err != nil {
		return errors.WithStack(err)
	}

	req, err := http.NewRequest(method, path, bytes.NewBuffer(jsonValue))
	if err != nil {
		return errors.WithStack(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("APIKEY %s", c.token))
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}

	if resp.StatusCode != http.StatusNoContent {
		return errors.WithStack(fmt.Errorf("controler %s %s error: %d", method, path, resp.StatusCode))
	}

	return nil
}

func (c controller) create(path string, target interface{}) error {
	return c.send("POST", c.address+controllerPath+path, target)
}

func (c controller) update(path string, target interface{}) error {
	return c.send("PUT", c.address+controllerPath+path, target)
}

func (c controller) updateID(path string, id string, target interface{}) error {
	return c.send("PUT", c.address+controllerPath+path+"/"+id, target)
}

func (c controller) delete(path, id string) error {
	fullPath := c.address + controllerPath + path + "/" + id
	//log.Debugf("DELETE %s", fullPath)
	req, err := http.NewRequest("DELETE", fullPath, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("APIKEY %s", c.token))

	resp, err := c.client.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}

	if resp.StatusCode != http.StatusNoContent {
		return errors.WithStack(fmt.Errorf("controler DELETE %s error: %d", fullPath, resp.StatusCode))
	}

	return nil
}

func (c controller) GetDigitalInputs() ([]models.DigitalInput, error) {
	var result []models.DigitalInput
	err := c.get("digitalinput", &result)
	return result, err
}

func (c controller) GetDosingPumps() ([]models.DosingPump, error) {
	var result []models.DosingPump
	err := c.get("dosingpump", &result)
	return result, err
}

func (c controller) GetLevelSensors() ([]models.LevelSensor, error) {
	var result []models.LevelSensor
	err := c.get("levelsensor", &result)
	return result, err
}

func (c controller) GetLights() ([]models.Light, error) {
	var result []models.Light
	err := c.get("light", &result)
	return result, err
}

func (c controller) GetLPorts() ([]models.LPort, error) {
	var result []models.LPort
	err := c.get("lport", &result)
	return result, err
}

func (c controller) GetProbes() ([]models.Probe, error) {
	var result []models.Probe
	err := c.get("probe", &result)
	return result, err
}

func (c controller) GetProgrammableLogics() ([]models.ProgrammableLogic, error) {
	var result []models.ProgrammableLogic
	err := c.get("programmablelogic", &result)
	return result, err
}

func (c controller) GetCurrentPumps() ([]models.CurrentPump, error) {
	var result []models.CurrentPump
	err := c.get("pump", &result)
	return result, err
}

func (c controller) GetSPorts() ([]models.SPort, error) {
	var result []models.SPort
	err := c.get("sport", &result)
	return result, err
}

func (c controller) GetCurrentPump(id string) (models.CurrentPump, error) {
	var result models.CurrentPump
	err := c.getID("pump", id, &result)
	return result, err
}

func (c controller) GetDigitalInput(id string) (models.DigitalInput, error) {
	var result models.DigitalInput
	err := c.getID("digitalinput", id, &result)
	return result, err
}

func (c controller) GetDosingPump(id string) (models.DosingPump, error) {
	var result models.DosingPump
	err := c.getID("dosingpump", id, &result)
	return result, err
}

func (c controller) GetLevelSensor(id string) (models.LevelSensor, error) {
	var result models.LevelSensor
	err := c.getID("levelsensor", id, &result)
	return result, err
}

func (c controller) GetLight(id string) (models.Light, error) {
	var result models.Light
	err := c.getID("light", id, &result)
	return result, err
}

func (c controller) GetLPort(id string) (models.LPort, error) {
	var result models.LPort
	err := c.getID("lport", id, &result)
	return result, err
}

func (c controller) GetProbe(id string) (models.Probe, error) {
	var result models.Probe
	err := c.getID("probe", id, &result)
	return result, err
}

func (c controller) GetProgrammableLogic(id string) (models.ProgrammableLogic, error) {
	var result models.ProgrammableLogic
	err := c.getID("programmablelogic", id, &result)
	return result, err
}

func (c controller) GetSPort(id string) (models.SPort, error) {
	var result models.SPort
	err := c.getID("sport", id, &result)
	return result, err
}

func (c controller) SetInfo(info models.Info, create bool) error {
	if create {
		return c.create("info", info)
	}

	return c.update("info", info)
}

func (c controller) SetDigitalInput(item models.DigitalInput, create bool) error {
	if create {
		return c.create("digitalinput", item)
	}
	return c.updateID("digitalinput", item.ID, item)

}

func (c controller) SetDosingPump(item models.DosingPump, create bool) error {
	if create {
		return c.create("dosingpump", item)
	}
	return c.updateID("dosingpump", item.ID, item)

}

func (c controller) SetLevelSensor(item models.LevelSensor, create bool) error {
	if create {
		return c.create("levelsensor", item)
	}
	return c.updateID("levelsensor", item.ID, item)

}

func (c controller) SetLight(item models.Light, create bool) error {
	if create {
		return c.create("light", item)
	}
	return c.updateID("light", item.ID, item)

}

func (c controller) SetLPort(item models.LPort, create bool) error {
	if create {
		return c.create("lport", item)
	}
	return c.updateID("lport", item.ID, item)

}

func (c controller) SetProbe(item models.Probe, create bool) error {
	if create {
		return c.create("probe", item)
	}
	return c.updateID("probe", item.ID, item)

}

func (c controller) SetProgrammableLogic(item models.ProgrammableLogic, create bool) error {
	if create {
		return c.create("programmablelogic", item)
	}
	return c.updateID("programmablelogic", item.Id, item)

}

func (c controller) SetCurrentPump(item models.CurrentPump, create bool) error {
	if create {
		return c.create("pump", item)
	}
	return c.updateID("pump", item.ID, item)

}

func (c controller) SetSPort(item models.SPort, create bool) error {
	if create {
		return c.create("sport", item)
	}
	return c.updateID("sport", item.ID, item)

}

func (c controller) DeleteDigitalInput(input models.DigitalInput) error {
	return c.delete("digitalinput", input.ID)
}

func (c controller) DeleteDosingPump(pump models.DosingPump) error {
	return c.delete("dosingpump", pump.ID)
}

func (c controller) DeleteLevelSensor(item models.LevelSensor) error {
	return c.delete("levelsensor", item.ID)
}

func (c controller) DeleteLight(item models.Light) error {
	return c.delete("light", item.ID)
}

func (c controller) DeleteLPort(item models.LPort) error {
	return c.delete("lport", item.ID)
}

func (c controller) DeleteProbe(item models.Probe) error {
	return c.delete("probe", item.ID)
}

func (c controller) DeleteProgrammableLogic(item models.ProgrammableLogic) error {
	return c.delete("programmablelogic", item.Id)
}

func (c controller) DeleteCurrentPump(item models.CurrentPump) error {
	return c.delete("pump", item.ID)
}

func (c controller) DeleteSPort(item models.SPort) error {
	return c.delete("sport", item.ID)
}

func (c controller) UpdateAssociations() error {
	req, err := http.NewRequest("POST", c.address+controllerPath+"/updateAssociations", nil)
	req.Header.Add("Authorization", fmt.Sprintf("APIKEY %s", c.token))
	resp, err := c.client.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.WithStack(fmt.Errorf("controler updateAssociations error getting: %d", resp.StatusCode))
	}

	return nil
}

// NewController creates a new controller
func NewController(address, token string) (Controller, error) {
	client := &http.Client{}
	return &controller{client, address, token}, nil
}
