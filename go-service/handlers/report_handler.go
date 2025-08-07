package handlers

import (
	"bytes"
	"github.com/gorilla/mux"
	"net/http"
)

func GenerateStudentReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["id"]

	student, err := fetchStudentData(studentID)
	if err != nil {
		http.Error(w, "Failed to fetch student data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	pdfBytes, err := createPDF(student)
	if err != nil {
		http.Error(w, "Failed to create PDF: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=report.pdf")
	w.WriteHeader(http.StatusOK)
	w.Write(pdfBytes)
}
