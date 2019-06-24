package routes

import (
	"encoding/json"
	"net/http"

	"github.com/cjburchell/go-uatu"

	"github.com/cjburchell/reefstatus-history/data"

	"github.com/gorilla/mux"
)

var historyData data.HistoryData

// SetupDataRoute setup the route
func SetupDataRoute(r *mux.Router, h data.HistoryData) {
	historyData = h
	dataRoute := r.PathPrefix("/data").Subrouter()
	dataRoute.HandleFunc("/log/{ID}", handleDayData).Methods("GET")
	dataRoute.HandleFunc("/logYear/{ID}", handleYearData).Methods("GET")
	dataRoute.HandleFunc("/logWeek/{ID}", handleWeekData).Methods("GET")
}

func handleDayData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	result, err := historyData.GetDayDataPoints(id)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened! " + err.Error()))
		return
	}

	reply, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply)
}

func handleWeekData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	result, err := historyData.GetWeekDataPoints(id)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened! " + err.Error()))
		return
	}

	reply, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply)
}

func handleYearData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	result, err := historyData.GetYearDataPoints(id)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened! " + err.Error()))
		return
	}

	reply, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply)
}
