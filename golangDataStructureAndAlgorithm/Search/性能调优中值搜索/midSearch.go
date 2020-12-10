package main

import "fmt"

//斐波那契搜索 mid=left+(right-left)*(data-left)/(right-left)
//二分查找
func binarySearch(arr []int, data int) int {
	left := 0             //最下面
	right := len(arr) - 1 //最上面
	for left < right {
		leftV := float64(data - arr[left])      //大段
		allV := float64(arr[right] - arr[left]) //整段
		diff := float64(right - left)
		mid := int(float64(right) + leftV/allV*diff)
		if mid < 0 || mid > len(arr) {
			return -1
		}
		//		mid := left + (right-left)/2
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
	arr := make([]int, 1000, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
		fmt.Println(arr[i])
	}

}
