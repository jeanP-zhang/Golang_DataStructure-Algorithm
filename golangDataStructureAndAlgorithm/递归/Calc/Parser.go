package Calc

import (
	"fmt"
	"strconv"
)

type (
	prefixParseFn func() Expression
	infixParseFn  func(Expression) Expression
)
type Parser struct {
	l              *Lexer
	curToken       Token //当前的
	peekToken      Token //提取的
	prefixParseFns map[string]prefixParseFn
	infixParseFns  map[string]infixParseFn
	errors         []string //处理错误
}

//map插入数据
func (p *Parser) RegisterPrefix(tokenType string, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) RegisterInfix(tokenType string, in infixParseFn) {
	p.infixParseFns[tokenType] = in
}
func NewParser(l *Lexer) *Parser {
	p := Parser{l: l, errors: []string{}}
	p.prefixParseFns = make(map[string]prefixParseFn)
	p.RegisterPrefix(INT, p.parseIntergerLiteral)
	p.RegisterPrefix(MINUS, p.parsePrefixExpression)
	p.RegisterPrefix(PLUS, p.parsePrefixExpression)
	p.RegisterPrefix(LPAREN, p.ParseGroupExpression)
	p.RegisterPrefix(G0, p.parsePrefixExpression)
	//p.RegisterInfix(PLUS, p.p) //调用函数
	p.infixParseFns = make(map[string]infixParseFn)
	p.RegisterInfix(PLUS, p.ParseInfixExpression)
	p.RegisterInfix(MINUS, p.ParseInfixExpression)
	p.RegisterInfix(SLASH, p.ParseInfixExpression)
	p.RegisterInfix(MOD, p.ParseInfixExpression)
	p.RegisterInfix(ASTERISK, p.ParseInfixExpression)
	p.RegisterInfix(BIG, p.ParseInfixExpression)
	p.nextToken()
	p.nextToken()
	return &p
}
func (p *Parser) Errors() []string {
	return p.errors //返回错误
}
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken() //循环截取数据
}
func (p *Parser) peekError(s string) {
	msg := fmt.Sprintf("bug %s %s", s, p.peekToken)
	p.errors = append(p.errors, msg)
}

//取出优先级
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	} else {
		return LOWEST
	}
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	} else {
		return LOWEST
	}
}

//提取判断
func (p *Parser) peekTokenIs(t string) bool {
	return p.peekToken.Type == t
}

//处理字符正确或者错误
func (p *Parser) expectPeek(t string) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
func (p *Parser) ParseExpression(precedence int) Expression {
	prefix := p.prefixParseFns[p.curToken.Type] //计算函数
	returnExp := prefix()
	for precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type] //前缀
		if infix == nil {
			return returnExp
		}
		p.nextToken()
		returnExp = infix(returnExp) //循环取出数据与函数
	}
	return returnExp
}

//1+ -1
func (p *Parser) parsePrefixExpression() Expression {
	expression := &PreFixExpression{Toke: p.curToken, Operator: p.curToken.Literal}
	p.nextToken()
	expression.Right = p.ParseExpression(PREFIX) //处理右边
	return expression
}
func (p *Parser) parseIntergerLiteral() Expression {
	lit := &IntergerLiteralExpression{Tokens: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("bug %s int %s")
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}
func (p *Parser) ParseGroupExpression() Expression {
	p.nextToken()
	exp := p.ParseExpression(LOWEST) //设置级别最低
	if !p.expectPeek(RPAREN) {
		return nil
	}
	return exp
}
func (p *Parser) ParseInfixExpression(left Expression) Expression {
	expressions := &InfixExpression{Toke: p.curToken, operator: p.curToken.Literal, Left: left}
	precedences := p.curPrecedence()
	p.nextToken() //循环继续
	expressions.Right = p.ParseExpression(precedences)
	return expressions
}
