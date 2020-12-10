package main

import (
	"code.qiangqiang.com/studygo/golang-数据结构与算法/threadsafe/Queue"
	"fmt"
)

func main() {
	h := Queue.NewHeap() //最小堆
	h.Insert(Queue.Int(8))
	h.Insert(Queue.Int(9))
	h.Insert(Queue.Int(7))
	h.Insert(Queue.Int(5))
	h.Insert(Queue.Int(4))
	h.Insert(Queue.Int(6))
	fmt.Println(h.Extract().(Queue.Int))
}
