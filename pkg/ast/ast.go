package ast

import "bytes"

// Node node
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement statement
type Statement interface {
	Node
	statementNode()
}

// Expression expression
type Expression interface {
	Node
	expressionNode()
}

// Program program
type Program struct {
	Statements []Statement
}

// TokenLiteral token literal
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		_, err := out.WriteString(s.String())

		if err != nil {
			panic(err)
		}
	}

	return out.String()
}
