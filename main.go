package main

import (
	"apiserver/handler"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/sirupsen/logrus"
)
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/get", handler.GetHandler).Methods(http.MethodGet)
	r.HandleFunc("/post", handler.PostHandler).Methods(http.MethodPost)
	//http.Handle("/", r)
	logrus.Info("Starting Server at port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		logrus.WithError(err).Fatal("Unable to spawn server")
	}
}
