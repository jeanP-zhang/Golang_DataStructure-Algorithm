package main

import "fmt"

//“wsad”上下左右
func run(direct string) {
	switch direct {
	case "w":
		if iPos-1 >= 0 && Data[iPos-1][jPos] < 2 {
			//简单控制
			Data[iPos][jPos], Data[iPos-1][jPos] = Data[iPos-1][jPos], Data[iPos][jPos]
			iPos -= 1
		}
	case "s":
		if iPos+1 <= M-1 && Data[iPos+1][jPos] < 2 {
			Data[iPos][jPos], Data[iPos+1][jPos] = Data[iPos+1][jPos], Data[iPos][jPos]
			iPos += 1
		}
	case "a":
		if jPos-1 >= 0 && Data[iPos][jPos-1] < 2 {
			Data[iPos][jPos], Data[iPos][jPos-1] = Data[iPos][jPos-1], Data[iPos][jPos]
			jPos -= 1
		}
	case "d":
		if jPos+1 <= N-1 && Data[iPos][jPos+1] < 2 {
			Data[iPos][jPos], Data[iPos][jPos+1] = Data[iPos][jPos+1], Data[iPos][jPos]
			jPos += 1
		}
	default:
		fmt.Println("输入错误")

	}
	show(Data)
}
func AiMoveOut() {
	AiData[0][0] = 1
	for iPos != M-1 && jPos != N-1 {
		if iPos-1 >= 0 && AiData[iPos-1][jPos] == 3 {
			AiData[iPos-1][jPos] = 0
			run("w")
		}
		if jPos-1 >= 0 && AiData[iPos][jPos-1] == 3 {
			AiData[iPos-1][jPos] = 0
			run("a")
		}
		if iPos+1 <= M && AiData[iPos+1][jPos] == 3 {
			AiData[iPos+1][jPos] = 0
			run("s")
		}
		if jPos+1 <= N-1 && AiData[iPos][jPos+1] == 3 {
			AiData[iPos][jPos+1] = 0
			run("d")
		}
	}
}
