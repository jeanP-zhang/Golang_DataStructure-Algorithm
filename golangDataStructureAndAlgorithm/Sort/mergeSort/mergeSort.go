package main

import "fmt"

func merge(leftArr,rightArr []int)[]int  {
	leftIndex:=0//左边索引
	rightIndex:=0//右边索引
	lastArr:=make([]int,0)
for leftIndex<len(leftArr)&&rightIndex<len(rightArr){
	if leftArr[leftIndex]<rightArr[rightIndex]{
lastArr=append(lastArr,leftArr[leftIndex])
leftIndex++
	}else if leftArr[leftIndex]>rightArr[rightIndex]{
		lastArr=append(lastArr,rightArr[rightIndex])
	rightIndex++
	}
	lastArr=append(lastArr,leftArr[leftIndex])
	lastArr=append(lastArr,rightArr[rightIndex])
	leftIndex++
	rightIndex++
}
for leftIndex<len(leftArr){
	lastArr=append(lastArr,leftArr[leftIndex])
	leftIndex++
}
for rightIndex<len(rightArr){
	lastArr=append(lastArr,rightArr[rightIndex])
	rightIndex++
}
return lastArr
}

func mergeSort(arr []int)[]int  {
	length:=len(arr)
	if length<=1{
		return  arr//小于10改用插入排序
	}
	mid:=length/2
	leftArr:=mergeSort(arr[:mid])
	rightArr:=mergeSort(arr[mid+1:])
	return merge(leftArr,rightArr)
}

func main() {
	a:=[]int{32,1,5,6,3,7,8,9,3,2,1,11,10,21}
	fmt.Println(mergeSort(a))
}