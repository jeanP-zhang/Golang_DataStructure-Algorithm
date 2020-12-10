package main

import (
	"fmt"
	"sync"
)

type Deque struct {
	array       []interface{}
	left, right int
	fixSize     int
	lock        *sync.RWMutex
}

func NewQueue(cap int) *Deque {

	if cap <= 0 {
		fmt.Println("队列容量必须大于0")
		return nil
	}
	deq := new(Deque)
	deq.array = make([]interface{}, cap)
	deq.fixSize = cap
	deq.lock = new(sync.RWMutex)
	return deq
}
func (deq *Deque) AddLeft(data interface{}) {
	if deq.right == deq.left && deq.left != 0 {
		panic("overflow")
	}
	deq.array[deq.left] = data
	deq.left = deq.left - 1
	if deq.left == -1 {
		deq.left = deq.fixSize - 1 //循环双端队列
	}
}

func (deq *Deque) AddRight(data interface{}) {
	if deq.right == deq.left && deq.right != 0 {
		panic("overflow")
	}
	deq.right = deq.right + 1
	deq.array[deq.right] = data
	//循环
}
func (deq *Deque) DelLeft() interface{} {
	if deq.fixSize == deq.left {
		panic("overflow")
	}
	deq.left = deq.left + 1
	if deq.left == deq.fixSize {
		deq.left = 0
	}
	data := deq.array[deq.left]
	return data

}
func (deq *Deque) DelRight() interface{} {
	if deq.right == deq.left {
		panic("overflow")
	}

	deq.right = deq.right - 1 //循环
	if deq.right == -1 {
		deq.right = deq.fixSize - 1
	}
	data := deq.array[deq.right]
	return data
}
func main() {
	deq := NewQueue(10)
	deq.AddLeft(1)
	fmt.Println(deq.left, deq.right, deq.array)
	deq.AddLeft(2)
	fmt.Println(deq.left, deq.right, deq.array)
	deq.AddLeft(3)
	fmt.Println(deq.left, deq.right, deq.array)
	deq.AddRight(4)
	fmt.Println(deq.left, deq.right, deq.array)
	deq.DelRight()

	fmt.Println(deq.left, deq.right, deq.array)
	deq.DelLeft()
	fmt.Println(deq.left, deq.right, deq.array)
}
