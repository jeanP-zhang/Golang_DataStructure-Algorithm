package Calc

type Lexer struct {
	input        string //输入字符串
	position     int    //位置
	readPosition int    //读取位置
	ch           byte   //读取一个字节
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar() //前进一个字符
	return l
}

//分离数据和操作
func (l *Lexer) NextToken() Token {
	var tok Token
	l.SkipWhiteSpace() //跳过垃圾字符
	switch l.ch {
	case '(':
		tok = NewToken(LPAREN, l.ch)
	case ')':
		tok = NewToken(RPAREN, l.ch)
	case '+':
		tok = NewToken(PLUS, l.ch)
	case '-':
		tok = NewToken(MINUS, l.ch)
	case '*':
		tok = NewToken(ASTERISK, l.ch)
	case '/':
		tok = NewToken(SLASH, l.ch)
	//case '0':
	//	tok.Literal = ""
	//	tok.Type = EOF //关闭
	case '!':
		tok = NewToken(G0, l.ch)
	case '>':
		tok = NewToken(BIG, l.ch)
	default:
		if IsDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.ReadNumber()
			return tok
		} else {
			tok = NewToken(ILLEGAL, l.ch)
		}
	}
	l.ReadChar()
	return tok
}
func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

//切割数字
func (l *Lexer) ReadNumber() string {
	positions := l.position
	for IsDigit(l.ch) {
		l.ReadChar()
	}
	return l.input[positions:l.position]
}
func (l *Lexer) SkipWhiteSpace() {
	if l.ch == '\t' || l.ch == ' ' || l.ch == '\r' || l.ch == '\n' {
		l.ReadChar()
	}
}
