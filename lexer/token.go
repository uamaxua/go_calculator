package lexer

import (
	"fmt"
	"math"
)

type TokenType int

const (
	NUMBER TokenType = iota
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	BRACKET_LEFT
	BRACKET_RIGHT
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
	return fmt.Sprintf("Token {type: %d, value: %f}", t.Type, t.Value)
}
