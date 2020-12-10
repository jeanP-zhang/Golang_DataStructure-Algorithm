package main

import "fmt"

func swap(arr []int, i, j int) {

	arr[i], arr[j] = arr[j], arr[i]

}
func QuickSort(arr []int) {

}
func findKLargest(arr []int, k int) int {
	return findKLargestGo(arr, 0, len(arr)-1, k)
}
func findKLargestGo(arr []int, left, right, k int) int {
	if left >= right {
		return arr[left]
	}
	query := partition(arr, left, right) //切割
	if query+1 == k {
		return arr[query] //第K大的数
	}
	if k < query+1 {
		return findKLargestGo(arr, left, query-1, k) //递归一直操作到区间为1
	}
	return findKLargestGo(arr, left, query+1, k)
}
func partition(arr []int, left, right int) int {
	pivot := right
	i := left
	for j := left; j < pivot; j++ {
		if arr[j] > arr[pivot] {
			swap(arr, i, j)
			i++
		}
	}
	swap(arr, i, pivot)
	return i
}
func main() {
	arr := []int{1, 3, 4, 51, 2, 5, 1, 5, 3, 6, 7, 324, 8, 2, 6357, 32, 2, 42, 342}
	fmt.Println(findKLargest(arr, 3))
	//提取QQ中间最大的100个，最小的100个
}
