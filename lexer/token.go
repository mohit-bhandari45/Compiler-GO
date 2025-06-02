package lexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"
	CHAR   = "CHAR"
	FLOAT  = "FLOAT"

	// Operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	EQ        = "=="
	NOT_EQ    = "!="
	LT        = "<"
	GT        = ">"
	LTE       = "<="
	GTE       = ">="
	AMPERSAND = "&"
	AND       = "&&"
	OR        = "||"
	PIPE      = "|"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	DOT       = "."
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Keywords
	FUNCTION = "FUNC"
	LET      = "LET"
	CONST    = "CONST"
	VAR      = "VAR"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	FOR      = "FOR"
	WHILE    = "WHILE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	NULL     = "NULL"
)

var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"let":    LET,
	"const":  CONST,
	"var":    VAR,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"for":    FOR,
	"while":  WHILE,
	"true":   TRUE,
	"FALSE":  FALSE,
	"null":   NULL,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
