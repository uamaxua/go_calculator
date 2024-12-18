package handlers

import (
	"fmt"
	"go_calculator/interpreter"
	"go_calculator/lexer"
	"go_calculator/parser"
	"log"
	"net/http"
)

func CalculateExpression(w http.ResponseWriter, r *http.Request) {
	expression := r.URL.Query().Get("expression")
	if expression == "" {
		log.Println("No expression provided")
		http.Error(w, "Bad Request: Missing expression parameter", http.StatusBadRequest)
		return
	}

	log.Printf("Calculating expression: %s", expression)

	tokens, lexerError := lexer.GenerateTokens(expression)
	if handleError(w, lexerError, "Not able to tokenize", http.StatusInternalServerError) {
		return
	}
	log.Printf("Tokens: %s\n", tokens)

	mathParser := parser.NewMathParser(tokens)
	node, parseError := mathParser.Parse()
	if handleError(w, parseError, "Not able to parse", http.StatusInternalServerError) {
		return
	}
	log.Printf("Flat tree: %s\n", node.String())

	result, interpreterError := interpreter.CalculateResult(node)
	if handleError(w, interpreterError, "Not able to calculate result", http.StatusInternalServerError) {
		return
	}

	w.WriteHeader(http.StatusOK)
	_, writeError := w.Write([]byte(fmt.Sprintf("Result: %.2f", result)))
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
