package parser

import (
	ast "github.com/mohit-bhandari45/Compiler-GO.git/internal/ast"
	lexer "github.com/mohit-bhandari45/Compiler-GO.git/internal/lexer"
)

// precedence levels (larger binds tighter)
const (
	_ int = iota
	LOWEST
	EQUALS      // == !=
	LESSGREATER // < <= > >=
	SUM         // + -
	PRODUCT     // * /
	PREFIX      // -X !X
	GROUP       // ( ... )
)

var Precedences = map[lexer.TokenType]int{
	lexer.EQ:       EQUALS,
	lexer.NOT_EQ:   EQUALS,
	lexer.LT:       LESSGREATER,
	lexer.GT:       LESSGREATER,
	lexer.LTE:      LESSGREATER,
	lexer.GTE:      LESSGREATER,
	lexer.PLUS:     SUM,
	lexer.MINUS:    SUM,
	lexer.SLASH:    PRODUCT,
	lexer.ASTERISK: PRODUCT,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression
)

type Parser struct {
	l *lexer.Lexer

	curToken  lexer.Token
	peekToken lexer.Token

	errors []string

	prefixFns map[lexer.TokenType]prefixParseFn;
	infixFns map[lexer.TokenType]infixParseFn;
}

func New(l *lexer.Lexer) *Parser{
	p := &Parser{
		l: l,

		errors: []string{},
		prefixFns: make(map[lexer.TokenType]prefixParseFn),
		infixFns: make(map[lexer.TokenType]infixParseFn),
	}

	// read two tokens, set curr and peek
	p.nextToken()
	p.nextToken()

	// register prefix parse functions
	p.registerPrefix(lexer.IDENT, p.parseIdentifier);
	p.registerPrefix(lexer.INT, p.parseIntegerLiteral)
	p.registerPrefix(lexer.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(lexer.STRING, p.parseStringLiteral)
	// p.registerPrefix(lexer.BANG, p.parsePrefixExpression)
	// p.registerPrefix(lexer.MINUS, p.parsePrefixExpression)
	// p.registerPrefix(lexer.LPAREN, p.parseGroupedExpression)

	return p;
}

func (p *Parser) nextToken(){
	p.curToken = p.peekToken;
	p.peekToken = p.l.NextToken();
}

func (p *Parser) registerPrefix(tt lexer.TokenType, fn prefixParseFn){
	p.prefixFns[tt] = fn;
}

// prefix parse functions
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	return &ast.IntegerLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	return &ast.FloatLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}