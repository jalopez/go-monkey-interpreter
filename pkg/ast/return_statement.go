package ast

import "github.com/jalopez/go-monkey-interpreter/pkg/token"

// ReturnStatement "return" statement
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (*ReturnStatement) statementNode() {}

// TokenLiteral token literal
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
