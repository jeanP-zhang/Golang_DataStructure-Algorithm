package main

import "fmt"

//单链表的节点
type Node struct {
	num  int
	next *Node
}

//环链表
var head, tail *Node

func AddNode(n *Node) { //没有节点的情况下
	if tail == nil {
		head = n
		n.next = head
		tail = n
	} else {
		tail.next = n
		n.next = head
		tail = n
	}
}
func ShowList(head *Node) {
	if head == nil {
		return
	} else {
		//循环约瑟夫环
		for head.next != nil && head != tail {
			fmt.Println(head.num)
			head = head.next
		}
		fmt.Println(head.num)
	}
}

//从第k个循环起，循环起第num个，留下最后一个
func joser(k, num int) {
	count := 1
	for i := 0; i <= k-1; i++ {
		head = head.next
		tail = tail.next //循坏到起点
	}
	for {
		count++ //开始记录次数
		head = head.next
		tail = tail.next
		if count == num {
			fmt.Println(head.num, "出局")
			tail.next = head.next
			head = head.next
			count = 1 //清零
		}
		if head == tail { //相等意味着只剩最后一个
			fmt.Println(head.num, "最后一个")
			break
		}
	}
}
func main() {
	for i := 0; i < 10; i++ {
		n := &Node{i, nil}
		AddNode(n)
	}
	joser(0, 3)
	ShowList(head)
}
