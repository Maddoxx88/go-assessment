package main

import (
	"fmt"
	"go-service/utils"
	"log"
	"net/http"
)

func main() {
	r := InitializeRouter()
	// Login at startup and get tokens map
	tokens, err := utils.LoginAndGetTokens()
	if err != nil {
		log.Fatalf("Failed to login at startup: %v", err)
	}

	// Set tokens in utils package for later use
	utils.SetTokens(tokens["accessToken"], tokens["refreshToken"], tokens["csrfToken"])
	fmt.Println("Logged in and tokens stored at startup")

	log.Println("Go service started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
