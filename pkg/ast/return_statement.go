package ast

import "github.com/jalopez/go-monkey-interpreter/pkg/token"

// ReturnStatement "return" statement
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (*ReturnStatement) statementNode() {} //nolint:golint,unused

// TokenLiteral token literal
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	return rs.TokenLiteral() + " " + rs.ReturnValue.String() + ";"
}

// ToJSON to json
func (rs *ReturnStatement) ToJSON() string {
	return `{"type":"return","value":` + rs.ReturnValue.ToJSON() + `}`
}
