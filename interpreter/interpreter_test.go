package interpreter_test

import (
	"go_calculator/interpreter"
	"go_calculator/parser"
	"testing"
)

func TestCalculateResult(t *testing.T) {
	tests := []struct {
		name        string
		node        parser.Node
		expected    float64
		expectError bool
	}{
		{
			name:        "Single number",
			node:        parser.NumberNode{Value: 42},
			expected:    42,
			expectError: false,
		},
		{
			name: "Addition",
			node: parser.AddNode{
				LeftNode:  parser.NumberNode{Value: 10},
				RightNode: parser.NumberNode{Value: 5},
			},
			expected:    15,
			expectError: false,
		},
		{
			name: "Subtraction",
			node: parser.SubtractNode{
				LeftNode:  parser.NumberNode{Value: 10},
				RightNode: parser.NumberNode{Value: 5},
			},
			expected:    5,
			expectError: false,
		},
		{
			name: "Multiplication",
			node: parser.MultiplyNode{
				LeftNoe:   parser.NumberNode{Value: 6},
				RightNode: parser.NumberNode{Value: 7},
			},
			expected:    42,
			expectError: false,
		},
		{
			name: "Division",
			node: parser.DivideNode{
				LeftNode:  parser.NumberNode{Value: 10},
				RightNode: parser.NumberNode{Value: 2},
			},
			expected:    5,
			expectError: false,
		},
		{
			name: "Division by zero",
			node: parser.DivideNode{
				LeftNode:  parser.NumberNode{Value: 10},
				RightNode: parser.NumberNode{Value: 0},
			},
			expected:    0,
			expectError: true,
		},
		{
			name: "Unary plus",
			node: parser.UnaryPlusNode{
				Node: parser.NumberNode{Value: 5},
			},
			expected:    5,
			expectError: false,
		},
		{
			name: "Unary minus",
			node: parser.UnaryMinusNode{
				Node: parser.NumberNode{Value: 5},
			},
			expected:    -5,
			expectError: false,
		},
		{
			name: "Complex expression: (5 + 3) * -2",
			node: parser.MultiplyNode{
				LeftNoe: parser.AddNode{
					LeftNode:  parser.NumberNode{Value: 5},
					RightNode: parser.NumberNode{Value: 3},
				},
				RightNode: parser.UnaryMinusNode{
					Node: parser.NumberNode{Value: 2},
				},
			},
			expected:    -16,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := interpreter.CalculateResult(tt.node)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect error but got one: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Expected %v but got %v", tt.expected, result)
				}
			}
		})
	}
}
