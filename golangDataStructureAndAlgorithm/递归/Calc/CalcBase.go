package Calc

const (
	ILLEGAL  = "ILLEGAL"
	EOF      = "EOF" //终止
	INT      = "INT"
	PLUS     = "+"
	MINUS    = "-"
	MOD      = "%"
	ASTERISK = "*"
	SLASH    = "/"
	LPAREN   = "("
	RPAREN   = ")"
	BIG      = ">" //5>1 return 1 1>5 return 0
	G0       = "!" //！任何数=0,!0=1
)
const (
	_       = iota
	LOWEST  //级别
	BIGGER  //
	SUM     //+-%
	PRODUCT //*
	PREFIX  //-1
	CALL    //1+(3))
)

//构造集合实现优先级
var precedences = map[string]int{
	PLUS:     SUM,
	MINUS:    SUM,
	MOD:      PRODUCT,
	SLASH:    PRODUCT,
	ASTERISK: PRODUCT,
	LPAREN:   CALL,
	G0:       PREFIX,
	BIG:      BIGGER,
}
