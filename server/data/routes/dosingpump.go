package routes

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/cjburchell/reefstatus-common/models"
	"github.com/cjburchell/reefstatus-data/repo"
)

type dosingpump struct {
	controller repo.Controller
}

func (p dosingpump) Set(item interface{}) error {
	resource, ok := item.(*models.DosingPump)
	if !ok {
		return errors.WithStack(fmt.Errorf("unable to cast to resource"))
	}

	return p.controller.SetDosingPump(*resource)
}

func (p dosingpump) GetEmpty() interface{} {
	return &models.DosingPump{}
}

func (p dosingpump) Get(id string) (interface{}, error) {
	return p.controller.GetDosingPump(id)
}

func (p dosingpump) Delete(id string) error {
	return p.controller.DeleteDosingPump(id)
}

func (p dosingpump) GetList() (interface{}, error) {
	return p.controller.GetDosingPumps()
}
