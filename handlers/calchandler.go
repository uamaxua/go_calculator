package handlers

import (
	"fmt"
	"go_calculator/lexer"
	"log"
	"net/http"
)

func CalculateExpression(w http.ResponseWriter, r *http.Request) {
	expression := r.URL.Query().Get("expression")
	if isNotValidRequestParam(w, expression) {
		return
	}
	log.Printf("Calculating expression: %s", expression)
	tokens, lexer_error := lexer.GenerateTokens(expression)
	if lexer_error != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if tokens != nil {
		fmt.Printf("Tokens: %s\n", tokens)
	}
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
