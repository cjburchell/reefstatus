package history_route

import (
	"encoding/json"
	"net/http"

	"github.com/cjburchell/reefstatus/server/history/data"

	"github.com/cjburchell/go-uatu"

	"github.com/gorilla/mux"
)

// SetupDataRoute setup the route
func SetupRoute(r *mux.Router, h data.HistoryData) {
	dataRoute := r.PathPrefix("/data").Subrouter()
	dataRoute.HandleFunc("/log/{ID}", func(writer http.ResponseWriter, request *http.Request) {
		handleDayData(writer, request, h)
	}).Methods("GET")
	dataRoute.HandleFunc("/logYear/{ID}", func(writer http.ResponseWriter, request *http.Request) {
		handleYearData(writer, request, h)
	}).Methods("GET")
	dataRoute.HandleFunc("/logWeek/{ID}", func(writer http.ResponseWriter, request *http.Request) {
		handleWeekData(writer, request, h)
	}).Methods("GET")
}

func handleDayData(w http.ResponseWriter, r *http.Request, historyData data.HistoryData) {
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

func handleWeekData(w http.ResponseWriter, r *http.Request, historyData data.HistoryData) {
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

func handleYearData(w http.ResponseWriter, r *http.Request, historyData data.HistoryData) {
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
