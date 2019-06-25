package controller_route

import (
	"fmt"

	"github.com/cjburchell/reefstatus/common/models"
	"github.com/cjburchell/reefstatus/server/data/repo"

	"github.com/pkg/errors"
)

type programmablelogic struct {
	controller repo.Controller
}

func (p programmablelogic) Set(item interface{}) error {
	resource, ok := item.(*models.ProgrammableLogic)
	if !ok {
		return errors.WithStack(fmt.Errorf("unable to cast to resource"))
	}

	return p.controller.SetProgrammableLogic(*resource)
}

func (p programmablelogic) GetEmpty() interface{} {
	return &models.ProgrammableLogic{}
}

func (p programmablelogic) Get(id string) (interface{}, error) {
	return p.controller.GetProgrammableLogic(id)
}

func (p programmablelogic) Delete(id string) error {
	return p.controller.DeleteProgrammableLogic(id)
}

func (p programmablelogic) GetList() (interface{}, error) {
	return p.controller.GetProgrammableLogics()
}
