package lexer

import "github.com/RafaLopesMelo/monkey-lang/internal/token"

// Only supports ASCII, since UTF-8 may have multiple bytes per char
type Lexer struct {
	input        string
	readPosition int  // current reading position in the input (after current char)
	position     int  // current position in the input (points to current char)
	ch           byte // current char under examination
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for NUL character
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = l.newToken(token.ASSIGN, '=')
	case ';':
		tok = l.newToken(token.SEMICOLON, ';')
	case '+':
		tok = l.newToken(token.PLUS, '+')
	case '(':
		tok = l.newToken(token.LPAREN, '(')
	case ')':
		tok = l.newToken(token.RPAREN, ')')
	case '{':
		tok = l.newToken(token.LBRACE, '{')
	case '}':
		tok = l.newToken(token.RBRACE, '}')
	case ',':
		tok = l.newToken(token.COMMA, ',')
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	}

	l.readChar()

	return tok
}

func (l *Lexer) newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar()

	return l
}
