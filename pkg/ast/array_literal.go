package ast

import "github.com/jalopez/go-monkey-interpreter/pkg/token"

// ArrayLiteral literal with array value
type ArrayLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

func (*ArrayLiteral) expressionNode() {} //nolint:golint,unused

// TokenLiteral token literal
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }

// String string representation
func (al *ArrayLiteral) String() string {
	out := "["
	for _, e := range al.Elements {
		out += e.String()
		out += ","
	}
	out = out[:len(out)-1]
	out += "]"
	return out
}

// ToJSON to json
func (al *ArrayLiteral) ToJSON() string {
	out := `{"type":"array","value":[`
	for _, e := range al.Elements {
		out += e.ToJSON()
		out += ","
	}
	out = out[:len(out)-1]
	out += `]}`
	return out
}
