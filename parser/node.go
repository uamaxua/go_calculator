package parser

import "fmt"

type Node interface {
	String() string
}

type NumberNode struct {
	Value float64
}

func (n NumberNode) String() string {
	return fmt.Sprintf("%v", n.Value)
}

type AddNode struct {
	LeftNode, RightNode Node
}

func (n AddNode) String() string {
	return fmt.Sprintf("|%s+%s|", n.LeftNode.String(), n.RightNode.String())
}

type SubtractNode struct {
	LeftNode, RightNode Node
}

func (n SubtractNode) String() string {
	return fmt.Sprintf("|%s-%s|", n.LeftNode.String(), n.RightNode.String())
}

type MultiplyNode struct {
	LeftNoe, RightNode Node
}

func (n MultiplyNode) String() string {
	return fmt.Sprintf("|%s*%s|", n.LeftNoe.String(), n.RightNode.String())
}

type DivideNode struct {
	LeftNode, RightNode Node
}

func (n DivideNode) String() string {
	return fmt.Sprintf("|%s/%s|", n.LeftNode.String(), n.RightNode.String())
}

type UnaryPlusNode struct {
	Node Node
}

func (n UnaryPlusNode) String() string {
	return fmt.Sprintf("|+%s|", n.Node.String())
}

type UnaryMinusNode struct {
	Node Node
}

func (n UnaryMinusNode) String() string {
	return fmt.Sprintf("|-%s|", n.Node.String())
}
