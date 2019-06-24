package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cjburchell/go-uatu"

	"github.com/cjburchell/reefstatus-common/communication"

	"github.com/cjburchell/reefstatus-commands/settings"
	"github.com/gorilla/mux"
)

var session communication.Session

// SetupCommandRoute setup the route
func SetupCommandRoute(r *mux.Router, s communication.Session) {
	session = s
	commandRoute := r.PathPrefix("/command").Subrouter()
	commandRoute.Use(tokenMiddleware)
	commandRoute.HandleFunc("/feedpause", handleFeedPasue).Methods("POST")
	commandRoute.HandleFunc("/thunderstorm", handleThunderstorm).Methods("POST")
	commandRoute.HandleFunc("/resetReminder/{Index}", handleResetReminder).Methods("POST")
	commandRoute.HandleFunc("/maintenance/{Index}", handleMaintenance).Methods("POST")
	commandRoute.HandleFunc("/clearlevelalarm/{ID}", handleClearLevelAlarm).Methods("POST")
	commandRoute.HandleFunc("/startwaterchange/{ID}", handleStartWaterChange).Methods("POST")
}

func tokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		auth := request.Header.Get("Authorization")
		if auth != "APIKEY "+settings.DataServiceToken {
			response.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(response, request)
	})
}

func handleFeedPasue(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleFeedPasue %s", r.URL.String())
	var body []byte
	r.Body.Read(body)
	var enable bool
	json.Unmarshal(body, &enable)
	communication.FeedPause(session, enable)
	reply, _ := json.Marshal(true)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(reply)
}

func handleThunderstorm(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleThunderstorm %s", r.URL.String())
	var body []byte
	r.Body.Read(body)
	var duration int
	json.Unmarshal(body, &duration)
	communication.Thunderstorm(session, duration)
	reply, _ := json.Marshal(true)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(reply)
}

func handleResetReminder(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleResetReminder %s", r.URL.String())
	vars := mux.Vars(r)
	index, _ := strconv.Atoi(vars["Index"])
	communication.ResetReminder(session, index)
	reply, _ := json.Marshal(true)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(reply)
}
func handleMaintenance(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleMaintenance %s", r.URL.String())
	vars := mux.Vars(r)
	index, _ := strconv.Atoi(vars["Index"])
	var body []byte
	r.Body.Read(body)
	var enable bool
	json.Unmarshal(body, &enable)
	communication.Maintenance(session, index, enable)
	reply, _ := json.Marshal(true)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(reply)
}
func handleClearLevelAlarm(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleClearLevelAlarm %s", r.URL.String())
	vars := mux.Vars(r)
	id, _ := vars["ID"]
	communication.ClearLevelAlarm(session, id)
	reply, _ := json.Marshal(true)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(reply)
}
func handleStartWaterChange(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleThunderstorm %s", r.URL.String())
	vars := mux.Vars(r)
	id, _ := vars["ID"]
	communication.WaterChange(session, id)
	reply, _ := json.Marshal(true)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(reply)
}
