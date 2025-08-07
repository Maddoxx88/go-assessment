package main

import (
	"log"
	"net/http"
)

func main() {
	r := InitializeRouter() // From router.go
	log.Println("PDF microservice running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
