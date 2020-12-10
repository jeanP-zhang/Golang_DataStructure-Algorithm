package main

import "fmt"

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func (this *ListNode) GetValue() interface{} {
	return this.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

//在某个节点后面插入节点
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	this.length++
	return true
}

//在某个节点前面插入节点
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == this.head {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true
}

//在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

//在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

//通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil
	this.length--
	return true
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

/*
单链表反转
时间复杂度：O(N)
*/
func (this *LinkedList) Reverse() {
	if nil == this.head || nil == this.head.next || nil == this.head.next.next {
		return
	}

	var pre *ListNode = nil
	cur := this.head.next
	for nil != cur {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}

	this.head.next = pre
}

/*
判断单链表是否有环
*/
func (this *LinkedList) HasCycle() bool {
	if nil != this.head {
		slow := this.head
		fast := this.head
		for nil != fast && nil != fast.next {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

/*
删除倒数第N个节点
*/
func (this *LinkedList) DeleteBottomN(n int) {
	if n <= 0 || nil == this.head || nil == this.head.next {
		return
	}

	fast := this.head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.next
	}

	if nil == fast {
		return
	}

	slow := this.head
	for nil != fast.next {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}

/*
获取中间节点
*/
func (this *LinkedList) FindMiddleNode() *ListNode {
	if nil == this.head || nil == this.head.next {
		return nil
	}
	if nil == this.head.next.next {
		return this.head.next
	}

	slow, fast := this.head, this.head
	for nil != fast && nil != fast.next {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

/*
两个有序单链表合并
*/
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if nil == l1 || nil == l1.head || nil == l1.head.next {
		return l2
	}
	if nil == l2 || nil == l2.head || nil == l2.head.next {
		return l1
	}

	l := &LinkedList{head: &ListNode{}}
	cur := l.head
	curl1 := l1.head.next
	curl2 := l2.head.next
	for nil != curl1 && nil != curl2 {
		if curl1.value.(int) > curl2.value.(int) {
			cur.next = curl2
			curl2 = curl2.next
		} else {
			cur.next = curl1
			curl1 = curl1.next
		}
		cur = cur.next
	}

	if nil != curl1 {
		cur.next = curl1
	} else if nil != curl2 {
		cur.next = curl2
	}

	return l
}

func main() {
	list := NewLinkedList()
	list.InsertToTail(2)
	list.InsertToHead(1)
	list.InsertToTail(3)
	list.InsertToTail(4)
	list.Reverse()
	list.Print()
	if list.HasCycle() {
		fmt.Println("has cycle")
	}
}
