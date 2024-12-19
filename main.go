package main

import (
	"fmt"
	"go_calculator/handlers"
	"go_calculator/listeners"
	"log"
	"net/http"
)

func main() {
	go initHttpServer()
	go initNatsListeners()
	select {}
}

func initHttpServer() {
	log.Println("Initializing HTTP server")
	http.Handle("/calc", http.HandlerFunc(handlers.CalculateExpression))
	port := 8080
	log.Println("HTTP server is starting on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func initNatsListeners() {
	log.Println("Initializing NATS listeners")
	listeners.ListenNatsCalcRequests()
}
