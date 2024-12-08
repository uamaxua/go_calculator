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
				NewTokenWithoutValue(PLUS),
				NewTokenWithoutValue(MINUS),
				NewTokenWithoutValue(MULTIPLY),
				NewTokenWithoutValue(DIVIDE),
				NewTokenWithoutValue(BRACKET_LEFT),
				NewTokenWithoutValue(BRACKET_RIGHT),
			},
			expectError: false,
		},
		{
			name:       "Simple expression",
			expression: "2 + 3 * 9.7",
			expected: []Token{
				NewToken(NUMBER, 2.000000),
				NewTokenWithoutValue(PLUS),
				NewToken(NUMBER, 3.000000),
				NewTokenWithoutValue(MULTIPLY),
				NewToken(NUMBER, 9.700000),
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
