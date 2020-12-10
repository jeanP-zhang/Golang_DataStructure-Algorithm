package main

import "fmt"

//3 2 9 1 5 7
//3
//21 3 957   双冒泡，小于我的往左，大于我的往右
//1 2 3
//5 7 9

func QuickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	splitData := arr[0] //第一个数字
	low := make([]int, 0, 0)
	high := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	mid = append(mid, splitData) //保存分离的数据
	for i := 1; i < length; i++ {
		if arr[i] < splitData {
			low = append(low, arr[i])
		} else if arr[i] > splitData {
			high = append(high, arr[i])
		} else {
			mid = append(mid, arr[i])
		}

	}
	low, high = QuickSort(low), QuickSort(high)
	ans := append(append(low, mid...), high...)
	return ans
}

//二分查找
func binarySearch(arr []int, data int) int {
	left := 0             //最下面
	right := len(arr) - 1 //最上面
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > data {
			right = mid - 1
		} else if arr[mid] < data {
			left = mid + 1
		} else {
			return arr[mid]
		}
	}
	return -1
}

func main() {
	arr := []int{1, 19, 4, 8, 3, 5, 4, 6, 18, 0}
	fmt.Println(QuickSort(arr)
	fmt.Scan()
}
