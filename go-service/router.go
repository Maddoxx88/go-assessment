package main

import (
	"github.com/gorilla/mux"
	"go-service/handlers"
	"net/http"
)

func InitializeRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods("GET")

	r.HandleFunc("/api/v1/students/{id}/report", handlers.GenerateStudentReport).Methods("GET")

	return r
}
