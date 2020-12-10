package main

import "fmt"

//按照斐波那契进行数据的切割，和二分查找类似

func fabSearch(arr []int, val int) int {
	length := len(arr)          //数组长度
	fabArr := makeFabArray(arr) //填充长度
	fmt.Println(fabArr)
	fillLength := fabArr[len(fabArr)-1] //填充的数组
	fillArr := make([]int, fillLength)
	for i, v := range arr {
		fillArr[i] = v
	}
	lastData := arr[length-1]
	for i := length; i < fillLength; i++ {
		fillArr[i] = lastData
	}
	left, mid, right := 0, 0, length //类似二分查找
	kIndex := len(fabArr) - 1
	for left < right {
		mid = left + fabArr[kIndex-1] - 1 //斐波那契切割
		if val < fillArr[mid] {
			right = mid - 1
			kIndex--
		} else if val > fillArr[mid] {
			left = mid + 1
			kIndex -= 2
		} else {
			if mid > right {
				return right
			}
			return mid
		}
	}

	return -1
}
func makeFabArray(arr []int) []int {
	length := len(arr)
	fibLen := 2
	first, second, third := 1, 2, 3
	for third < length { //找出最接近的斐波那契
		first, second, third = second, third, first+second
		fibLen++
	}
	fb := make([]int, fibLen) //开辟数组
	fb[0] = 1
	fb[1] = 1
	for i := 2; i < fibLen; i++ {
		fb[i] = fb[i-1] + fb[i-2]
	}
	return fb
}
func main() {
	arr := make([]int, 1000, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	fmt.Println(fabSearch(arr, 135))
}
