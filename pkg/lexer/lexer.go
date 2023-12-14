package lexer

import "github.com/jalopez/go-monkey-interpreter/pkg/token"

// Lexer lexer
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	line         int  // current line number
	column       int  // current column number
	ch           byte // current char under examination
}

// New creates a new lexer
func New(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 0,
	}
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
			column := l.column
			l.readChar()
			literal := string(ch) + string(l.ch)
			nextToken = token.Token{
				Type:    token.EQ,
				Literal: literal,
				Line:    l.line,
				Column:  column,
			}
		} else {
			nextToken = newToken(token.ASSIGN, l)
		}
	case ';':
		nextToken = newToken(token.SEMICOLON, l)
	case '(':
		nextToken = newToken(token.LPAREN, l)
	case ')':
		nextToken = newToken(token.RPAREN, l)
	case ',':
		nextToken = newToken(token.COMMA, l)
	case '+':
		nextToken = newToken(token.PLUS, l)
	case '{':
		nextToken = newToken(token.LBRACE, l)
	case '}':
		nextToken = newToken(token.RBRACE, l)
	case '-':
		nextToken = newToken(token.MINUS, l)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			column := l.column
			l.readChar()
			literal := string(ch) + string(l.ch)
			nextToken = token.Token{
				Type:    token.NOTEQ,
				Literal: literal,
				Line:    l.line,
				Column:  column,
			}
		} else {
			nextToken = newToken(token.BANG, l)
		}
	case '/':
		nextToken = newToken(token.SLASH, l)
	case '*':
		nextToken = newToken(token.ASTERISK, l)
	case '<':
		nextToken = newToken(token.LT, l)
	case '>':
		nextToken = newToken(token.GT, l)
	case '"':
		nextToken.Line = l.line
		nextToken.Column = l.column

		str, ok := l.readString()

		if !ok {
			nextToken.Type = token.ILLEGAL
			nextToken.Literal = "Unterminated string"
		}

		nextToken.Type = token.STRING
		nextToken.Literal = str
	case 0:
		nextToken.Literal = ""
		nextToken.Type = token.EOF
	default:
		switch {
		case isLetter(l.ch):
			nextToken.Line = l.line
			nextToken.Column = l.column
			nextToken.Literal = l.readIdentifier()
			nextToken.Type = token.LookupIdent(nextToken.Literal)
			return nextToken
		case isDigit(l.ch):
			nextToken.Line = l.line
			nextToken.Column = l.column

			literal, ok := l.readNumber()

			nextToken.Type = token.INT

			if !ok {
				nextToken.Type = token.ILLEGAL
			}

			nextToken.Literal = literal
			return nextToken
		default:
			nextToken = newToken(token.ILLEGAL, l)
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
	l.column++
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

func (l *Lexer) readString() (string, bool) {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' {
			break
		}

		if l.ch == '\n' || l.ch == 0 {
			return "", false
		}

		if l.ch == '\\' {
			l.readChar()
		}
	}

	return l.input[position:l.position], true
}

func (l *Lexer) readNumber() (string, bool) {
	position := l.position
	for isDigit(l.ch) || isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position], isValidNumber(l.input[position:l.position])
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		if l.ch == '\n' {
			l.line++
			l.column = 0
		}

		l.readChar()
	}
}

func newToken(tokenType token.Type, l *Lexer) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(l.ch),
		Line:    l.line,
		Column:  l.column,
	}
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
