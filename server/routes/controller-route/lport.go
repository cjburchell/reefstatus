package controller_route

import (
	"fmt"

	"github.com/cjburchell/reefstatus/common/models"
	"github.com/cjburchell/reefstatus/server/data/repo"

	"github.com/pkg/errors"
)

type lport struct {
	controller repo.Controller
}

func (p lport) Set(item interface{}) error {
	resource, ok := item.(*models.LPort)
	if !ok {
		return errors.WithStack(fmt.Errorf("unable to cast to resource"))
	}

	return p.controller.SetLPort(*resource)
}

func (p lport) GetEmpty() interface{} {
	return &models.LPort{}
}

func (p lport) Get(id string) (interface{}, error) {
	return p.controller.GetLPort(id)
}

func (p lport) Delete(id string) error {
	return p.controller.DeleteLPort(id)
}

func (p lport) GetList() (interface{}, error) {
	return p.controller.GetLPorts()
}
