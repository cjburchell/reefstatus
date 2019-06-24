package routes

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/cjburchell/reefstatus-common/models"
	"github.com/cjburchell/reefstatus-data/repo"
)

type pump struct {
	controller repo.Controller
}

func (p pump) Set(item interface{}) error {
	resource, ok := item.(*models.CurrentPump)
	if !ok {
		return errors.WithStack(fmt.Errorf("unable to cast to resource"))
	}

	return p.controller.SetCurrentPump(*resource)
}

func (p pump) GetEmpty() interface{} {
	return &models.CurrentPump{}
}

func (p pump) Get(id string) (interface{}, error) {
	return p.controller.GetCurrentPump(id)
}

func (p pump) Delete(id string) error {
	return p.controller.DeleteCurrentPump(id)
}

func (p pump) GetList() (interface{}, error) {
	return p.controller.GetCurrentPumps()
}
