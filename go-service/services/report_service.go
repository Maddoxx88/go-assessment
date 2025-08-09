package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-service/models"
	"go-service/utils"
	"net/http"
)

func GeneratePDFReport(id string, w http.ResponseWriter) error {
	tokens, err := utils.GetTokens()
	fmt.Println("Current tokens:")
	fmt.Println("accessToken:", tokens["accessToken"])
	fmt.Println("refreshToken:", tokens["refreshToken"])
	fmt.Println("csrfToken:", tokens["csrfToken"])
	if err != nil {
		return fmt.Errorf("could not get tokens: %v", err)
	}

	err = fetchAndGeneratePDF(id, w, tokens["accessToken"], tokens["refreshToken"], tokens["csrfToken"])
	if errors.Is(err, utils.ErrUnauthorized) {
		// Token expired → re-login
		newTokens, loginErr := utils.LoginAndGetTokens()
		if loginErr != nil {
			return fmt.Errorf("login failed: %v", loginErr)
		}
		utils.SetTokens(newTokens["accessToken"], newTokens["refreshToken"], newTokens["XSRF-TOKEN"])
		return fetchAndGeneratePDF(id, w, newTokens["accessToken"], newTokens["refreshToken"], newTokens["XSRF-TOKEN"])
	}
	return err
}

func fetchAndGeneratePDF(id string, w http.ResponseWriter, accessToken, refreshToken, csrfToken string) error {
	req, err := http.NewRequest("GET", "http://localhost:5007/api/v1/students/"+id, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Cookie", fmt.Sprintf("accessToken=%s;refreshToken=%s;XSRF-TOKEN=%s", accessToken, refreshToken, csrfToken))
	req.Header.Set("x-csrf-Token", csrfToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error fetching student data: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return utils.ErrUnauthorized
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 status: %d", res.StatusCode)
	}

	var resp models.StudentResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	if !resp.Success {
		return fmt.Errorf("API responded with success=false")
	}

	pdfBytes, err := utils.CreatePDF(resp.Data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=report.pdf")
	w.Write(pdfBytes)
	return nil
}
