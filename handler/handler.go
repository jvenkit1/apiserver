package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/bitly/go-simplejson"
)

var Counter uint64

func init() {
	Counter = 0
}

// get request handler
func GetHandler(w http.ResponseWriter, r *http.Request){
	Counter += 1

	logrus.Infof("GET Counter is %d", Counter)

	json := simplejson.New()
	json.Set("Counter", Counter)

	payload, err := json.MarshalJSON()
	if err != nil {
		logrus.WithError(err).Error("Unable to marshal json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func PostHandler(w http.ResponseWriter, r *http.Request){
	Counter += 1

	logrus.Infof("POST Counter is %d", Counter)

	json := simplejson.New()
	json.Set("Counter", Counter)

	payload, err := json.MarshalJSON()
	if err != nil {
		logrus.WithError(err).Error("Unable to marshal json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
