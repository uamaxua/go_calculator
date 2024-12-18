package lexer

import (
	"testing"
)

func TestGenerateTokens(t *testing.T) {
	tests := []struct {
		name        string
		expression  string
		expected    []Token
		expectError bool
	}{
		{
			name:       "Basic operators",
			expression: "+-*/()",
			expected: []Token{
				NewTokenWithoutValue(Plus),
				NewTokenWithoutValue(Minus),
				NewTokenWithoutValue(Multiply),
				NewTokenWithoutValue(Divide),
				NewTokenWithoutValue(LeftBracket),
				NewTokenWithoutValue(RightBracket),
			},
			expectError: false,
		},
		{
			name:       "Simple expression",
			expression: "2 + 3 * 9.7",
			expected: []Token{
				NewToken(Number, 2.000000),
				NewTokenWithoutValue(Plus),
				NewToken(Number, 3.000000),
				NewTokenWithoutValue(Multiply),
				NewToken(Number, 9.700000),
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens, err := GenerateTokens(tt.expression)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error but got one: %v", err)
				}
				if !tokensEqualIgnoringNaN(tt.expected, tokens) {
					t.Errorf("Expected tokens %v, but got %v", tt.expected, tokens)
				}
			}
		})
	}
}

func tokensEqualIgnoringNaN(expected, actual []Token) bool {
	if len(expected) != len(actual) {
		return false
	}
	for i := range expected {
		if expected[i].Type != actual[i].Type {
			return false
		}
		if !(isNaN(expected[i].Value) && isNaN(actual[i].Value)) {
			if expected[i].Value != actual[i].Value {
				return false
			}
		}
	}
	return true
}

func isNaN(value float64) bool {
	return value != value
}
