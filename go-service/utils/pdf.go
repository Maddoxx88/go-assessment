package utils

import (
	"bytes"
	"fmt"
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

	// ID as int to string
	pdf.Cell(40, 10, fmt.Sprintf("ID: %d", student.ID))
	pdf.Ln(8)

	pdf.Cell(40, 10, "Name: "+student.Name)
	pdf.Ln(8)

	pdf.Cell(40, 10, "Email: "+student.Email)
	pdf.Ln(8)

	// System Access
	pdf.Cell(40, 10, fmt.Sprintf("System Access: %t", student.SystemAccess))
	pdf.Ln(8)

	pdf.Cell(40, 10, "Phone: "+student.Phone)
	pdf.Ln(8)

	pdf.Cell(40, 10, "Gender: "+student.Gender)
	pdf.Ln(8)

	// Format DOB if present
	if student.Dob != nil {
		pdf.Cell(40, 10, "DOB: "+student.Dob.Format("2006-01-02"))
	} else {
		pdf.Cell(40, 10, "DOB: N/A")
	}
	pdf.Ln(8)

	// Nullable string helper
	printNullable := func(label string, value *string) {
		if value != nil {
			pdf.Cell(40, 10, label+": "+*value)
		} else {
			pdf.Cell(40, 10, label+": N/A")
		}
		pdf.Ln(8)
	}

	printNullable("Class", student.Class)
	printNullable("Section", student.Section)
	printNullable("Roll", student.Roll)

	pdf.Cell(40, 10, "Father Name: "+student.FatherName)
	pdf.Ln(8)

	printNullable("Father Phone", student.FatherPhone)
	pdf.Cell(40, 10, "Mother Name: "+student.MotherName)
	pdf.Ln(8)
	printNullable("Mother Phone", student.MotherPhone)

	printNullable("Guardian Name", student.GuardianName)
	printNullable("Guardian Phone", student.GuardianPhone)
	printNullable("Relation of Guardian", student.RelationOfGuardian)
	printNullable("Current Address", student.CurrentAddress)
	printNullable("Permanent Address", student.PermanentAddress)

	// AdmissionDate as string, print or N/A
	if student.AdmissionDate != nil && *student.AdmissionDate != "" {
		pdf.Cell(40, 10, "Admission Date: "+*student.AdmissionDate)
	} else {
		pdf.Cell(40, 10, "Admission Date: N/A")
	}
	pdf.Ln(8)

	printNullable("Reporter Name", student.ReporterName)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
