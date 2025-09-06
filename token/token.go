package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IEDNT = "IDENT" // add, x, y, ..
	INT  = "INT" // 0, 9, 2,

	ASSIGN  = "="
	PLUS = "+"
	SUBSTRACT = "-"
	MULTIPLY = "*"

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"
)