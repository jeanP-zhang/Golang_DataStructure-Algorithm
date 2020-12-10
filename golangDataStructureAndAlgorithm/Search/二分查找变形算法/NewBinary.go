package main

import "fmt"

//找到第一个等于3的
//找到最后一个等于3的
//找到第一个大于等于2的
//
func NewBinarySearch(arr []int ,data int)int{
low:=0
high:=len(arr)-1
index:=-1
for low<=high{
	mid:=low+(high-low)/2
	if arr[mid]>data{
		high=mid-1
	}else if arr[mid]<data{
		low=mid+1
	}else{
		if mid==len(arr)-1||arr[mid+1]!=data{
			index=mid
			return index
		}else{
			low=mid+1//递归继续查找
		}

	}
}
return index
}
func main()  {
	arr:=[]int{1,2,3,4,5,6,6,7,8,8,9,10,11,12}
	for i:=0;i<len(arr) ;i++  {
		fmt.Println("index",arr[i])
	}
}
