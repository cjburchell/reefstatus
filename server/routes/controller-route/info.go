package controller_route

import (
	"encoding/json"
	"net/http"

	"github.com/cjburchell/reefstatus/common/models"
	"github.com/cjburchell/reefstatus/server/data/repo"

	"github.com/cjburchell/go-uatu"
)

func handleInfo(w http.ResponseWriter, r *http.Request, controller repo.Controller) {
	info, err := controller.GetInfo()
	handleGet(w, r.URL.String(), info, err)
}

func handleSetInfo(writer http.ResponseWriter, request *http.Request, controller repo.Controller) {
	log.Debugf("handleSetInfo %s", request.URL.String())

	decoder := json.NewDecoder(request.Body)
	var info models.Info
	err := decoder.Decode(&info)
	if err != nil {
		log.Error(err, "Unmarshal Failed")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.SetInfo(info)
	if err != nil {
		log.Error(err, "Unable to Set Info")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
