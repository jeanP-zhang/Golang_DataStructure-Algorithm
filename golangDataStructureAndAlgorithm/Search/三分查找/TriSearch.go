package main

import "fmt"

//二分查找比三分查找更快
func TriSearch(arr []int, data int) int {
	low := 0
	high := len(arr) - 1 //确定底部与高部
	for low <= high {
		mid1 := low + (high-low)/3
		mid2 := high - (high-low)/3
		midData1 := arr[mid1]
		midData2 := arr[mid2]
		if midData1 == data {
			return mid1
		} else if midData2 == data {
			return mid2
		}
		if midData1 < data {
			low = mid1 + 1
		} else if midData2 > data {
			high = mid2 - 1
		} else {
			high = high + 1
			low = low - 1
		}

	}
	return -1
}

func main() {
	arr := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	fmt.Println(TriSearch(arr, 367))
}
