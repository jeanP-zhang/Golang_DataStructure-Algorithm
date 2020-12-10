package main

import (
	"code.qiangqiang.com/studygo/golang-数据结构与算法/链表/SingleLink"
	"fmt"
)

func main2() {
	list := SingleLink.NewSingleLinkList()
	node1 := SingleLink.NewSingleLinkNode(1)
	//	fmt.Println(&node1)
	node2 := SingleLink.NewSingleLinkNode(2)
	//fmt.Println(&node2)
	node3 := SingleLink.NewSingleLinkNode(3)
	//fmt.Println(&node3)
	//node4 := SingleLink.NewSingleLinkNode(4)
	fmt.Println("------------------------------")
	//fmt.Println(list.Head)
	list.InsertNodeFront(node1)
	//fmt.Println(list, &list.Head.Next)
	list.InsertNodeFront(node2)
	//fmt.Println(list, &list.Head.Next)
	list.InsertNodeBack(node3)
	//fmt.Println(list, &list.Head.Next)
	list.InsertNodeValueFront(4, node1)
	list.InsertNodeValueFront(4, node2)
	list.InsertNodeValueFront(4, node3)
	fmt.Println(list.String())
}

//func main1() {
//	list := SingleLink.NewSingleLinkList()
//	path := ""
//	sqFile, _ := os.Open(path)
//	br := bufio.NewReader(sqFile)
//}
/*链表的中间节点：双指针A、B，A一次走两步，B一次走一步，

 */
