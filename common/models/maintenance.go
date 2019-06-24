package models

import (
	"fmt"
)

// Maintenance model
type Maintenance struct {
	DisplayName string
	Index       int
	IsActive    bool
	Duration    int
	TimeLeft    int
}

// NewMaintenance creates object
func NewMaintenance(index int) *Maintenance {
	var maintenance Maintenance
	maintenance.Index = index
	maintenance.DisplayName = fmt.Sprintf("Maintenance%d", index+1)
	return &maintenance
}
