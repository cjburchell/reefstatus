package routes

import (
	"encoding/json"
	"net/http"

	"github.com/cjburchell/go-uatu"

	"github.com/cjburchell/reefstatus-data/associations"

	"github.com/cjburchell/reefstatus-data/repo"

	"github.com/cjburchell/reefstatus-data/settings"
	"github.com/gorilla/mux"
)

type crud interface {
	Get(id string) (interface{}, error)
	Delete(id string) error
	Set(interface{}) error
	GetList() (interface{}, error)
	GetEmpty() interface{}
}

func setupCrud(path string, r *mux.Router, resource crud) {
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
		handleGet(writer, request.URL.String(), resourceItems, err)
	}).Methods("GET")
	r.HandleFunc(path+"/{id}", func(writer http.ResponseWriter, request *http.Request) {
		resourceItem, err := resource.Get(mux.Vars(request)["id"])
		handleGet(writer, request.URL.String(), resourceItem, err)
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
func SetupControllerRoute(r *mux.Router, c repo.Controller) {
	controllerRoute := r.PathPrefix("/controller").Subrouter()
	controllerRoute.Use(tokenMiddleware)
	controllerRoute.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		handleInfo(writer, request, c)
	}).Methods("GET")
	controllerRoute.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		handleSetInfo(writer, request, c)
	}).Methods("PUT")
	controllerRoute.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		handleSetInfo(writer, request, c)
	}).Methods("POST")

	setupCrud("/probe", controllerRoute, probe{c})
	setupCrud("/levelsensor", controllerRoute, levelsensor{c})
	setupCrud("/sport", controllerRoute, sport{c})
	setupCrud("/lport", controllerRoute, lport{c})
	setupCrud("/digitalinput", controllerRoute, digitalinput{c})
	setupCrud("/pump", controllerRoute, pump{c})
	setupCrud("/programmablelogic", controllerRoute, programmablelogic{c})
	setupCrud("/dosingpump", controllerRoute, dosingpump{c})
	setupCrud("/light", controllerRoute, light{c})

	controllerRoute.HandleFunc("/updateAssociations", func(writer http.ResponseWriter, request *http.Request) {
		log.Debugf("Controller Update Associations Get %s", request.URL.String())
		associations.Update(c)
		writer.WriteHeader(http.StatusOK)
	}).Methods("POST")
}

func tokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		auth := request.Header.Get("Authorization")
		if auth != "APIKEY "+settings.DataServiceToken {
			response.WriteHeader(http.StatusUnauthorized)

			log.Warnf("Unauthorized %s != %s", auth, settings.DataServiceToken)
			return
		}

		next.ServeHTTP(response, request)
	})
}

func handleGet(response http.ResponseWriter, url string, result interface{}, err error) {
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
	response.Write(reply)
}
