package main

import (
	"fmt"
)

//ToDo:main1x
func main() {
	isok := AiOut(AiData, 0, 0)
	if isok {
		fmt.Println("可以走出")
		show(AiData)
	} else {
		fmt.Println("走不出")
	}
}
func main1x() {
	show(Data)
	for {
		var inputStr string
		fmt.Scan(&inputStr)
		run(inputStr)
	}
}
