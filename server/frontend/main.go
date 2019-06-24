package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"

	"github.com/cjburchell/go-uatu"

	"github.com/cjburchell/reefstatus-frontend/settings"
	"github.com/gorilla/mux"
)

func main() {
	err := log.Setup(settings.Log)
	if err != nil {
		log.Warn(err, "Unable to Connect to logger")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ReefStatus/dist/ReefStatus/index.html")
	})

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("ReefStatus/dist/ReefStatus"))))

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
