package main

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"

	"github.com/cjburchell/reefstatus-history/settings"

	"github.com/cjburchell/go-uatu"

	"github.com/cjburchell/reefstatus-history/save"

	"github.com/cjburchell/reefstatus-history/routes"
	"github.com/gorilla/mux"

	"github.com/cjburchell/reefstatus-common/data"

	"github.com/cjburchell/reefstatus-common/communication"
	history "github.com/cjburchell/reefstatus-history/data"
)

const hourLogRate = time.Hour
const dayLogRate = time.Hour * 24

func main() {
	err := log.Setup(settings.Log)
	if err != nil {
		log.Fatal(err, "Unable to Connect to logger")
	}

	historyData, err := history.New(settings.MongoUrl)
	if err != nil {
		log.Fatal(err, "Unable to Connect to history database")
	}

	controller, err := data.NewController(settings.DataServiceAddress, settings.DataServiceToken)
	if err != nil {
		log.Fatal(err, "Unable to Connect to data database")
	}

	r := mux.NewRouter()
	routes.SetupDataRoute(r, historyData)

	loggedRouter := handlers.LoggingHandler(log.Writer{Level: log.DEBUG}, r)

	log.Print("Starting Server at port 8092")
	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         ":8092",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	session, err := communication.NewSession(settings.PubSubAddress, settings.PubSubToken)
	if err != nil {
		log.Fatal(err, "Unable to Connect to pub sub")
	}

	defer session.Close()
	ch, err := session.QueueSubscribe(communication.UpdateHistoryMessage, "history")
	if err != nil {
		log.Fatal(err, "Unable to Subscribe to UpdateMessage")
	}

	firstTime := true
	for {
		<-ch
		log.Debug("Saving Day Data")
		err = save.Day(historyData, controller)
		if err != nil {
			log.Error(err, "Error Saving Day data")
		}

		if firstTime {
			go updateWeekHistory(historyData, controller)
			go updateYearHistory(historyData, controller)
			firstTime = false
		}
	}
}

func updateWeekHistory(historyData history.HistoryData, controller data.ControllerService) {
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

func updateYearHistory(historyData history.HistoryData, controller data.ControllerService) {
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
