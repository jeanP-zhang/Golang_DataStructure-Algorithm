package main

import (
	"code.qiangqiang.com/studygo/golang-数据结构与算法/双链表/doubleLinkList"
	"container/list"
	"fmt"
)

t (
	"code.qiangqiang.com/studygo/golang-数据结构与算法/双链表/doubleLinkList"
	"fmt"
)



import (
	"code.qiangqiang.com/studygo/golang-数据结构与算法/双链表/doubleLinkList"
	"fmt"
	"container/list"
)

func main() {
	lists := doubleLinkList.NewDoubleLinkList()
	node1 := doubleLinkList.NewDoubleLinkNode(1)
	node2 := doubleLinkList.NewDoubleLinkNode(22)
	node3 := doubleLinkList.NewDoubleLinkNode(33)
	node4 := doubleLinkList.NewDoubleLinkNode(12)
	lists = lists.InsertHead(node1)
	//	fmt.Println(lists)
	lists = lists.InsertHead(node2)
	//fmt.Println(lists)
	lists = lists.InsertBack(node3)
	//fmt.Println(lists)
	lists = lists.InsertBack(node4)
	//	fmt.Println(lists.String())
	lists.InsertBack(node1)
	lists.InsertValueBack(node2, node1)
	//lists.InsertValueHead(node3, node2)
	fmt.Println(lists)
newlist:=list.New()
}
