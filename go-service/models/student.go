package models

import "time"

type Student struct {
	ID                 int        `json:"id"`
	Name               string     `json:"name"`
	Email              string     `json:"email"`
	SystemAccess       bool       `json:"systemAccess"`
	Phone              string     `json:"phone"`
	Gender             string     `json:"gender"`
	Dob                *time.Time `json:"dob"`     // pointer to handle null values
	Class              *string    `json:"class"`   // nullable
	Section            *string    `json:"section"` // nullable
	Roll               *string    `json:"roll"`    // nullable
	FatherName         string     `json:"fatherName"`
	FatherPhone        *string    `json:"fatherPhone"` // nullable
	MotherName         string     `json:"motherName"`
	MotherPhone        *string    `json:"motherPhone"`        // nullable
	GuardianName       *string    `json:"guardianName"`       // nullable
	GuardianPhone      *string    `json:"guardianPhone"`      // nullable
	RelationOfGuardian *string    `json:"relationOfGuardian"` // nullable
	CurrentAddress     *string    `json:"currentAddress"`     // nullable
	PermanentAddress   *string    `json:"permanentAddress"`   // nullable
	AdmissionDate      *string    `json:"admissionDate"`      // nullable, use string or time.Time pointer if format known
	ReporterName       *string    `json:"reporterName"`       // nullable
}

// For decoding the full API response:
type StudentResponse struct {
	Success bool    `json:"success"`
	Data    Student `json:"data"`
}
