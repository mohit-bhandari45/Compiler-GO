package lexer

// lexer.go
// A reasonably-featured lexer for a toy language.
// Exports: Token, TokenType, New (constructor), (*Lexer).NextToken()

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENT  TokenType = "IDENT"  // foo, bar, _name
	INT    TokenType = "INT"    // 123
	FLOAT  TokenType = "FLOAT"  // 1.23
	STRING TokenType = "STRING" // "hello"

	// Operators
	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	BANG     TokenType = "!"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"

	LT     TokenType = "<"
	GT     TokenType = ">"
	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="
	LE     TokenType = "<="
	GE     TokenType = ">="

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	COLON     TokenType = ":"

	LPAREN   TokenType = "("
	RPAREN   TokenType = ")"
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
	WHILE    TokenType = "WHILE"
	FOR      TokenType = "FOR"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"while":  WHILE,
	"for":    FOR,
}

func lookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// Lexer implementation
type Lexer struct {
	input        string
	position     int  // current char index
	readPosition int  // next char index
	ch           byte // current char under examination
}

// New creates a new Lexer for the given input source.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar advances the lexer by one byte (stores into l.ch).
// Uses 0 as EOF sentinel.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
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

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	// handle comments
	if l.ch == '/' && l.peekChar() == '/' {
		l.skipLineComment()
		l.skipWhitespace()
	}
	if l.ch == '/' && l.peekChar() == '*' {
		l.skipBlockComment()
		l.skipWhitespace()
	}

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(PLUS, l.ch)
	case '-':
		tok = newToken(MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(BANG, l.ch)
		}
	case '*':
		tok = newToken(ASTERISK, l.ch)
	case '/':
		tok = newToken(SLASH, l.ch)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: LE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: GE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(GT, l.ch)
		}
	case ',':
		tok = newToken(COMMA, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case ':':
		tok = newToken(COLON, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case '{':
		tok = newToken(LBRACE, l.ch)
	case '}':
		tok = newToken(RBRACE, l.ch)
	case '[':
		tok = newToken(LBRACKET, l.ch)
	case ']':
		tok = newToken(RBRACKET, l.ch)
	case '"':
		// readString consumes the quoted content and moves the lexer past the closing quote
		tok.Type = STRING
		tok.Literal = l.readString()
		return tok
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			lit := l.readIdentifier()
			tok.Type = lookupIdent(lit)
			tok.Literal = lit
			return tok
		} else if isDigit(l.ch) {
			lit, typ := l.readNumber()
			tok.Type = typ
			tok.Literal = lit
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tt TokenType, ch byte) Token {
	return Token{Type: tt, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// readIdentifier reads an identifier starting at current position and returns it.
// It stops at the first char that is not letter/digit/underscore.
func (l *Lexer) readIdentifier() string {
	start := l.position
	// allow digits in identifier after first char (like foo2)
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

// readNumber reads an integer or float and returns the literal and the token type.
func (l *Lexer) readNumber() (string, TokenType) {
	start := l.position
	hasDot := false

	for isDigit(l.ch) || (l.ch == '.' && !hasDot && isDigit(l.peekChar())) {
		if l.ch == '.' {
			hasDot = true
		}
		l.readChar()
	}

	if hasDot {
		return l.input[start:l.position], FLOAT
	}
	return l.input[start:l.position], INT
}

// readString reads a double-quoted string, supports basic escapes like \" and \\.
// It advances the lexer to the char after the closing quote.
func (l *Lexer) readString() string {
	// current l.ch == '"'
	// consume opening quote
	l.readChar()
	start := l.position

	for l.ch != '"' && l.ch != 0 {
		// handle escape by skipping next char (keeps literal as-is)
		if l.ch == '\\' && l.peekChar() != 0 {
			l.readChar() // skip backslash
			l.readChar() // skip escaped char
			continue
		}
		l.readChar()
	}

	lit := l.input[start:l.position]

	// consume closing quote if present
	if l.ch == '"' {
		l.readChar()
	}

	return lit
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) skipLineComment() {
	// assumes l.ch == '/' and peekChar() == '/'
	// Advance until newline or EOF
	l.readChar() // move to second '/'
	l.readChar() // move past second '/'
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

func (l *Lexer) skipBlockComment() {
	// assumes l.ch == '/' and peekChar() == '*'
	l.readChar() // move to '*'
	l.readChar() // move past '*'
	for l.ch != 0 {
		if l.ch == '*' && l.peekChar() == '/' {
			l.readChar() // move to '/'
			l.readChar() // move past '/'
			return
		}
		l.readChar()
	}
	// EOF reached before end of comment -> simply return (unterminated block comment will result in EOF token later)
}