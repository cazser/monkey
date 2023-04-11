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
	INT = "INT"
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH= "/"

	LT = "<"
	GT = ">"
	COMMA = ","
	SEMICOLON = ";"
	LPAREN ="("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	//关键字
	FUNCTION = "FUNCTION"
	LET = "LET"
  TRUE = "true"
	FALSE = "false"
	IF = "if"
	ELSE = "else"
	RETURN = "return"

	EQ = "=="
	NOT_EQ = "!="

	STRING = "STRING"
)


var keywords = map[string] TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType{
	if tok, ok := keywords[ident]; ok{
		return tok;
	}
	return IDENT;
}