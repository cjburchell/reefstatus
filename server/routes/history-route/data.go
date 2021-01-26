package history_route

import (
	"encoding/json"
	"net/http"

	"github.com/cjburchell/reefstatus/server/history/data"

	logger "github.com/cjburchell/uatu-go"

	"github.com/gorilla/mux"
)

// SetupDataRoute setup the route
func SetupRoute(r *mux.Router, h data.HistoryData, log logger.ILog) {
	dataRoute := r.PathPrefix("/data").Subrouter()
	dataRoute.HandleFunc("/log/{ID}", func(writer http.ResponseWriter, request *http.Request) {
		handleDayData(writer, request, h, log)
	}).Methods("GET")
	dataRoute.HandleFunc("/logYear/{ID}", func(writer http.ResponseWriter, request *http.Request) {
		handleYearData(writer, request, h, log)
	}).Methods("GET")
	dataRoute.HandleFunc("/logWeek/{ID}", func(writer http.ResponseWriter, request *http.Request) {
		handleWeekData(writer, request, h, log)
	}).Methods("GET")
}

func handleDayData(w http.ResponseWriter, r *http.Request, historyData data.HistoryData, log logger.ILog) {
	vars := mux.Vars(r)
	id := vars["ID"]

	result, err := historyData.GetDayDataPoints(id)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("500 - Something bad happened! " + err.Error()))
		if err != nil{
			log.Error(err)
		}
		return
	}

	reply, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(reply)
	if err != nil{
		log.Error(err)
	}
}

func handleWeekData(w http.ResponseWriter, r *http.Request, historyData data.HistoryData, log logger.ILog) {
	vars := mux.Vars(r)
	id := vars["ID"]

	result, err := historyData.GetWeekDataPoints(id)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("500 - Something bad happened! " + err.Error()))
		if err != nil{
			log.Error(err)
		}
		return
	}

	reply, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(reply)
	if err != nil{
		log.Error(err)
	}
}

func handleYearData(w http.ResponseWriter, r *http.Request, historyData data.HistoryData, log logger.ILog) {
	vars := mux.Vars(r)
	id := vars["ID"]

	result, err := historyData.GetYearDataPoints(id)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("500 - Something bad happened! " + err.Error()))
		if err != nil{
			log.Error(err)
		}
		return
	}

	reply, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(reply)
	if err != nil{
		log.Error(err)
	}
}
