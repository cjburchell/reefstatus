package routes

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/cjburchell/reefstatus-common/models"
	"github.com/cjburchell/reefstatus-data/repo"
)

type digitalinput struct {
	controller repo.Controller
}

func (p digitalinput) Set(item interface{}) error {
	resource, ok := item.(*models.DigitalInput)
	if !ok {
		return errors.WithStack(fmt.Errorf("unable to cast to resource"))
	}

	return p.controller.SetDigitalInput(*resource)
}

func (p digitalinput) GetEmpty() interface{} {
	return &models.DigitalInput{}
}

func (p digitalinput) Get(id string) (interface{}, error) {
	return p.controller.GetDigitalInput(id)
}

func (p digitalinput) Delete(id string) error {
	return p.controller.DeleteDigitalInput(id)
}

func (p digitalinput) GetList() (interface{}, error) {
	return p.controller.GetDigitalInputs()
}
