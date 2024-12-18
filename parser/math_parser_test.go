package parser

import (
	"go_calculator/lexer"
	"testing"
)

func TestGenerateTokens(t *testing.T) {
	tests := []struct {
		name        string
		tokens      []lexer.Token
		expected    string
		expectError bool
	}{
		{
			name: "Parse simple expression: 2 + 3 * 9.7",
			tokens: []lexer.Token{
				lexer.NewToken(lexer.Number, 2),
				lexer.NewTokenWithoutValue(lexer.Plus),
				lexer.NewToken(lexer.Number, 3),
				lexer.NewTokenWithoutValue(lexer.Multiply),
				lexer.NewToken(lexer.Number, 9.7),
			},
			expected:    "|2+|3*9.7||",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(tt.tokens)
			result, resultError := p.Parse()
			treeString := result.String()
			if tt.expectError {
				if resultError == nil {
					t.Errorf("Expected an error but got none")
				}
			} else {
				if resultError != nil {
					t.Errorf("Did not expect an error but got one: %v", resultError)
				}
				if tt.expected != treeString {
					t.Errorf("Expected nodes represations %v, but got %v", tt.expected, treeString)
				}
			}
		})
	}
}
