package Calc

//主函数
func Calc(input string) int64 {
	lexer := NewLexer(input)
	parser := NewParser(lexer)
	exp := parser.ParseExpression(LOWEST)
	return Eval(exp)
}
func Eval(exp Expression) int64 {
	switch node := exp.(type) {
	case *IntergerLiteralExpression:
		return node.Value
	case *PreFixExpression:
		rightValue := Eval(node.Right)
		return evalPrefixExpression(node.Operator, rightValue)
	case *InfixExpression:
		leftv := Eval(node.Left)
		rightv := Eval(node.Right)
		return evalInfixExpression(leftv, node.operator, rightv) //计算
	}

	return 0
}

//1+-1
func evalPrefixExpression(operator string, right int64) int64 {
	if operator == "+" {
		return right
	} else if operator == "-" {
		return -1 * right
	} else if operator == "!" {
		if right == 0 {
			return 1
		}
		return 0
	} else if operator == ">" {
		return 1
	} else {
		return 0
	}
}
func evalInfixExpression(left int64, operator string, right int64) int64 {
	switch operator {
	case ">":
		if left > right {
			return 1
		} else {
			return 0
		}
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		if right == 0 {
			return 0
		} else {
			return left / right
		}
	case "%":
		return left % right
	default:
		return 0
	}
}

//判断是否为数字字符
func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
