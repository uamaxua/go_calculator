package parser

import (
	"errors"
	"go_calculator/lexer"
)

type Parser struct {
	tokens       []lexer.Token
	currentToken *lexer.Token
	position     int
}

func NewParser(tokens []lexer.Token) *Parser {
	parser := &Parser{tokens: tokens, position: -1}
	parser.nextToken()
	return parser
}

func (p *Parser) nextToken() {
	p.position++
	if p.position < len(p.tokens) {
		p.currentToken = &p.tokens[p.position]
	} else {
		p.currentToken = nil
	}
}

func (p *Parser) raiseSyntaxError() error {
	return errors.New("invalid syntax")
}

func (p *Parser) Parse() (Node, error) {
	if p.currentToken == nil {
		return nil, nil
	}
	result, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	if p.currentToken != nil {
		return nil, p.raiseSyntaxError()
	}
	return result, nil
}

func (p *Parser) parseExpression() (Node, error) {
	result, err := p.parseTerm()
	if err != nil {
		return nil, err
	}
	for p.currentToken != nil && (p.currentToken.Type == lexer.Plus || p.currentToken.Type == lexer.Minus) {
		token := p.currentToken
		p.nextToken()
		if token.Type == lexer.Plus {
			right, err := p.parseTerm()
			if err != nil {
				return nil, err
			}
			result = AddNode{LeftNode: result, RightNode: right}
		} else if token.Type == lexer.Minus {
			right, err := p.parseTerm()
			if err != nil {
				return nil, err
			}
			result = SubtractNode{LeftNode: result, RightNode: right}
		}
	}
	return result, nil
}

func (p *Parser) parseTerm() (Node, error) {
	result, err := p.parseFactor()
	if err != nil {
		return nil, err
	}
	for p.currentToken != nil && (p.currentToken.Type == lexer.Multiply || p.currentToken.Type == lexer.Divide) {
		token := p.currentToken
		p.nextToken()
		if token.Type == lexer.Multiply {
			right, err := p.parseFactor()
			if err != nil {
				return nil, err
			}
			result = MultiplyNode{LeftNoe: result, RightNode: right}
		} else if token.Type == lexer.Divide {
			right, err := p.parseFactor()
			if err != nil {
				return nil, err
			}
			result = DivideNode{LeftNode: result, RightNode: right}
		}
	}
	return result, nil
}

func (p *Parser) parseFactor() (Node, error) {
	token := p.currentToken
	if token.Type == lexer.LeftBracket {
		p.nextToken()
		result, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if p.currentToken == nil || p.currentToken.Type != lexer.RightBracket {
			return nil, p.raiseSyntaxError()
		}
		p.nextToken()
		return result, nil
	} else if token.Type == lexer.Number {
		p.nextToken()
		return NumberNode{Value: token.Value}, nil
	} else if token.Type == lexer.Plus {
		p.nextToken()
		node, err := p.parseFactor()
		if err != nil {
			return nil, err
		}
		return UnaryPlusNode{Node: node}, nil
	} else if token.Type == lexer.Minus {
		p.nextToken()
		node, err := p.parseFactor()
		if err != nil {
			return nil, err
		}
		return UnaryMinusNode{Node: node}, nil
	}
	return nil, p.raiseSyntaxError()
}
