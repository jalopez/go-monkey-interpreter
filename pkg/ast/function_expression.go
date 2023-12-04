package ast

import "github.com/jalopez/go-monkey-interpreter/pkg/token"

// FunctionLiteral function literal
type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (*FunctionLiteral) expressionNode() {} //nolint:golint,unused

// TokenLiteral token literal
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }

// String string representation
func (fl *FunctionLiteral) String() string {
	out := fl.TokenLiteral()
	out += "("

	for i, p := range fl.Parameters {
		if i != 0 {
			out += ", "
		}

		out += p.String()
	}

	out += ") "
	out += fl.Body.String()

	return out
}
