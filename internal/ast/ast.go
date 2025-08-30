package ast

import (
	"bytes"
	"strings"

	lexer "github.com/mohit-bhandari45/Compiler-GO.git/internal/lexer"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// ------- Expressions -------- //

// Identifier represents variable names, function names, or any user-defined identifiers
// in the source code (e.g., x, y, foo, bar).
// Note: This does NOT include string literals like "Mohit" or "Rohit".
type Identifier struct {
	Token lexer.Token // The token corresponding to this identifier (Type: IDENT)
	Value string      // The name of the identifier as a string
}

// Marks this node as an Expression (required by the Expression interface)
func (i *Identifier) expressionNode() {}

// Returns the literal value of the token (as it appeared in the source code)
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// Returns a string representation of the identifier (useful for printing the AST)
func (i *Identifier) String() string {
	return i.Value
}

// ---------------------------------------------------------------------------- //

// IntegerLiteral represents integer numbers in the source code (e.g., 5, 42, 100).
type IntegerLiteral struct {
	Token lexer.Token // The token corresponding to this integer (Type: INT)
	Value string      // The literal value of the integer as a string
}

// Marks this node as an Expression (required by the Expression interface)
func (il *IntegerLiteral) expressionNode() {}

// Returns the literal value of the token (as it appeared in the source code)
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// Returns a string representation of the integer (useful for printing the AST)
func (il *IntegerLiteral) String() string {
	return il.Value
}

// ---------------------------------------------------------------------------- //

// FloatLiteral represents floating-point numbers in the source code
// (e.g., 3.14, 0.5, 10.0).
type FloatLiteral struct {
	Token lexer.Token // The token corresponding to this float (Type: FLOAT)
	Value string      // The literal value of the float as a string
}

// Marks this node as an Expression (required by the Expression interface)
func (fl *FloatLiteral) expressionNode() {}

// Returns the literal value of the token as it appeared in the source code
func (fl *FloatLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

// Returns a string representation of the float (useful for printing the AST)
func (fl *FloatLiteral) String() string {
	return fl.Value
}

// ---------------------------------------------------------------------------- //

// StringLiteral represents string values in the source code (e.g., "hello", "world").
// Unlike identifiers, these include quotes in the source but are stored without quotes in Value.
type StringLiteral struct {
	Token lexer.Token // The token corresponding to this string literal (Type: STRING)
	Value string      // The content of the string without quotes
}

// Marks this node as an Expression (required by the Expression interface)
func (sl *StringLiteral) expressionNode() {}

// Returns the literal value of the token as it appeared in the source code (with quotes)
func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

// Returns a string representation of the string literal (useful for printing the AST),
// adding quotes around the Value for readability.
func (sl *StringLiteral) String() string {
	return strconvQuote(sl.Value)
}

// ---------------------------------------------------------------------------- //

// PrefixExpression represents unary operations in the source code,
// such as !foo or -5.
type PrefixExpression struct {
	Token    lexer.Token // The token corresponding to the operator ("!" or "-")
	Operator string      // The operator as a string, e.g., "!" or "-"
	Right    Expression  // The expression the operator is applied to
}

// Marks this node as an Expression (required by the Expression interface)
func (pe *PrefixExpression) expressionNode() {}

// Returns the literal value of the operator token as it appeared in the source code
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// Returns a string representation of the prefix expression (useful for printing the AST)
// Example: (!foo) or (-5)
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

// ---------------------------------------------------------------------------- //

// InfixExpression represents binary operations in the source code,
// such as a + b, a * b, or a == b.
type InfixExpression struct {
	Token    lexer.Token // The token corresponding to the operator (e.g., "+", "*", "==")
	Left     Expression  // The left-hand side expression
	Operator string      // The operator as a string (e.g., "+", "*", "==")
	Right    Expression  // The right-hand side expression
}

// Marks this node as an Expression (required by the Expression interface)
func (ie *InfixExpression) expressionNode() {}

// Returns the literal value of the operator token as it appeared in the source code
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// Returns a string representation of the infix expression (useful for printing the AST)
// Example: (a + b) or (x * y)
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

// ---------------------------------------------------------------------------- //

// BlockStatement represents a block of statements enclosed by braces `{ ... }`
// Commonly used in function bodies, if/else blocks, loops, etc.
type BlockStatement struct {
	Token      lexer.Token // The token corresponding to the opening brace `{`
	Statements []Statement // A slice of statements contained in this block
}
// Marks this node as a Statement (required by the Statement interface)
func (bs *BlockStatement) statementNode() {}
// Returns the literal value of the token as it appeared in the source code
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}
// Returns a string representation of the entire block (useful for printing the AST)
// Concatenates the string representations of all statements in the block
func (bs *BlockStatement) String() string {
	var out strings.Builder
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// ---------------------------------------------------------------------------- //

// helpers
func strconvQuote(s string) string {
	// minimal quoting to make String() readable
	var b bytes.Buffer
	b.WriteByte('"')
	b.WriteString(s)
	b.WriteByte('"')
	return b.String()
}
