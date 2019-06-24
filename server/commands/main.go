package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"

	"github.com/cjburchell/reefstatus-commands/routes"
	"github.com/cjburchell/reefstatus-commands/settings"

	"github.com/cjburchell/go-uatu"

	"github.com/cjburchell/reefstatus-common/communication"

	"github.com/gorilla/mux"
)

func main() {
	err := log.Setup(settings.Log)
	if err != nil {
		log.Warn(err, "Unable to Connect to logger")
	}

	r := mux.NewRouter()
	session, err := communication.NewSession(settings.PubSubAddress, settings.PubSubToken)
	if err != nil {
		log.Fatal(err, "Unable to connect to pub sub")
	}

	defer session.Close()
	routes.SetupCommandRoute(r, session)

	loggedRouter := handlers.LoggingHandler(log.Writer{Level: log.DEBUG}, r)

	log.Print("Starting Server at port 8091")
	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         ":8091",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	srv.Shutdown(ctx)

	log.Print("shutting down")
	os.Exit(0)
}
