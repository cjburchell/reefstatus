package main

import (
	"net/http"
	"time"

	"github.com/cjburchell/reefstatus/server/history"

	"github.com/cjburchell/reefstatus/server/alert"

	"github.com/cjburchell/reefstatus/server/alert/state"

	"github.com/cjburchell/reefstatus/common/communication"

	"github.com/cjburchell/reefstatus/server/data/repo"
	"github.com/cjburchell/reefstatus/server/settings"

	historyData "github.com/cjburchell/reefstatus/server/history/data"

	"github.com/gorilla/handlers"

	"github.com/cjburchell/go-uatu"
	logSettings "github.com/cjburchell/go-uatu/settings"
	"github.com/gorilla/mux"
)

func main() {
	err := logSettings.SetupLogger()
	if err != nil {
		log.Warn("Unable to Connect to logger")
	}

	controllerRepo, err := repo.NewController(settings.RedisAddress, settings.RedisPassword)
	if err != nil {
		log.Fatal(err, "Unable to Connect to controller repo")
	}

	historyRepo, err := historyData.New(settings.MongoUrl)
	if err != nil {
		log.Fatal(err, "Unable to Connect to history database")
	}

	stateRepo, err := state.New(settings.MongoUrl)
	if err != nil {
		log.Fatal(err, "Unable to setup state repo")
	}

	session, err := communication.NewSession(settings.PubSubAddress, settings.PubSubToken)
	if err != nil {
		log.Fatal(err, "Unable to connect to Pub Sub")
	}

	defer session.Close()

	go RunUpdate(controllerRepo, session, stateRepo, historyRepo)

	r := mux.NewRouter()
	controller_route.SetupRoute(r, controllerRepo)
	command_route.SetupRoute(r, session)
	history_route.SetupRoute(r, historyRepo)
	client_route.SetupRoute(r)
	loggedRouter := handlers.LoggingHandler(log.Writer{Level: log.DEBUG}, r)

	log.Print("Starting Server at port 8090")
	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         ":8090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error(err)
	}
}

func RunUpdate(controllerRepo repo.Controller, session communication.SubscribeSession, stateRepo state.StateData, historyRepo historyData.HistoryData) {
	ch, err := session.QueueSubscribe(communication.UpdateMessage, "history")
	if err != nil {
		log.Fatal(err, "Unable to Subscribe to UpdateMessage")
	}

	firstTime := true
	for {
		<-ch
		go alert.Check(controllerRepo, stateRepo)
		go history.Update(historyRepo, controllerRepo, firstTime)
		firstTime = false
	}
}
