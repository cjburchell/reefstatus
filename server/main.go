package main

import (
	"net/http"
	"time"

	"github.com/cjburchell/settings-go"

	clientRoute "github.com/cjburchell/reefstatus/server/routes/client-route"
	commandRoute "github.com/cjburchell/reefstatus/server/routes/command-route"
	controllerRoute "github.com/cjburchell/reefstatus/server/routes/controller-route"
	historyRoute "github.com/cjburchell/reefstatus/server/routes/history-route"
	appSettings "github.com/cjburchell/reefstatus/server/settings"
	"github.com/cjburchell/tools-go/env"

	"github.com/cjburchell/reefstatus/server/history"

	"github.com/cjburchell/reefstatus/server/alert"

	"github.com/cjburchell/reefstatus/server/alert/state"

	"github.com/cjburchell/reefstatus/common/communication"

	"github.com/cjburchell/reefstatus/server/data/repo"

	historyData "github.com/cjburchell/reefstatus/server/history/data"

	"github.com/gorilla/handlers"

	logger "github.com/cjburchell/uatu-go"
	"github.com/gorilla/mux"
)

func main() {
	set := settings.Get(env.Get("ConfigFile", ""))
	log := logger.Create(set)
	appConfig := appSettings.Get(set)

	controllerRepo, err := repo.NewController(appConfig.RedisAddress, appConfig.RedisPassword)
	if err != nil {
		log.Fatal(err, "Unable to Connect to controller repo")
	}

	historyRepo, err := historyData.New(appConfig.MongoUrl)
	if err != nil {
		log.Fatal(err, "Unable to Connect to history database")
	}

	stateRepo, err := state.New(appConfig.MongoUrl)
	if err != nil {
		log.Fatal(err, "Unable to setup state repo")
	}

	session, err := communication.NewSession(appConfig.PubSubAddress, appConfig.PubSubToken, log)
	if err != nil {
		log.Fatal(err, "Unable to connect to Pub Sub")
	}

	defer session.Close()

	go RunUpdate(controllerRepo, session, stateRepo, historyRepo, log, appConfig)

	r := mux.NewRouter()
	controllerRoute.SetupRoute(r, controllerRepo, log, appConfig.DataServiceToken)
	commandRoute.SetupRoute(r, session, log, appConfig.DataServiceToken)
	historyRoute.SetupRoute(r, historyRepo, log)
	clientRoute.SetupRoute(r)
	loggedRouter := handlers.LoggingHandler(log.GetWriter(logger.DEBUG), r)

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

func RunUpdate(controllerRepo repo.Controller, session communication.SubscribeSession, stateRepo state.StateData, historyRepo historyData.HistoryData, log logger.ILog, config appSettings.Config) {
	ch, err := session.QueueSubscribe(communication.UpdateMessage, "history")
	if err != nil {
		log.Fatal(err, "Unable to Subscribe to UpdateMessage")
	}

	firstTime := true
	for {
		<-ch
		go alert.Check(controllerRepo, stateRepo, log, config.SendOnReminder, config.SlackDestination)
		go history.Update(historyRepo, controllerRepo, firstTime, log)
		firstTime = false
	}
}
