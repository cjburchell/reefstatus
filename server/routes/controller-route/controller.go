package controller_route

import (
	"encoding/json"
	"net/http"

	"github.com/cjburchell/reefstatus/server/routes/token"

	"github.com/cjburchell/reefstatus/server/data/associations"
	"github.com/cjburchell/reefstatus/server/data/repo"
	logger "github.com/cjburchell/uatu-go"

	"github.com/gorilla/mux"
)

type crud interface {
	Get(id string) (interface{}, error)
	Delete(id string) error
	Set(interface{}) error
	GetList() (interface{}, error)
	GetEmpty() interface{}
}

func setupCrud(path string, r *mux.Router, resource crud, log logger.ILog) {
	put := func(writer http.ResponseWriter, request *http.Request) {
		item := resource.GetEmpty()
		err := json.NewDecoder(request.Body).Decode(item)
		if err != nil {
			log.Errorf(err, "Unmarshal Failed %s", request.URL.String())
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		err = resource.Set(item)
		if err != nil {
			log.Errorf(err, "Unable to Set resource %s", request.URL.String())
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusNoContent)
	}

	r.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		resourceItems, err := resource.GetList()
		handleGet(writer, request.URL.String(), resourceItems, err, log)
	}).Methods("GET")
	r.HandleFunc(path+"/{id}", func(writer http.ResponseWriter, request *http.Request) {
		resourceItem, err := resource.Get(mux.Vars(request)["id"])
		handleGet(writer, request.URL.String(), resourceItem, err, log)
	}).Methods("GET")
	r.HandleFunc(path, put).Methods("POST")
	r.HandleFunc(path+"/{id}", func(writer http.ResponseWriter, request *http.Request) {
		id := mux.Vars(request)["id"]
		_, err := resource.Get(id)
		if err != nil {
			log.Errorf(err, "Unable to PUT Resource %s", request.URL.String())
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		put(writer, request)
	}).Methods("PUT")
	r.HandleFunc(path+"/{id}", func(writer http.ResponseWriter, request *http.Request) {
		id := mux.Vars(request)["id"]
		_, err := resource.Get(id)
		if err != nil {
			log.Errorf(err, "Unable to DELETE Resource %s", request.URL.String())
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		err = resource.Delete(id)
		if err != nil {
			log.Errorf(err, "Unable to DELETE Resource %s", request.URL.String())
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusNoContent)
	}).Methods("DELETE")
}

// SetupControllerRoute setup the route
func SetupRoute(r *mux.Router, c repo.Controller, log logger.ILog) {
	controllerRoute := r.PathPrefix("api/v1/controller").Subrouter()
	controllerRoute.Use(token.Middleware)
	controllerRoute.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		handleInfo(writer, request, c, log)
	}).Methods("GET")
	controllerRoute.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		handleSetInfo(writer, request, c, log)
	}).Methods("PUT")
	controllerRoute.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		handleSetInfo(writer, request, c, log)
	}).Methods("POST")

	setupCrud("/probe", controllerRoute, probe{c}, log)
	setupCrud("/levelsensor", controllerRoute, levelsensor{c}, log)
	setupCrud("/sport", controllerRoute, sport{c}, log)
	setupCrud("/lport", controllerRoute, lport{c}, log)
	setupCrud("/digitalinput", controllerRoute, digitalinput{c}, log)
	setupCrud("/pump", controllerRoute, pump{c}, log)
	setupCrud("/programmablelogic", controllerRoute, programmablelogic{c}, log)
	setupCrud("/dosingpump", controllerRoute, dosingpump{c}, log)
	setupCrud("/light", controllerRoute, light{c}, log)

	controllerRoute.HandleFunc("/updateAssociations", func(writer http.ResponseWriter, request *http.Request) {
		log.Debugf("Controller Update Associations Get %s", request.URL.String())
		associations.Update(c)
		writer.WriteHeader(http.StatusOK)
	}).Methods("POST")
}

func handleGet(response http.ResponseWriter, url string, result interface{}, err error, log logger.ILog) {
	log.Debugf("Controller Handle Get %s", url)

	if err == repo.ErrNotFound {
		log.Debugf("Unable to find resource: %s", url)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		log.Errorf(err, "Error handling controller data %s", url)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	reply, _ := json.Marshal(result)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	_, err = response.Write(reply)
	if err != nil {
		log.Error(err)
	}
}
