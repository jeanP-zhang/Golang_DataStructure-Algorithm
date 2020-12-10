package main

import (
	"fmt"
)

//选择排序
func SelectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length; i++ {
		min := i //索引标记
		for j := i; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i] //数据交换
		}
	}
	return arr
}

//适用于排序数量有限的情况
func bucketSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	num := length
	max := SelectSortMax(arr)     //极大值
	index := 0                    //索引
	buckets := make([][]int, num) //创造二维数组
	for i := 0; i < length; i++ {
		index = arr[i] * (num - 1) / max                //木桶的自动分配算法
		buckets[index] = append(buckets[index], arr[i]) //木桶计数+1
	}
	fmt.Println(buckets)
	tmpPose := 0 //木桶排序
	for i := 0; i < num; i++ {

		bucketsLen := len(buckets[i]) //q求某一段长度
		if bucketsLen > 0 {
			buckets[i] = SelectSort(buckets[i]) //木桶内部数据排序
			copy(arr[tmpPose:], buckets[i])     //拷贝数据
			tmpPose += bucketsLen               //定位
		}
	}
	return arr
}

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
func main() {
	arr := []int{1, 2, 3, 3, 3, 3, 4, 4, 5, 5, 6, 6, 6, 6, 6, 7}
	bucketSort(arr)
	fmt.Println(arr)
}
