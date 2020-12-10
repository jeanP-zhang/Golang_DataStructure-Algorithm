package Calc

type Token struct {
	Type    string //类型
	Literal string //意义
}

func NewToken(tokenType string, c byte) Token {
	return Token{tokenType, string(c)}
}
