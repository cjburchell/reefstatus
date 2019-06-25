package controller_route

import (
	"fmt"

	"github.com/cjburchell/reefstatus/common/models"
	"github.com/cjburchell/reefstatus/server/data/repo"

	"github.com/pkg/errors"
)

type levelsensor struct {
	controller repo.Controller
}

func (p levelsensor) Set(item interface{}) error {
	resource, ok := item.(*models.LevelSensor)
	if !ok {
		return errors.WithStack(fmt.Errorf("unable to cast to resource"))
	}

	return p.controller.SetLevelSensor(*resource)
}

func (p levelsensor) GetEmpty() interface{} {
	return &models.LevelSensor{}
}

func (p levelsensor) Get(id string) (interface{}, error) {
	return p.controller.GetLevelSensor(id)
}

func (p levelsensor) Delete(id string) error {
	return p.controller.DeleteLevelSensor(id)
}

func (p levelsensor) GetList() (interface{}, error) {
	return p.controller.GetLevelSensors()
}
