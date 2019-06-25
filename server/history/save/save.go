package save

import (
	"time"

	"github.com/cjburchell/reefstatus/server/data/repo"

	"github.com/cjburchell/reefstatus/server/history/model"

	historyData "github.com/cjburchell/reefstatus/server/history/data"
)

// Day save data
func Day(historyData historyData.HistoryData, controller repo.Controller) error {
	now := time.Now()
	probes, err := controller.GetProbes()
	if err != nil {
		return err
	}

	for _, probe := range probes {
		err := historyData.SaveDayData(model.Data{Type: probe.ID, Value: probe.Value, Time: now})
		if err != nil {
			return err
		}
	}

	return nil
}

// Week save data
func Week(historyData historyData.HistoryData, controller repo.Controller) error {

	now := time.Now()
	probes, err := controller.GetProbes()
	if err != nil {
		return err
	}

	for _, probe := range probes {
		average, err := calculateLastHourAverage(probe.ID, historyData)
		if err != nil {
			return err
		}

		err = historyData.SaveWeekData(model.Data{Type: probe.ID, Value: average, Time: now})
		if err != nil {
			return err
		}
	}
	return nil
}

// Year save data
func Year(historyData historyData.HistoryData, controller repo.Controller) error {
	now := time.Now()
	probes, err := controller.GetProbes()
	if err != nil {
		return err
	}

	for _, probe := range probes {
		average, err := calculateLastDayAverage(probe.ID, historyData)
		if err != nil {
			return err
		}

		err = historyData.SaveYearData(model.Data{Type: probe.ID, Value: average, Time: now})
		if err != nil {
			return err
		}
	}
	return nil
}

func average(data []model.Data) float64 {
	if len(data) == 0 {
		return 0
	}

	sum := float64(0)
	for _, item := range data {
		sum += item.Value
	}

	return sum / float64(len(data))
}

func calculateLastHourAverage(dataType string, historyData historyData.HistoryData) (float64, error) {
	dataPoints, err := historyData.GetDataPointsFromLastHour(dataType)
	if err != nil {
		return 0, err
	}
	return average(dataPoints), nil
}

func calculateLastDayAverage(dataType string, historyData historyData.HistoryData) (float64, error) {
	dataPoints, err := historyData.GetDayDataPoints(dataType)
	if err != nil {
		return 0, err
	}
	return average(dataPoints), nil
}
