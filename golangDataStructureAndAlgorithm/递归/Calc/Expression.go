package Calc

import "bytes"

//接口
type Expression interface {
	String() string
}
type PreFixExpression struct {
	Toke     Token
	Operator string
	Right    Expression
}

//整数求值
type IntergerLiteralExpression struct {
	Tokens Token
	Value  int64
}

func (il *IntergerLiteralExpression) String() string {
	return il.Tokens.Literal
}

//对括号内部计算
func (pe *PreFixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Toke     Token
	Left     Expression
	operator string
	Right    Expression
}

func (in *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(in.Left.String())
	out.WriteString("")
	out.WriteString(in.operator)
	out.WriteString("")
	out.WriteString(in.Right.String())
	out.WriteString(")")
	return out.String()
}
