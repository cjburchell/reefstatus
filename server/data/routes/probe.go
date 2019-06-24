package routes

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/cjburchell/reefstatus-common/models"
	"github.com/cjburchell/reefstatus-data/repo"
)

type probe struct {
	controller repo.Controller
}

func (p probe) Set(item interface{}) error {
	probe, ok := item.(*models.Probe)
	if !ok {
		return errors.WithStack(fmt.Errorf("unable to cast to probe"))
	}

	return p.controller.SetProbe(*probe)
}

func (p probe) GetEmpty() interface{} {
	return &models.Probe{}
}

func (p probe) Get(id string) (interface{}, error) {
	return p.controller.GetProbe(id)
}

func (p probe) Delete(id string) error {
	return p.controller.DeleteProbe(id)
}

func (p probe) GetList() (interface{}, error) {
	return p.controller.GetProbes()
}
