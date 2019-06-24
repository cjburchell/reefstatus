package repo

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/cjburchell/reefstatus-common/models"

	"github.com/go-redis/redis"
)

// ErrNotFound when a element is not found
var ErrNotFound = errors.New("not found")

// Controller repo interface
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

	SetInfo(info models.Info) error
	SetDigitalInput(input models.DigitalInput) error
	SetDosingPump(pump models.DosingPump) error
	SetLevelSensor(models.LevelSensor) error
	SetLight(models.Light) error
	SetLPort(models.LPort) error
	SetProbe(models.Probe) error
	SetProgrammableLogic(models.ProgrammableLogic) error
	SetCurrentPump(models.CurrentPump) error
	SetSPort(models.SPort) error

	DeleteDigitalInput(id string) error
	DeleteDosingPump(id string) error
	DeleteLevelSensor(id string) error
	DeleteLight(id string) error
	DeleteLPort(id string) error
	DeleteProbe(id string) error
	DeleteProgrammableLogic(id string) error
	DeleteCurrentPump(id string) error
	DeleteSPort(id string) error
}

type controller struct {
	client *redis.Client
}

// NewController create a new controller
func NewController(address string, password string) (Controller, error) {
	c := controller{}
	err := c.setup(address, password)
	return c, err
}

func (controller *controller) setup(address string, password string) error {
	controller.client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0, // use default DB
	})

	_, err := controller.client.Ping().Result()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func buildKey(key string, id string) string {
	return fmt.Sprintf("%s:%s", key, id)
}

func (controller controller) getItem(key string, item interface{}) error {
	result, err := controller.client.Get(key).Result()
	if err == redis.Nil {
		return ErrNotFound
	}

	if err != nil {
		return errors.WithStack(err)
	}

	err = json.Unmarshal([]byte(result), item)
	if err != nil {
		return errors.WithStack(err)
	}

	return err
}

func (controller controller) setItem(key string, item interface{}) error {
	data, err := json.Marshal(item)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = controller.client.Set(key, string(data), 0).Result()
	if err != nil {
		return errors.WithStack(err)
	}

	return err
}

func (controller controller) deleteItem(key string) error {
	_, err := controller.client.Del(key).Result()
	if err == redis.Nil {
		return ErrNotFound
	}

	if err != nil {
		return errors.WithStack(err)
	}

	return err
}

func (controller controller) getKeys(key string) ([]string, error) {
	keys, _, err := controller.client.Scan(0, key+":*", 100).Result()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return keys, err
}

const (
	infoKey              = "reefstatus:info"
	digitalInputKey      = "reefstatus:digitalinput"
	dosingPumpKey        = "reefstatus:dosingpump"
	levelSensorKey       = "reefstatus:levelSensor"
	lightsKey            = "reefstatus:light"
	lPortKey             = "reefstatus:lPort"
	probeKey             = "reefstatus:probe"
	programmableLogicKey = "reefstatus:programmableLogic"
	pumpKey              = "reefstatus:pump"
	sPortKey             = "reefstatus:sPort"
)

