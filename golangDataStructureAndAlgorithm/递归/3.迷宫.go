package main

import (
	"fmt"

)

const (
	M = 10
	N = 10
)

var Data = [M][N]int{
	{1, 0, 2, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 2, 0, 0, 2, 0, 0, 0},
	{0, 0, 0, 0, 2, 0, 0, 2, 2, 0},
	{0, 0, 0, 0, 2, 0, 2, 0, 0, 0},
	{2, 2, 2, 0, 0, 2, 0, 2, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 2, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 2, 2, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 2, 0},
}
var AiData = [M][N]int{
	{1, 0, 2, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 2, 0, 0, 2, 0, 0, 0},
	{0, 0, 0, 0, 2, 0, 0, 2, 2, 0},
	{0, 0, 0, 0, 2, 0, 2, 0, 0, 0},
	{2, 2, 2, 0, 0, 2, 0, 2, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 2, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 2, 2, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 2, 0},
}
var iPos, jPos = 0, 0     //定义初始位置
var canGoOut = false //走不出迷宫
func show(arr [M][N]int)  {
	for i:=0;i<M ;i++  {
		for j:=0;j<N ;j++  {
			fmt.Printf("%4d",arr[i][j])
		}
		fmt.Println(" ")
	}
}
func Find(dataX [M][N]int,i,j int)bool  {
dataX[i][j]=3
	if i==M-1&&j==N-1{
		canGoOut=true
		Data=dataX
		fmt.Println("迷宫可以走出来")
	}else {
	if j+1<=N-1&&dataX[i][j+1]<2&&canGoOut!=true{
		Find(dataX,i,j+1)
	}
		if j-1>=0&&dataX[i][j-1]<2&&canGoOut!=true{
			Find(dataX,i,j-1)
		}
		if i+1<=M-1&&dataX[i+1][j]<2&&canGoOut!=true{
			Find(dataX,i+1,j)
		}
		if i-1>=0&&dataX[i-1][j]<2&&canGoOut!=true{
			Find(dataX,i-1,j)
		}
	}
	return canGoOut
}
func main()  {
	isok:=Find(Data,0,0)
if isok{
	fmt.Println("可以走出")
	show(Data)
}
}