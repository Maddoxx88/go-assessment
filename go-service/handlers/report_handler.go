package handlers

import (
	"github.com/gorilla/mux"
	"go-service/services"
	"net/http"
)

func GenerateStudentReport(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	studentID := params["id"]

	err := services.GeneratePDFReport(studentID, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
