package lexer

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

var operationsMap = map[rune]TokenType{
	'+': Plus,
	'-': Minus,
	'*': Multiply,
	'/': Divide,
	'(': LeftBracket,
	')': RightBracket,
}

func GenerateTokens(expression string) ([]Token, error) {
	var tokens []Token
	var containsDecimalPoint bool
	for i := 0; i < len(expression); i++ {
		character := rune(expression[i])
		if tokenType, ok := operationsMap[character]; ok {
			tokens = append(tokens, NewTokenWithoutValue(tokenType))
		} else if unicode.IsDigit(character) || character == '.' {
			j := i
			for j < len(expression) {
				curr := rune(expression[j])
				if curr == '.' {
					if containsDecimalPoint {
						return nil, errors.New("invalid number format: multiple decimal points")
					}
					containsDecimalPoint = true
				} else if !unicode.IsDigit(curr) {
					break
				}
				j++
			}
			num, err := strconv.ParseFloat(expression[i:j], 64)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, Token{Type: Number, Value: num})
			containsDecimalPoint = false
			i = j - 1
		} else if !unicode.IsSpace(character) {
			return nil, errors.New(fmt.Sprintf("invalid character found: %v", character))
		}
	}
	return tokens, nil
}
