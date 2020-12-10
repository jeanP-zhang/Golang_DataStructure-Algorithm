package main

import "fmt"

const IntMax = int(^uint(0) >> 1) //移位操作
type Node struct {
	Value int //数据
	Next  int //下一个索引
}

var NL []Node //链表集合
func ListSort() {
	var i, low, high int
	for i = 2; i < len(NL); i++ {
		low = 0
		high = NL[0].Next
		for NL[high].Value < NL[i].Value { //寻找一个邻居的数据 在NL[max]和NL[i]之间,插入NL[min]
			low = high
			high = NL[high].Next
		}
		NL[low].Next = i
		NL[i].Next = high //插入数据到中间
	}
}
func Arrange() {
	p := NL[0].Next
	for i := 1; i < len(NL); i++ {
		for p < i {
			p = NL[p].Next
		}
		q := NL[p].Next
		if p != i {
			NL[p].Value, NL[i].Value = NL[i].Value, NL[p].Value
			NL[p].Next = NL[i].Next //修改next
			NL[i].Next = p          //地址的插入
		}
		p = q //寻找下一个
	}
	for i := 1; i < len(NL); i++ {
		fmt.Println(NL[i].Value)
	}
}
func InitList(arr []int) {
	var node Node
	node = Node{IntMax, 1} //哨兵
	NL = append(NL, node)
	for i := 1; i <= len(arr); i++ {
		node = Node{arr[i-1], 0} //插入一个数据
		NL = append(NL, node)
	}
}
func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 6, 4, 5, 10}
	InitList(arr) //初始化
	ListSort()    //排序
	fmt.Println(NL)
	fmt.Println("----------------------")
	Arrange()
}
