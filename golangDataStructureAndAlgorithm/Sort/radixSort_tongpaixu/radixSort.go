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
func RadixSort(arr []int) []int {
	max := SelectSortMax(arr)              //寻找数组的极大值
	for bit := 1; max/bit > 0; bit *= 10 { //按照数量级分段
		arr = BitSort(arr, bit)
	}
	return arr
}
func BitSort(arr []int, bit int) []int {
	length := len(arr)
	bitCounts := make([]int, 10)
	for i := 0; i < length; i++ {
		num := (arr[i] / bit) % 10 //分层处理
		bitCounts[num]++           //统计余数相等个数
	}
	for i := 1; i < 10; i++ {
		bitCounts[i] += bitCounts[i-1] //叠加，计算位置
	}
	tmp := make([]int, 10)
	for i := length - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10
		tmp[bitCounts[num]-1] = arr[i]
		bitCounts[num]-- //计算排序的位置
	}
	for i := 0; i < length; i++ {
		arr[i] = tmp[i] //保存数组
	}
	return arr
}
func main() {
	arr := []int{11, 91, 222, 878, 348, 7123, 4213, 6232, 5123, 1011}
	fmt.Println("结果是：", RadixSort(arr))
}
