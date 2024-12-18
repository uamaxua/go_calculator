package interpreter

import (
	"fmt"
	"go_calculator/parser"
)

func CalculateResult(node parser.Node) (float64, error) {
	switch n := node.(type) {
	case parser.NumberNode:
		return n.Value, nil

	case parser.AddNode:
		left, err := CalculateResult(n.LeftNode)
		if err != nil {
			return 0, err
		}
		right, err := CalculateResult(n.RightNode)
		if err != nil {
			return 0, err
		}
		return left + right, nil

	case parser.SubtractNode:
		left, err := CalculateResult(n.LeftNode)
		if err != nil {
			return 0, err
		}
		right, err := CalculateResult(n.RightNode)
		if err != nil {
			return 0, err
		}
		return left - right, nil

	case parser.MultiplyNode:
		left, err := CalculateResult(n.LeftNoe)
		if err != nil {
			return 0, err
		}
		right, err := CalculateResult(n.RightNode)
		if err != nil {
			return 0, err
		}
		return left * right, nil

	case parser.DivideNode:
		left, err := CalculateResult(n.LeftNode)
		if err != nil {
			return 0, err
		}
		right, err := CalculateResult(n.RightNode)
		if err != nil {
			return 0, err
		}
		if right == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return left / right, nil

	case parser.UnaryPlusNode:
		value, err := CalculateResult(n.Node)
		if err != nil {
			return 0, err
		}
		return +value, nil

	case parser.UnaryMinusNode:
		value, err := CalculateResult(n.Node)
		if err != nil {
			return 0, err
		}
		return -value, nil

	default:
		return 0, fmt.Errorf("unknown node type: %T", n)
	}
}
