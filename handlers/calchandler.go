package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func CalculateExpression(w http.ResponseWriter, r *http.Request) {
	expression := r.URL.Query().Get("expression")
	if isNotValidRequestParam(w, expression) {
		return
	}
	log.Printf("Calculating expression: %s", expression)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf("Provided expression: %s", expression)))
	if err != nil {
		log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func isNotValidRequestParam(w http.ResponseWriter, e string) bool {
	if e == "" {
		log.Println("No expression provided")
		http.Error(w, "Bad Request: Missing expression parameter", http.StatusBadRequest)
		return true
	}
	return false
}
