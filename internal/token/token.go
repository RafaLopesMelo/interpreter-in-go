package token

// Could be a int or byte for better performence, but string makes it easier
type TokenType string

const (
	// Token/character that we do not know about
	ILLEGAL TokenType = "ILLEGAL"

	// End of file
	EOF TokenType = "EOF"

	// Identifiers + literals
	IDENT TokenType = "IDENT" // add, foobar, x, y
	INT   TokenType = "INT"

	// Operators
	ASSIGN TokenType = "="
	PLUS   TokenType = "+"

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}