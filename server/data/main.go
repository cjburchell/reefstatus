package main

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"

	"github.com/cjburchell/reefstatus-data/settings"

	"github.com/cjburchell/go-uatu"

	"github.com/cjburchell/reefstatus-data/repo"

	"github.com/cjburchell/reefstatus-data/routes"
	"github.com/gorilla/mux"
)

func main() {
	err := log.Setup(settings.Log)
	if err != nil {
		log.Warn("Unable to Connect to logger")
	}

	controllerRepo, err := repo.NewController(settings.RedisAddress, settings.RedisPassword)
	if err != nil {
		log.Fatalf(err, "Unable to Connect to controller repo")
	}

	r := mux.NewRouter()
	routes.SetupControllerRoute(r, controllerRepo)
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
