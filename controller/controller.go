package controller

import (
	"encoding/json"
	"md-tnt-mgmt/iface"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Controller struct {
	svc iface.Service
	log *logrus.Logger
}

func New(svc iface.Service, log *logrus.Logger) *Controller {
	return &Controller{
		svc: svc,
		log: log,
	}
}

// encodeJSONResponse
// encodes the given data into a JSON response and writes it to the provided http.ResponseWriter.
func encodeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
