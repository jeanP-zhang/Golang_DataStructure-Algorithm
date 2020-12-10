package main

import "fmt"

//动态规划思路
const N = 5
const W = 6

var weight = []int{3, 1, 1, 1, 1}
var val = []int{9, 3, 2, 6, 7}
var dp [N + 1][W + 1]int
var record [N][W + 1]int

func solve2() int {
	length := len(val)
	for i := 0; i < length; i++ {
		for j := weight[i]; j <= W; j++ {
			dp[i+1][j] = max(dp[i][j], dp[i][j-weight[i]]+val[i])
		}
	}
	return dp[N][W]
}
func init() {
	for i := 0; i < N; i++ {
		for j := 0; j <= W; j++ {
			record[i][j] = -1
		}
	}
}
func max(a, b int) (c int) {
	if a > b {
		c = a
		return
	}
	c = b
	return
}

//i是以及有的重量，total总量
func solve(i, total int) int {
	fmt.sc
	result := 0 //结果
	if i >= N {
		return result
	}
	if record[i][total] != -1 {
		return record[i][total] //如果数据已经记录，直接返回
	}
	if weight[i] > total {
		record[i][total] = solve(i+1, total) //当前物品大于总量，跳出，计算下一个
	} else {
		//递归求最大值，退出一个再加一个
		result = max(solve(i+1, total), solve(i+1, total-weight[i])+val[i])
	}
	record[i][total] = result
	return record[i][total]
}
func main() {
	fmt.Println(solve(0, W))
	fmt.Println(solve2())
	fmt.Println(dp)
}
