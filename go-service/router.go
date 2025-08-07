package main

import (
	"go-service/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods("GET")
	r.HandleFunc("/api/v1/students/{id}/report", handlers.GenerateStudentReport).Methods("GET")
	return r
}
