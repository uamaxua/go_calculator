package main

import (
	"fmt"
	"go_calculator/handlers"
	"log"
	"net/http"
)

func main() {
	handleRequests()
}

func handleRequests() {
	http.Handle("/calc", http.HandlerFunc(handlers.CalculateExpression))
	port := 8080
	log.Println("HTTP server is starting on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
