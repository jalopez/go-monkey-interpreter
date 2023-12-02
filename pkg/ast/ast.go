package ast

// Node node
type Node interface {
	TokenLiteral() string
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
