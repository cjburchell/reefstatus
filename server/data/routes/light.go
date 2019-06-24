package routes

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/cjburchell/reefstatus-common/models"
	"github.com/cjburchell/reefstatus-data/repo"
)

type light struct {
	controller repo.Controller
}

func (p light) Set(item interface{}) error {
	resource, ok := item.(*models.Light)
	if !ok {
		return errors.WithStack(fmt.Errorf("unable to cast to resource"))
	}

	return p.controller.SetLight(*resource)
}

func (p light) GetEmpty() interface{} {
	return &models.Light{}
}

func (p light) Get(id string) (interface{}, error) {
	return p.controller.GetLight(id)
}

func (p light) Delete(id string) error {
	return p.controller.DeleteLight(id)
}

func (p light) GetList() (interface{}, error) {
	return p.controller.GetLights()
}