func (controller controller) GetInfo() (models.Info, error) {
	info := models.NewInfo()
	err := controller.getItem(infoKey, &info)
	return info, err
}
func (controller controller) GetDigitalInputs() ([]models.DigitalInput, error) {
	keys, err := controller.getKeys(digitalInputKey)
	if err != nil {
		return nil, err
	}

	items := make([]models.DigitalInput, len(keys))
	for idx, key := range keys {
		err = controller.getItem(key, &items[idx])
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}
func (controller controller) GetDosingPumps() ([]models.DosingPump, error) {
	keys, err := controller.getKeys(dosingPumpKey)
	if err != nil {
		return nil, err
	}

	items := make([]models.DosingPump, len(keys))
	for idx, key := range keys {
		err = controller.getItem(key, &items[idx])
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}
func (controller controller) GetLevelSensors() ([]models.LevelSensor, error) {
	keys, err := controller.getKeys(levelSensorKey)
	if err != nil {
		return nil, err
	}

	items := make([]models.LevelSensor, len(keys))
	for idx, key := range keys {
		err = controller.getItem(key, &items[idx])
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}
func (controller controller) GetLights() ([]models.Light, error) {
	keys, err := controller.getKeys(lightsKey)
	if err != nil {
		return nil, err
	}

	items := make([]models.Light, len(keys))
	for idx, key := range keys {
		err = controller.getItem(key, &items[idx])
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}
func (controller controller) GetLPorts() ([]models.LPort, error) {
	keys, err := controller.getKeys(lPortKey)
	if err != nil {
		return nil, err
	}

	items := make([]models.LPort, len(keys))
	for idx, key := range keys {
		err = controller.getItem(key, &items[idx])
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}
func (controller controller) GetProbes() ([]models.Probe, error) {
	keys, err := controller.getKeys(probeKey)
	if err != nil {
		return nil, err
	}

	items := make([]models.Probe, len(keys))
	for idx, key := range keys {
		err = controller.getItem(key, &items[idx])
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}
func (controller controller) GetProgrammableLogics() ([]models.ProgrammableLogic, error) {
	keys, err := controller.getKeys(programmableLogicKey)
	if err != nil {
		return nil, err
	}

	items := make([]models.ProgrammableLogic, len(keys))
	for idx, key := range keys {
		err = controller.getItem(key, &items[idx])
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}
func (controller controller) GetCurrentPumps() ([]models.CurrentPump, error) {
	keys, err := controller.getKeys(pumpKey)
	if err != nil {
		return nil, err
	}

	items := make([]models.CurrentPump, len(keys))
	for idx, key := range keys {
		err = controller.getItem(key, &items[idx])
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}
func (controller controller) GetSPorts() ([]models.SPort, error) {
	keys, err := controller.getKeys(sPortKey)
	if err != nil {
		return nil, err
	}

	items := make([]models.SPort, len(keys))
	for idx, key := range keys {
		err = controller.getItem(key, &items[idx])
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (controller controller) GetDigitalInput(id string) (models.DigitalInput, error) {
	var result models.DigitalInput
	err := controller.getItem(buildKey(digitalInputKey, id), &result)
	return result, err
}
func (controller controller) GetCurrentPump(id string) (models.CurrentPump, error) {
	var result models.CurrentPump
	err := controller.getItem(buildKey(pumpKey, id), &result)
	return result, err
}
func (controller controller) GetDosingPump(id string) (models.DosingPump, error) {
	var result models.DosingPump
	err := controller.getItem(buildKey(dosingPumpKey, id), &result)
	return result, err
}
func (controller controller) GetLevelSensor(id string) (models.LevelSensor, error) {
	var result models.LevelSensor
	err := controller.getItem(buildKey(levelSensorKey, id), &result)
	return result, err
}
func (controller controller) GetLight(id string) (models.Light, error) {
	var result models.Light
	err := controller.getItem(buildKey(lightsKey, id), &result)
	return result, err
}
func (controller controller) GetLPort(id string) (models.LPort, error) {
	var result models.LPort
	err := controller.getItem(buildKey(lPortKey, id), &result)
	return result, err
}
func (controller controller) GetProbe(id string) (models.Probe, error) {
	var result models.Probe
	err := controller.getItem(buildKey(probeKey, id), &result)
	return result, err
}
func (controller controller) GetProgrammableLogic(id string) (models.ProgrammableLogic, error) {
	var result models.ProgrammableLogic
	err := controller.getItem(buildKey(programmableLogicKey, id), &result)
	return result, err
}
func (controller controller) GetSPort(id string) (models.SPort, error) {
	var result models.SPort
	err := controller.getItem(buildKey(sPortKey, id), &result)
	return result, err
}

func (controller controller) SetInfo(info models.Info) error {
	return controller.setItem(infoKey, info)
}
func (controller controller) SetDigitalInput(input models.DigitalInput) error {
	return controller.setItem(buildKey(digitalInputKey, input.ID), input)
}
func (controller controller) SetDosingPump(pump models.DosingPump) error {
	return controller.setItem(buildKey(digitalInputKey, pump.ID), pump)
}
func (controller controller) SetLevelSensor(sensor models.LevelSensor) error {
	return controller.setItem(buildKey(levelSensorKey, sensor.ID), sensor)
}
func (controller controller) SetLight(item models.Light) error {
	return controller.setItem(buildKey(lightsKey, item.ID), item)
}
func (controller controller) SetLPort(item models.LPort) error {
	return controller.setItem(buildKey(lPortKey, item.ID), item)
}
func (controller controller) SetProbe(item models.Probe) error {
	return controller.setItem(buildKey(probeKey, item.ID), item)
}
func (controller controller) SetProgrammableLogic(item models.ProgrammableLogic) error {
	return controller.setItem(buildKey(programmableLogicKey, item.Id), item)
}
func (controller controller) SetCurrentPump(item models.CurrentPump) error {
	return controller.setItem(buildKey(pumpKey, item.ID), item)
}
func (controller controller) SetSPort(item models.SPort) error {
	return controller.setItem(buildKey(sPortKey, item.ID), item)
}

func (controller controller) DeleteDigitalInput(id string) error {
	return controller.deleteItem(buildKey(digitalInputKey, id))
}
func (controller controller) DeleteDosingPump(id string) error {
	return controller.deleteItem(buildKey(dosingPumpKey, id))
}
func (controller controller) DeleteLevelSensor(id string) error {
	return controller.deleteItem(buildKey(levelSensorKey, id))
}
func (controller controller) DeleteLight(id string) error {
	return controller.deleteItem(buildKey(lightsKey, id))
}
func (controller controller) DeleteLPort(id string) error {
	return controller.deleteItem(buildKey(lPortKey, id))
}
func (controller controller) DeleteProbe(id string) error {
	return controller.deleteItem(buildKey(probeKey, id))
}
func (controller controller) DeleteProgrammableLogic(id string) error {
	return controller.deleteItem(buildKey(programmableLogicKey, id))
}
func (controller controller) DeleteCurrentPump(id string) error {
	return controller.deleteItem(buildKey(pumpKey, id))
}
func (controller controller) DeleteSPort(id string) error {
	return controller.deleteItem(buildKey(sPortKey, id))
}
