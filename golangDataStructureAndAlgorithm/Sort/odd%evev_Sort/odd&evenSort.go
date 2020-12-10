package main

import "fmt"

func OddEven(arr []int)[]int{
	//tmp:=0
	isSorted:=false//奇数位，偶数位都不需要交换则完成排序
	for !isSorted{
		isSorted=true
	for i:=1;i<len(arr)-1;i=i+2{//奇数位
if arr[i]>arr[i+1]{
	arr[i],arr[i+1]=arr[i+1],arr[i]
	isSorted=false
}
	}
	for i:=0;i<len(arr)-1;i=i+2{//偶数位
		if arr[i]>arr[i+1]{
			arr[i],arr[i+1]=arr[i+1],arr[i]
			isSorted=false
		}
	}
	}
	return arr
}
func main() {
	a:=[]int{32,1,5,6,3,7,8,9,3,2,1,11,10,21}
	fmt.Println(OddEven(a))
}
