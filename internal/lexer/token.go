// internal/lexer/token.go
package lexer;

type TokenType string;

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special
	ILLEGAL TokenType = "illegal"
	EOF     TokenType = "eof"

	// Identifiers + literals
	IDENT  TokenType = "ident"  // for int, int8, string, a, name ...etc
	INT    TokenType = "int"  // for 34, 45, 23 ...etc
	FLOAT  TokenType = "float"  // for 12.3, 52.3 ..etc
	STRING TokenType = "string"  // for "mohit", "right" ...etc

	// Operators
	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	BANG     TokenType = "!"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"
	LT       TokenType = "<"
	GT       TokenType = ">"
	EQ       TokenType = "=="
	NOT_EQ   TokenType = "!="
	LTE      TokenType = "<="
	GTE      TokenType = ">="

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	COLON     TokenType = ":"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"
	LBRACKET  TokenType = "["
	RBRACKET  TokenType = "]"

	// Keywords
	BREAK    TokenType = "break"
	CASE     TokenType = "case"
	CHAN     TokenType = "chan"
	CONST    TokenType = "const"
	CONTINUE TokenType = "continue"

	DEFAULT     TokenType = "default"
	DEFER       TokenType = "defer"
	ELSE        TokenType = "else"
	FALLTHROUGH TokenType = "fallthrough"
	FOR         TokenType = "for"

	FUNC   TokenType = "func"
	GO     TokenType = "go"
	GOTO   TokenType = "goto"
	IF     TokenType = "if"
	IMPORT TokenType = "import"

	INTERFACE TokenType = "interface"
	MAP       TokenType = "map"
	PACKAGE   TokenType = "package"
	RANGE     TokenType = "range"
	RETURN    TokenType = "return"

	SELECT TokenType = "select"
	STRUCT TokenType = "struct"
	SWITCH TokenType = "switch"
	TYPE   TokenType = "type"
	VAR    TokenType = "var"
)

var keywords = map[string]TokenType{
	"break":       BREAK,
	"case":        CASE,
	"chan":        CHAN,
	"const":       CONST,
	"continue":    CONTINUE,
	"default":     DEFAULT,
	"defer":       DEFER,
	"else":        ELSE,
	"fallthrough": FALLTHROUGH,
	"for":         FOR,
	"func":        FUNC,
	"go":          GO,
	"goto":        GOTO,
	"if":          IF,
	"import":      IMPORT,
	"interface":   INTERFACE,
	"map":         MAP,
	"package":     PACKAGE,
	"range":       RANGE,
	"return":      RETURN,
	"select":      SELECT,
	"struct":      STRUCT,
	"switch":      SWITCH,
	"type":        TYPE,
	"var":         VAR,
}

func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}