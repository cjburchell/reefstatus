package command_route

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cjburchell/reefstatus/server/routes/token"

	"github.com/cjburchell/reefstatus/server/commands"

	"github.com/cjburchell/reefstatus/common/communication"

	"github.com/cjburchell/go-uatu"

	"github.com/gorilla/mux"
)

var session communication.Session

// SetupCommandRoute setup the route
func SetupRoute(r *mux.Router, s communication.Session) {
	session = s
	commandRoute := r.PathPrefix("api/v1/command").Subrouter()
	commandRoute.Use(token.Middleware)
	commandRoute.HandleFunc("/feedpause", func(writer http.ResponseWriter, request *http.Request) {
		handleFeedPasue(writer, request, s)
	}).Methods("POST")
	commandRoute.HandleFunc("/thunderstorm", func(writer http.ResponseWriter, request *http.Request) {
		handleThunderstorm(writer, request, s)
	}).Methods("POST")
	commandRoute.HandleFunc("/resetReminder/{Index}", func(writer http.ResponseWriter, request *http.Request) {
		handleResetReminder(writer, request, s)
	}).Methods("POST")
	commandRoute.HandleFunc("/maintenance/{Index}", func(writer http.ResponseWriter, request *http.Request) {
		handleMaintenance(writer, request, s)
	}).Methods("POST")
	commandRoute.HandleFunc("/clearlevelalarm/{ID}", func(writer http.ResponseWriter, request *http.Request) {
		handleClearLevelAlarm(writer, request, s)
	}).Methods("POST")
	commandRoute.HandleFunc("/startwaterchange/{ID}", func(writer http.ResponseWriter, request *http.Request) {
		handleStartWaterChange(writer, request, s)
	}).Methods("POST")
}

func handleFeedPasue(w http.ResponseWriter, r *http.Request, session communication.Session) {
	log.Printf("handleFeedPasue %s", r.URL.String())
	var body []byte
	r.Body.Read(body)
	var enable bool
	json.Unmarshal(body, &enable)
	commands.FeedPause(session, enable)
	reply, _ := json.Marshal(true)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(reply)
}

func handleThunderstorm(w http.ResponseWriter, r *http.Request, session communication.Session) {
	log.Printf("handleThunderstorm %s", r.URL.String())
	var body []byte
	r.Body.Read(body)
	var duration int
	json.Unmarshal(body, &duration)
	commands.Thunderstorm(session, duration)
	reply, _ := json.Marshal(true)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(reply)
}

func handleResetReminder(w http.ResponseWriter, r *http.Request, session communication.Session) {
	log.Printf("handleResetReminder %s", r.URL.String())
	vars := mux.Vars(r)
	index, _ := strconv.Atoi(vars["Index"])
	commands.ResetReminder(session, index)
	reply, _ := json.Marshal(true)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(reply)
}
func handleMaintenance(w http.ResponseWriter, r *http.Request, session communication.Session) {
	log.Printf("handleMaintenance %s", r.URL.String())
	vars := mux.Vars(r)
	index, _ := strconv.Atoi(vars["Index"])
	var body []byte
	r.Body.Read(body)
	var enable bool
	json.Unmarshal(body, &enable)
	commands.Maintenance(session, index, enable)
	reply, _ := json.Marshal(true)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(reply)
}
func handleClearLevelAlarm(w http.ResponseWriter, r *http.Request, session communication.Session) {
	log.Printf("handleClearLevelAlarm %s", r.URL.String())
	vars := mux.Vars(r)
	id, _ := vars["ID"]
	commands.ClearLevelAlarm(session, id)
	reply, _ := json.Marshal(true)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(reply)
}
func handleStartWaterChange(w http.ResponseWriter, r *http.Request, session communication.Session) {
	log.Printf("handleThunderstorm %s", r.URL.String())
	vars := mux.Vars(r)
	id, _ := vars["ID"]
	commands.WaterChange(session, id)
	reply, _ := json.Marshal(true)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(reply)
}
