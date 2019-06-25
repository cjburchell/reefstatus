package controller_route

import (
	"fmt"

	"github.com/cjburchell/reefstatus/common/models"
	"github.com/cjburchell/reefstatus/server/data/repo"

	"github.com/pkg/errors"
)

type sport struct {
	controller repo.Controller
}

func (p sport) Set(item interface{}) error {
	resource, ok := item.(*models.SPort)
	if !ok {
		return errors.WithStack(fmt.Errorf("unable to cast to resource"))
	}

	return p.controller.SetSPort(*resource)
}

func (p sport) GetEmpty() interface{} {
	return &models.SPort{}
}

func (p sport) Get(id string) (interface{}, error) {
	return p.controller.GetSPort(id)
}

func (p sport) Delete(id string) error {
	return p.controller.DeleteSPort(id)
}

func (p sport) GetList() (interface{}, error) {
	return p.controller.GetSPorts()
}
