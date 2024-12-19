package calculator

import (
	"fmt"
	"go_calculator/interpreter"
	"go_calculator/lexer"
	"go_calculator/parser"
	"log"
	"math"
)

func Calculate(exp string) (float64, error) {
	if exp == "" {
		return math.NaN(), fmt.Errorf("no expression provided")
	}
	log.Printf("Calculating expression: %s", exp)

	tokens, lexerError := lexer.GenerateTokens(exp)
	if lexerError != nil {
		return math.NaN(), fmt.Errorf("not able to tokenize: %v", lexerError)
	}
	log.Printf("Tokens: %s\n", tokens)

	mathParser := parser.NewMathParser(tokens)
	node, parseError := mathParser.Parse()
	if parseError != nil {
		return math.NaN(), fmt.Errorf("not able to parse: %v", parseError)
	}
	log.Printf("Flat tree: %s\n", node.String())

	result, interpreterError := interpreter.CalculateResult(node)
	if interpreterError != nil {
		return math.NaN(), fmt.Errorf("not able to calculate result: %v", interpreterError)
	}

	log.Printf("Calculation result: %.2f", result)
	return result, nil
}
