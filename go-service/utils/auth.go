package utils

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

var (
	ErrUnauthorized = errors.New("unauthorized: token expired or invalid")
	accessToken     string
	refreshToken    string
	csrfToken       string
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
		CsrfToken    string `json:"csrfToken"`
	} `json:"data"`
}

func LoginAndGetTokens() (map[string]string, error) {
	loginURL := "http://localhost:5007/api/v1/auth/login"

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	payload := strings.NewReader(`{"username":"admin@school-admin.com", "password":"3OU4zn3q6Zh9"}`)

	req, err := http.NewRequest("POST", loginURL, payload)
	if err != nil {
		return nil, fmt.Errorf("error creating login request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending login request: %v", err)
	}
	defer res.Body.Close()

	u, _ := url.Parse(loginURL)
	cookies := jar.Cookies(u)

	tokens := make(map[string]string)
	for _, c := range cookies {
		switch c.Name {
		case "accessToken":
			tokens["accessToken"] = c.Value
		case "refreshToken":
			tokens["refreshToken"] = c.Value
		case "csrfToken":
			tokens["csrfToken"] = c.Value
		}
	}

	if tokens["accessToken"] == "" || tokens["refreshToken"] == "" || tokens["csrfToken"] == "" {
		return nil, errors.New("tokens not found in login response cookies")
	}

	return tokens, nil
}

func GetTokens() (map[string]string, error) {
	if accessToken == "" || refreshToken == "" || csrfToken == "" {
		return nil, fmt.Errorf("tokens not set")
	}
	return map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"csrfToken":    csrfToken,
	}, nil
}

func SetTokens(a, r, c string) {
	accessToken = a
	refreshToken = r
	csrfToken = c
}
