package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}
const (
	ILLEGAL="ILEGAL"
	EOF = "EOF"
	IDENT = "IDENT"
	ASSIGN = "="
	PLUS = "+"
	COMMA = ","
	SEMICOLON = ";"
	LPAREN ="("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	FUNCTION = "FUNCTION"
	LET = "LET"
)