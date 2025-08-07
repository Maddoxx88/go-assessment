package utils

import (
	"bytes"
	"github.com/jung-kurt/gofpdf"
	"go-service/models"
)

func CreatePDF(student models.Student) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(40, 10, "Student Report")
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "ID: "+student.ID)
	pdf.Ln(8)
	pdf.Cell(40, 10, "Name: "+student.Name)
	pdf.Ln(8)
	pdf.Cell(40, 10, "Email: "+student.Email)
	pdf.Ln(8)
	pdf.Cell(40, 10, "Grade: "+student.Grade)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
