package main

import (
	"fmt"
	"math/rand"
	"time"
)

//洗牌算法
func isOrder(list []int)bool  {
	for i:=1;i<len(list);i++{
		if list [i-1]>list[i]{
			return false
		}
	}
	return true
}
//洗牌算法
//理解为洗牌
func randList(list []int)  {
	data:=make([]int,len(list))
	copy(data ,list)
	rand.Seed(time.Now().UnixNano())//定义随机数种子
	index:=rand.Perm(len(list))//随机选择一个切片
	for i,k:=range index{
		list[i]=data[k]
	}
}
func main()  {
	list:=[]int{1,9,2,8,3,7,4,5,22,30,25}
fmt.Println(list)
	count :=0
for {
	if isOrder(list){
		fmt.Println("排序完成",list)
		break
	}else{
		randList(list)
		count++
	}
}
fmt.Println(count)
}
