package lexer

import (
	"fmt"
	"math"
)

type TokenType string

const (
	Number       TokenType = "NUMBER"
	Plus         TokenType = "PLUS"
	Minus        TokenType = "MINUS"
	Multiply     TokenType = "MULTIPLY"
	Divide       TokenType = "DIVIDE"
	LeftBracket  TokenType = "LEFT_BRACKET"
	RightBracket TokenType = "RIGHT_BRACKET"
)

type Token struct {
	Type  TokenType
	Value float64
}

func NewToken(tokenType TokenType, value float64) Token {
	return Token{tokenType, value}
}

func NewTokenWithoutValue(tokenType TokenType) Token {
	return Token{tokenType, math.NaN()}
}

func (t Token) String() string {
	return fmt.Sprintf("Token {type: %s, value: %f}", t.Type, t.Value)
}
