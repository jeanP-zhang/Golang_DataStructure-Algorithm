package main

import "fmt"

func SelectSortMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	} else {
		max := arr[0]
		for i := 1; i < length; i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}
}
func countSort(arr []int) []int {
	max := SelectSortMax(arr)          //寻找最大值
	sortedArr := make([]int, len(arr)) //排序之后存储
	countArr := make([]int, len(arr))  //统计次数
	for _, v := range arr {
		countArr[v]++
	}
	fmt.Println("第一次统计次数", countArr) //统计次数
	for i := 1; i <= max; i++ {
		countArr[i] += countArr[i-1] //叠加
	}
	fmt.Println("次数叠加", countArr) //统计次数
	for _, v := range arr {
		sortedArr[countArr[v]-1] = v //展开数据
		countArr[v]--                //递减
	}
	return sortedArr
}
func main() {
	arr := []int{1, 2, 3, 4, 4, 3, 2, 1, 2, 5, 5, 3, 4, 3, 2, 1}
	fmt.Println(countSort(arr))
}
