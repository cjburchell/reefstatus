package models

import (
	"fmt"

	"github.com/cjburchell/reefstatus/common/profilux/types"
)

type ProgrammableLogic struct {
	DisplayName string
	Index       int
	Function    types.LogicFunction
	Input1      types.PortMode
	Input2      types.PortMode
	Id          string
	Type        string
}

const ProgrammableLogicType = "ProgrammableLogic"

func NewProgrammableLogic(index int) ProgrammableLogic {
	var programmableLogic ProgrammableLogic
	programmableLogic.Type = ProgrammableLogicType
	programmableLogic.Id = GetID(ProgrammableLogicType, index)
	programmableLogic.Index = index
	programmableLogic.DisplayName = fmt.Sprintf("Programable Logic %d", index+1)
	return programmableLogic
}
