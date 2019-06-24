package models

import "fmt"

// BaseInfo base information for most models
type BaseInfo struct {
	ID          string
	Type        string
	DisplayName string
	Units       string
}

// GetID Generates an id
func GetID(typ string, index int) string {
	return fmt.Sprintf("%s%d", typ, 1+index)
}
