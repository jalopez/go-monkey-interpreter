package lexer

import "github.com/jalopez/go-monkey-interpreter/pkg/token"

// Lexer lexer
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New creates a new lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken returns the next token of the input
func (l *Lexer) NextToken() token.Token {
	var nextToken token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			nextToken = token.Token{Type: token.EQ, Literal: literal}
		} else {
			nextToken = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		nextToken = newToken(token.SEMICOLON, l.ch)
	case '(':
		nextToken = newToken(token.LPAREN, l.ch)
	case ')':
		nextToken = newToken(token.RPAREN, l.ch)
	case ',':
		nextToken = newToken(token.COMMA, l.ch)
	case '+':
		nextToken = newToken(token.PLUS, l.ch)
	case '{':
		nextToken = newToken(token.LBRACE, l.ch)
	case '}':
		nextToken = newToken(token.RBRACE, l.ch)
	case '-':
		nextToken = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			nextToken = token.Token{Type: token.NOTEQ, Literal: literal}
		} else {
			nextToken = newToken(token.BANG, l.ch)
		}
	case '/':
		nextToken = newToken(token.SLASH, l.ch)
	case '*':
		nextToken = newToken(token.ASTERISK, l.ch)
	case '<':
		nextToken = newToken(token.LT, l.ch)
	case '>':
		nextToken = newToken(token.GT, l.ch)
	case 0:
		nextToken.Literal = ""
		nextToken.Type = token.EOF
	default:
		switch {
		case isLetter(l.ch):
			nextToken.Literal = l.readIdentifier()
			nextToken.Type = token.LookupIdent(nextToken.Literal)
			return nextToken
		case isDigit(l.ch):
			literal, ok := l.readNumber()

			nextToken.Type = token.INT

			if !ok {
				nextToken.Type = token.ILLEGAL
			}

			nextToken.Literal = literal
			return nextToken
		default:
			nextToken = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return nextToken
}

// readChar reads the next character in the input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL"
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() (string, bool) {
	position := l.position
	for isDigit(l.ch) || isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position], isValidNumber(l.input[position:l.position])
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isValidNumber(literal string) bool {
	for _, ch := range literal {
		if !isDigit(byte(ch)) {
			return false
		}
	}
	return true
}