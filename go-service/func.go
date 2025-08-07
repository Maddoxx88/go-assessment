import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Class string `json:"class"`
	Marks int    `json:"marks"`
}

func fetchStudentData(id string) (*Student, error) {
	url := fmt.Sprintf("http://localhost:5007/api/v1/students/%s", id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var student Student
	if err := json.Unmarshal(body, &student); err != nil {
		return nil, err
	}

	return &student, nil
}

func createPDF(student *Student) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Student Report")
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Name: %s", student.Name))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Email: %s", student.Email))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Class: %s", student.Class))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Marks: %d", student.Marks))

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	return buf.Bytes(), err
}
