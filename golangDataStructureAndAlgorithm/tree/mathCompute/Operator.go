package main

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

const (
	//四则运算符号
	AVALABLE_Code = "+-*/^√()"
	//小数点符号
	AVALABLE_DECIMAL_CODE = "1234567890.E"
	//配套参数
	AVALABLE_PARAMETER_CODE = "abbcdefghijklmnopqrstuvwxyz"
)

//判断计算符
func isAvailableCode(c string) bool {
	return strings.IndexAny(AVALABLE_Code, c) != -1
}
func (op *Operator) GetOpers() []string {
	return op.opers
}
func (op *Operator) GetSuffixExpression() []string {
	return op.suffixExpression
}

//判断小数
func isBelongToDecimal(c string) bool {
	return strings.IndexAny(AVALABLE_DECIMAL_CODE, c) != -1
}

//判断参数
func isParmerterCode(c string) bool {
	return strings.IndexAny(AVALABLE_PARAMETER_CODE, c) != -1
}

//计算
func exeSingleExpression(left float64, right float64, exp string) float64 {
	switch exp {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	case "^":
		return math.Pow(left, right)
	case "√":
		return math.Pow(left, 1/right)
	default:
		return 0
	}
}
func changeParameter(Parameter string, str []string) string {
	for i := 0; i < len(str); i += 2 {
		if Parameter == str[i] {
			return str[i+1]
		}
	}
	return Parameter
}

//四则类运算
type Operator struct {
	sentence         string   //1+（2*3)文字表达式
	opers            []string //表达式存储
	suffixExpression []string //后缀表达式
}

func NewOperator(sentence string) (*Operator, error) {
	//初始化分配内存
	o := &Operator{sentence, make([]string, 0, len(sentence)), make([]string, 0, len(sentence))}
	o.init()
	err := o.setSuffixExpression() //处理后缀表达式
	return o, err
}

//初始化,字符串切割opers 1 + 2
func (op *Operator) init() {
	value := make([]rune, 0, len(op.sentence)) //定义数组开辟内存
	flag := false
	for _, c := range op.sentence {
		if isBelongToDecimal(string(c)) { //处理数字
			value = append(value, c)
			flag = true
			if c == rune('E') {
				flag = false
			}

		} else if (c == '-' || c == '+') && !flag {
			value = append(value, c)
			flag = true
		} else if isAvailableCode(string(c)) {
			if flag && len(value) > 0 {
				op.opers = append(op.opers, string(value))
				flag = false
			}
			value = value[:0]
			op.opers = append(op.opers, string(c))
		} else if isParmerterCode(string(c)) {
			if flag && len(value) > 0 {
				op.opers = append(op.opers, string(value))
			}
			value = value[:0]
			op.opers = append(op.opers, string(c))
			flag = false
		} else if c == 'I' {
			if flag && len(value) > 0 {
				op.opers = append(op.opers, string(value))
			}
			value = value[:0]
			op.opers = append(op.opers, string(c))
			flag = false
		}
	}
	//处理最后一段
	if flag && len(value) > 0 {
		op.opers = append(op.opers, string(value))
	}
}

//中缀表达式转化为后缀表达式，自动具备了顺序
func (op *Operator) setSuffixExpression() error {
	ops := NewStack()
	var top interface{} //任何类型，数字，运算符
	for i := 0; i < len(op.opers); i++ {
		cur := op.opers[i] //获得当前字符
		if cur == "(" {
			ops.Push(cur)
		} else if cur == ")" {
			for {
				top = ops.Pop() //弹出数据
				if top == "(" {
					break
				}
				op.suffixExpression = append(op.suffixExpression, M2String(top))
				//括号之间的数据截取
			}

		} else if cur == "+" || cur == "-" || cur == "*" || cur == "/" || cur == "^" || cur == "√" {
			for {
				if ops.IsEmpty() || ops.Peak() == "(" || ((cur == "*" || cur == "/") && (ops.Peak() == "-" || ops.Peak() == "+")) || ((cur == "^" || cur == "√") && (ops.Peak() == "+" || ops.Peak() == "-" || ops.Peak() == "*" || ops.Peak() == "/")) {
					ops.Push(cur)
					break
				} else {
					top = ops.Pop()
					op.suffixExpression = append(op.suffixExpression, M2String(top))
				}
			}
		} else {
			op.suffixExpression = append(op.suffixExpression, M2String(cur))
		}
	}
	for {
		if ops.IsEmpty() { //如果栈为空，跳出循环
			break
		}
		top := ops.Pop()
		if top != "(" {
			op.suffixExpression = append(op.suffixExpression, M2String(top))
		} else {
			return errors.New("错误的表达式")
		}
	}
	return nil
}

//
func (op *Operator) Execute(str []string) (value float64, err error) {
	temp := NewStack()
	for i := 0; i < len(op.suffixExpression); i++ {
		st := changeParameter(op.suffixExpression[i], str)
		if val, err := strconv.ParseFloat(strings.TrimSpace(st), 64); err == nil {
			temp.Push(val) //压入数据
		} else {
			exp := op.suffixExpression[i] //取出数据
			if exp == "I" {
				v1 := temp.Pop()       //取得弹出的数据
				temp.Push(M2int64(v1)) //压入数据,计算好的结果
			} else {
				rights := temp.Pop()       //取出数据
				right := M2float64(rights) //数据转化
				lefts := temp.Pop()
				left := M2float64(lefts)
				temp.Push(exeSingleExpression(left, right, exp))
			}
		}
	}
	value = M2float64(temp.Pop())
	return value, err
}
