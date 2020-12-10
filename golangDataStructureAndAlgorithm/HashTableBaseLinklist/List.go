package main
import "fmt"

type DoubleLinkNode struct {
	Value interface{}
	Prev  *DoubleLinkNode //上一个节点
	Next  *DoubleLinkNode //下一个节点
}

func NewDoubleLinkNode(value interface{}) *DoubleLinkNode {
	return &DoubleLinkNode{value, nil, nil}
}
func (node *DoubleLinkNode) Values() interface{} {
	return node.Value
}

//返回上一个节点
func (node *DoubleLinkNode) PrevNode() *DoubleLinkNode {
	return node.Prev
}

//返回下一个节点
func (node *DoubleLinkNode) NextNode() *DoubleLinkNode {
	return node.Next
}
type DoubleLinkList struct {
	Head   *DoubleLinkNode
	Length int
}

func NewDoubleLinkList() *DoubleLinkList {
	head := NewDoubleLinkNode(nil)
	return &DoubleLinkList{head, 0}
}

//返回双向链表长度
func (dList *DoubleLinkList) GerLength() int {
	return dList.Length
}
func (dList *DoubleLinkList) GetFirstNode() *DoubleLinkNode {
	return dList.Head.Next
}
func (dList *DoubleLinkList) InsertHead(node *DoubleLinkNode) *DoubleLinkList {
	pHead := dList.Head
	if pHead.Next == nil {
		node.Next = nil
		pHead.Next = node
		node.Prev = pHead
		dList.Length++

	} else {
		node.Next = pHead.Next //下一个节点
		pHead.Next.Prev = node //标记上一个节点
		pHead.Next = node      //标记头部节点
		node.Prev = pHead
		dList.Length++

	}
	return dList
}
func (dList *DoubleLinkList) InsertBack(node *DoubleLinkNode) *DoubleLinkList {
	pHead := dList.Head
	if pHead.Next == nil {
		node.Next = nil
		pHead.Next = node
		node.Prev = pHead
		dList.Length++

	} else {
		for pHead.Next != nil {
			pHead = pHead.Next
		}
		pHead.Next = node
		node.Prev = pHead
		dList.Length++

	}
	return dList
}
func (dList *DoubleLinkList) String() string {
	var listString1 string
	var listString2 string
	p := dList.Head
	i := 0
	if p.Next != nil {
		listString1 += fmt.Sprintf("%v-->", p.Next.Value)
		p = p.Next //循环
		i++
	}
	listString1 += fmt.Sprintf("nil")
	listString1 += "\n"
	for p != dList.Head {
		listString2 += fmt.Sprintf("%v<--", p.Prev.Value)
		p = p.Prev
	}
	listString2 += fmt.Sprintf("nil")
	return listString1 + listString2 + "\n"
}
func (dList *DoubleLinkList) InsertValueHead(dest *DoubleLinkNode, node *DoubleLinkNode) bool {
	pHead := dList.Head
	for pHead.Next != nil && pHead.Next != dest {
		pHead = pHead.Next
	}
	if pHead.Next == dest {
		if pHead.Next != nil {
			pHead.Next.Prev = node //500100的赋值
		}
		node.Next = pHead.Next //300300
		node.Prev = pHead
		pHead.Next = node //400500
		dList.Head = pHead
		dList.Length++
		return true
	} else {
		return false
	}
}
func (dList *DoubleLinkList) InsertValueBack(dest *DoubleLinkNode, node *DoubleLinkNode) bool {
	pHead := dList.Head
	for pHead.Next != nil && pHead.Next != dest {
		pHead = pHead.Next
	}
	if pHead.Next == dest {
		if pHead.Next.Next != nil {
			pHead.Next.Next.Prev = node //500100的赋值
		}
		node.Next = pHead.Next.Next //300300
		pHead.Next.Next = node      //500100
		node.Prev = pHead.Next      //300200
		dList.Head = pHead
		dList.Length++
		return true
	} else {
		return false
	}
}
func (dList *DoubleLinkList) GetNodeAtIndex(index int) *DoubleLinkNode {
	if index > dList.Length-1 || index < 0 {
		return nil
	}
	pHead := dList.Head
	for index > -1 {
		pHead = pHead.Next
		index--
	}
	return pHead
}

func (dList *DoubleLinkList) DeleteNodeAtIndex(index int) bool {
	if index > dList.Length-1 || index < 0 {
		return false
	}
	pHead := dList.Head
	for index > 0 {
		pHead = pHead.Next
		index--
	}
	if pHead.Next.Next != nil {
		pHead.Next.Next.Prev = pHead
	}
	pHead.Next = pHead.Next.Next
	dList.Length--
	return true
}
func (dList *DoubleLinkList) Each(f func(node DoubleLinkNode)){
	for node:=dList.Head;node!=nil;node=node.Next{
		f(*node)
	}
}