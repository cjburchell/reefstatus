package main

import (
	"net/http"

	"github.com/cjburchell/tools-go/env"

	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {

	clientLocation := env.Get("CLIENT_LOCATION", "client/dist/ReefStatus")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, clientLocation+"/index.html")
	})

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(clientLocation))))
}
