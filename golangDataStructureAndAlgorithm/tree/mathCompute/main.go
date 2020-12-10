package main

import "fmt"

//1+(2-1)*3+(1+2)*(1+（1+（2-1）)）
func main() {
	op, err := NewOperator("I+a")
	if err != nil {
		fmt.Println(err)
	}
	//for i := 0; i < len(op.opers); i++ {
	//	fmt.Printf("%T,%v\n", Stoi64(op.opers[i]), Stoi64(op.opers[i]))
	//}
	fmt.Println("------------------------------------------------------------------")
	//fmt.Println(op.suffixExpression)
	value, err := op.Execute([]string{"I", "12", "a", "18"})
	fmt.Println(value)
}
