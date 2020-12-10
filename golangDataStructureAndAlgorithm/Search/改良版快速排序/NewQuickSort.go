package main

import (
	"fmt"
	"math/rand"
)

//快速排序的核心代码
func NewQuickSort(arr []int) []int {
	sortForMerge(arr, 0, len(arr))
	return arr
}
func swap(arr []int, i, j int) {

	arr[i], arr[j] = arr[j], arr[i]

}
func sortForMerge(arr []int, left, right int) {
	for i := left; i < right; i++ {
		temp := arr[i]
		var j int
		for j = i; j > left && arr[j-1] > temp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp //插入
	}
}

//快速排序递归
func QuickSort(arr []int, left, right int) {
	if right-left < 3 { //数组剩下三个直接插入排序
		sortForMerge(arr, left, right)
	} else {
		//随机找一个数字
		swap(arr, left, rand.Int()%(right-left+1)+left)
		vData := arr[left] //坐标数据，比我小左边，比我大右边
		lt := left         //arr[left+1,lt]<vData
		gt := right + 1    //arr[gt.... right]>vData
		i := left + 1      //arr[lt+1,...i]==vData
		for i < gt {
			if arr[i] < vData {
				swap(arr, i, lt+1)
				lt++
				i++
			} else if arr[i] > vData {
				swap(arr, i, gt-1) //移动到大于的地方
				gt--
			} else {
				i++
			}
		}
		swap(arr, left, lt)         //交换头部位置
		QuickSort(arr, left, lt-1)  //递归处理小于那一段
		QuickSort(arr, gt+1, right) //递归处理大于那一段
	}
}
func main() {
	arr := []int{1, 4, 5, 6, 2, 6, 73, 7, 9, 3, 6, 2, 56, 48, 7, 3, 2, 5}
	fmt.Println("未排序", arr)
	fmt.Println("已排序", NewQuickSort(arr))
}
