package handlers

import (
	"fmt"
	"go_calculator/calculator"
	"log"
	"net/http"
)

func CalculateExpression(w http.ResponseWriter, r *http.Request) {
	expression := r.URL.Query().Get("expression")
	calResult, err := calculator.Calculate(expression)
	if handleError(w, err, "There was an during calculations", http.StatusInternalServerError) {
		return
	}
	w.WriteHeader(http.StatusOK)
	_, writeError := w.Write([]byte(fmt.Sprintf("Result: %.2f", calResult)))
	if handleError(w, writeError, "Failed to write response", http.StatusInternalServerError) {
		return
	}
}

func handleError(w http.ResponseWriter, err error, message string, statusCode int) bool {
	if err != nil {
		log.Printf("%s: %v", message, err)
		http.Error(w, fmt.Sprintf("%s: %v", message, err), statusCode)
		return true
	}
	return false
}
