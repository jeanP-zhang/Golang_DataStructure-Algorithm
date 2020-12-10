package main

import "fmt"

type pos struct {
	x, y int
}

//递归转化为栈

//解决递归走出来
func AiOut(AiData [M][N]int, i int, j int) bool {
	AiData[i][j] = 3
	if i == M-1 && j == N-1 {
		canGoOut = true
		fmt.Println("gameOver")
	} else {
		if j+1 <= N-1 && AiData[i][j+1] < 2 && canGoOut != true {
			AiOut(AiData, i, j+1) //递归一次
		}

		if i+1 <= M-1 && AiData[i+1][j] < 2 && canGoOut != true {
			AiOut(AiData, i+1, j) //递归一次
		}

		if j-1 >= 0 && AiData[i][j-1] < 2 && canGoOut != true {
			AiOut(AiData, i, j-1) //递归一次
		}

		if i-1 >= 0 && AiData[i-1][j] < 2 && canGoOut != true {
			AiOut(AiData, i-1, j) //递归一次
		}
		if canGoOut != true {
			AiData[i][j] = 3 //走不通设置为0
		}
		return canGoOut
	}

}
