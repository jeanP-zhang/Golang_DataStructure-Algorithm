package main

import "fmt"

func main() {
	a:=[]int{32,1,5,6,3,7,8,9,3,2,1,11,10,21}
fmt.Println(QuickSort(a))
}
func QuickSort(arr []int)[]int{
	length:=len(arr)//数组长度
if length<=1{
	return arr
}else{
	splitData:=arr[0]//以第一个为基准
	low,high,mid:=make([]int,0,0),make([]int,0,0),make([]int,0,0)//存储比我小、大、相等的
	mid=append(mid,splitData)
	for i:=1;i<length ;i++  {
		if arr[i]<splitData{
			low=append(low,arr[i])
		}else if arr[i]>splitData{
			high=append(high,arr[i])
		}else{
			mid=append(mid,arr[i])
		}
	}
	low,high=QuickSort(low),QuickSort(high)//切割递归处理
	myArray:=append(append(low,mid...),high...)
	return myArray
}
}
