package SingleLink

import (
	"fmt"
	"strings"
)

//单链表适用于增加删除比较多的数据场合
type SingleLink interface {
	//增删查改
	GetFirstNode() *SingleLinkNode            //抓取头部节点
	InsertNodeFront(node *SingleLinkNode)     //头部插入
	InsertNodeBack(node *SingleLinkNode)      //尾部插入
	GetNodeAtIndex(index int) *SingleLinkNode //根据索引抓取指定位置数据
	InsertNodeValueFront(dest interface{}, node *SingleLinkNode) bool
	InsertNodeValueBack(dest interface{}, node *SingleLinkNode) bool
	DeleteNode(dest *SingleLinkNode) bool //删除节点
	DeleteIndex(index int) bool           //删除指定位置的节点
	String() string                       //返回链表字符串
}

//链表的结构
type SingleLinkList struct {
	Head   *SingleLinkNode
	Length int
}

//创建链表
func NewSingleLinkList() *SingleLinkList {
	head := NewSingleLinkNode(nil)
	return &SingleLinkList{head, 0}
}
func (list *SingleLinkList) GetFirstNode() *SingleLinkNode {
	return list.Head.Next
}
func (list *SingleLinkList) InsertNodeFront(node *SingleLinkNode) {
	if list.Head.Value == nil {
		list.Head.Next = node
		node.Next = nil
		list.Length++
	} else {
		bank := list.Head
		node.Next = bank.Next
		bank.Next = node
		list.Length++
	}
}
func (list *SingleLinkList) InsertNodeBack(node *SingleLinkNode) {
	if list.Head.Value == nil {
		list.Head.Next = node
		node.Next = nil
		list.Length++
	} else {
		bank := list.Head
		for bank.Next != nil {
			bank = bank.Next
		}
		bank.Next = node
		list.Length++
	}
}
func (list *SingleLinkList) GetNodeAtIndex(index int) *SingleLinkNode {
	if index > list.Length-1 || index < 0 {
		return nil
	} else {
		pHead := list.Head
		for index > -1 {
			pHead = pHead.Next
			index--
		}
		return pHead
	}

}
func (list *SingleLinkList) DeleteNode(dest *SingleLinkNode) bool {
	isDelete := false
	pNext := list.Head
	if dest == nil {
		return false
	} else {
		for pNext.Next != nil && pNext.Value != dest {
			pNext = pNext.Next
		}
		if pNext.Next == dest {
			pNext.Next = pNext.Next.Next
			list.Length--
		}
	}
	return isDelete
}
func (list *SingleLinkList) DeleteIndex(index int) bool {
	if index > list.Length-1 || index < 0 {
		return false
	} else {
		pHead := list.Head
		for index > 0 {
			pHead = pHead.Next
			index--
		}
		pHead.Next = pHead.Next.Next
		list.Length--
		return true
	}

}
func (list *SingleLinkList) String() string {
	var listString string
	p := list.Head
	i := 0
	if p.Next != nil {
		listString += fmt.Sprintf("%v-->", p.Next.Value)
		p = p.Next //循环
		i++

	}
	listString += fmt.Sprintf("nil")
	return listString
}
func (list *SingleLinkList) InsertNodeValueFront(dest interface{}, node *SingleLinkNode) bool {
	pHead := list.Head
	isFind := false
	for pHead.Next != nil {
		if pHead.Value == dest {
			isFind = true
			break
		}
		pHead = pHead.Next
	}
	if isFind {
		node.Next = pHead.Next
		pHead.Next = node
		list.Length++
		return true
	} else {
		return false
	}
}
func (list *SingleLinkList) InsertNodeValueBack(dest, node *SingleLinkNode) bool {
	pHead := list.Head
	isFind := false
	for pHead.Next != nil {
		if pHead.Next.Value == dest {
			isFind = true
			break
		}
		pHead = pHead.Next
	}
	if isFind {
		node.Next = pHead.Next
		pHead.Next = node
		list.Length++
		return true
	} else {
		return false
	}

}
func (list *SingleLinkList) FindString(data string) {
	phead := list.Head
	for phead.Next != nil {
		if strings.Contains(phead.Value.(string), data) {
			fmt.Println(phead.Value)
		}
		phead = phead.Next
	}
}
func (list *SingleLinkList) GetMid() (*SingleLinkNode, int) {
	if list.Head.Next == nil {
		return nil, -1
	}
	pHead1, pHead2 := list.Head, list.Head
	i := 0
	for pHead2 != nil && pHead2.Next != nil {
		pHead1 = pHead1.Next
		pHead2 = pHead2.Next.Next
		i++
	}
	return pHead1, i //中间节点
}

//链表反转
func (list *SingleLinkList) ReverseList() {
	if list.Head == nil || list.Head.Next == nil {
		return //链表为空或者只有一个节点
	} else {
		var pre *SingleLinkNode  //前面节点
		var cur = list.Head.Next //当前节点

		for cur != nil {
			curNext := cur.Next //后面节点
			cur.Next = pre
			pre = cur
			cur = curNext
		}
		list.Head.Next = pre
	}
}
