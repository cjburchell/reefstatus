package data

import (
	"time"

	"github.com/cjburchell/reefstatus-history/model"
)

// HistoryData interface
type HistoryData interface {
	SaveDayData(data model.Data) error
	SaveWeekData(data model.Data) error
	SaveYearData(data model.Data) error
	GetLastTimeYearDataWasSaved() (time.Time, error)
	GetLastTimeWeekDataWasSaved() (time.Time, error)
	GetDayDataPoints(dataType string) ([]model.Data, error)
	GetYearDataPoints(dataType string) ([]model.Data, error)
	GetWeekDataPoints(dataType string) ([]model.Data, error)
	GetDataPointsFromLastHour(dataType string) ([]model.Data, error)
}

// New History data
func New(address string) (HistoryData, error) {
	data := mongoData{}
	err := data.setup(address)
	return data, err
}
