package history

import (
	"time"

	"github.com/cjburchell/reefstatus/server/data/repo"

	historyData "github.com/cjburchell/reefstatus/server/history/data"
	"github.com/cjburchell/reefstatus/server/history/save"

	"github.com/cjburchell/go-uatu"
)

const hourLogRate = time.Hour
const dayLogRate = time.Hour * 24

func Update(history historyData.HistoryData, controller repo.Controller, firstTime bool) {

	log.Debug("Saving Day Data")
	err := save.Day(history, controller)
	if err != nil {
		log.Error(err, "Error Saving Day data")
	}

	if firstTime {
		go updateWeekHistory(history, controller)
		go updateYearHistory(history, controller)
	}
}

func updateWeekHistory(historyData historyData.HistoryData, controller repo.Controller) {
	lastHourSavedTime, err := historyData.GetLastTimeWeekDataWasSaved()
	if err != nil {
		log.Error(err)
		return
	}

	timeSinceLastHourSaved := time.Duration(int64(time.Second) * (time.Now().Unix() - lastHourSavedTime.Unix()))
	duration := hourLogRate
	if timeSinceLastHourSaved > hourLogRate {
		err = save.Week(historyData, controller)
		if err != nil {
			log.Error(err)
			return
		}
	} else if !lastHourSavedTime.IsZero() {
		duration = hourLogRate - timeSinceLastHourSaved
	} else {
		err = save.Week(historyData, controller)
		if err != nil {
			log.Error(err)
			return
		}
	}

	for {
		log.Debugf("SaveWeekHistory Sleeping for %s", duration.String())
		<-time.After(duration)
		err := save.Week(historyData, controller)
		if err != nil {
			log.Error(err)
		}
		duration = hourLogRate
	}
}

func updateYearHistory(historyData historyData.HistoryData, controller repo.Controller) {
	lastHourSavedTime, err := historyData.GetLastTimeYearDataWasSaved()
	if err != nil {
		log.Error(err)
		return
	}

	timeSinceLastHourSaved := time.Duration(int64(time.Millisecond) * (time.Now().Unix() - lastHourSavedTime.Unix()))
	duration := dayLogRate
	if timeSinceLastHourSaved > dayLogRate {
		err = save.Year(historyData, controller)
		if err != nil {
			log.Error(err)
			return
		}
	} else if !lastHourSavedTime.IsZero() {
		duration = dayLogRate - timeSinceLastHourSaved
	} else {
		err = save.Year(historyData, controller)
		if err != nil {
			log.Error(err)
			return
		}
	}

	for {
		log.Debugf("SaveYearHistory Sleeping for %s", duration.String())
		<-time.After(duration)
		err := save.Year(historyData, controller)
		if err != nil {
			log.Error(err)
		}
		duration = dayLogRate
	}
}
