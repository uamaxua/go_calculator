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
	if isNotValidRequestParam(w, expression) {
		return
	}
	log.Printf("Calculating expression: %s", expression)
	tokens, lexerError := lexer.GenerateTokens(expression)
	if lexerError != nil {
		http.Error(w, fmt.Sprintf("Not able to tokenize: %v", lexerError), http.StatusInternalServerError)
	}
	if tokens != nil {
		log.Printf("Tokens: %s\n", tokens)
	}
	mathParser := parser.NewMathParser(tokens)
	node, parseError := mathParser.Parse()
	if parseError != nil {
		http.Error(w, fmt.Sprintf("Not able to parse: %v", parseError), http.StatusInternalServerError)
	}
	if node != nil {
		log.Printf("Flat tree: %s\n", node.String())
	}
	result, interpreterError := interpreter.CalculateResult(node)
	if interpreterError != nil {
		http.Error(w, fmt.Sprintf("Not able to caclulate result: %v", interpreterError), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf("Resul: %.2f", result)))
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
