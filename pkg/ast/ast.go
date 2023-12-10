package ast

// Node node
type Node interface {
	TokenLiteral() string
	String() string
	ToJSON() string
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
	out := ""

	for _, s := range p.Statements {
		out += s.String()
	}

	return out
}

// ToJSON to json
func (p *Program) ToJSON() string {
	out := "["

	for _, s := range p.Statements {
		out += s.ToJSON()
		out += ","
	}
	out = out[:len(out)-1]
	out += "]"

	return out
}
