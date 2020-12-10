package main

import "fmt"

//arr需要排序的数组，lastArr排序排列好的数组，level代表级别，tread代表线程数量
/*func QuickSortAddrThread(arr []int, lastArr chan int, level int, thread int) {
	level = level * 2 //每加深一个级别就多一个线程
	if len(arr) == 0 {
		close(lastArr)
		return
	} else if len(arr) == 1 {
		lastArr <- arr[0] //为一个数据放入管道
		close(lastArr)
		return
	} else {
		less := make([]int, 0)        //比我小的数据
		greater := make([]int, 0)     //比我大的数据
		midder := make([]int, 0)      //与我相等的数据
		left := arr[0]                //取得第一个数据
		midder = append(midder, left) //中间存放相等的数据
		for i := 1; i < len(arr); i++ {
			if arr[i] < left {
				less = append(less, arr[i]) //处理小于
			} else if arr[i] > left {
				greater = append(greater, arr[i]) //处理大于
			} else {
				midder = append(midder, arr[i]) //处理等于
			}
		}
		leftCh, rightCh := make(chan int, len(less)), make(chan int, len(greater))
		if level <= thread { //限制线程数量，如果线程超过执行数量，顺序调用，否则并发调用
			fmt.Println("111111")
			go QuickSortThread(less, leftCh, level, thread)
			go QuickSortThread(greater, rightCh, level, thread)
		} else {
			fmt.Println("22222")
			QuickSortThread(less, leftCh, level, thread)
			QuickSortThread(greater, rightCh, level, thread)
		}
		for i := range leftCh {
			lastArr <- i
		}
		for _, v := range midder {
			lastArr <- v
		}
		for i := range rightCh {
			lastArr <- i
		}
		close(lastArr)
		return
	}

}*/
func QuickSort(arr []int, addr []int) []int {
	length := len(addr)
	if length <= 1 {
		return arr
	} else {
		splitData := arr[addr[0]]
		low := make([]int, 0)
		mid := make([]int, 0)
		high := make([]int, 0)
		mid = append(mid, splitData)
		for i := 1; i < length; i++ {
			if arr[i] < splitData {
				low = append(low, addr[i])
			} else if arr[i] > splitData {
				high = append(high, addr[i])
			} else {
				mid = append(mid, addr[i])
			}
		}
		low, high = QuickSort(arr, low), QuickSort(arr, high)
		myArr := append(append(low, mid...), high...)
		return myArr
	}

}
func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 6, 4, 5, 10}
	arrAddr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//lastArr := make(chan int) //管道，将排好得数据压入管道
	QuickSort(arr, arrAddr)
	for i := 0; i < len(arr); i++ { //显示管道的每一个数据
		fmt.Println(arr[i], arrAddr[i])
	}
}
