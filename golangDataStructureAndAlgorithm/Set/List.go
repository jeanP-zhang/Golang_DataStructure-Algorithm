package main

type Object interface {
}
type MatchFun func(data1 Object, data2 Object) int //函数实现对比

type Node struct {
	Data Object
	Next *Node
}
type List struct {
	Size    uint64
	Head    *Node
	Tail    *Node //define your match function
	myMatch MatchFun
}

//define functions
func defaultMatch(data1 Object, data2 Object) int {
	if data2 == data1 {
		return 0
	} else {
		return 1
	}
}

//choose suitable match
func (list *List) match(data1, data2 Object) int {
	var matc MatchFun = nil
	if (*list).myMatch == nil {
		matc = defaultMatch
	} else {
		matc = (*list).myMatch
	}
	return matc(data1, data2)
}
func (list *List) creatNode(data Object) *Node {
	node := new(Node)
	node.Data = data
	node.Next = nil
	return node
}
func nextNode(node *Node) *Node {
	return node.Next
}
func (list *List) getHead() *Node {
	return list.Head
}
func (list *List) getTail() *Node {
	return list.Tail
}
func (node *Node) getData() Object {
	return node.Data
}
func (list *List) insertAfterNode(node *Node, data Object) bool {
	//TODO:
	return true
}
func (list *List) RemoveAt(index uint64) Object {
	size := list.GetSize()
	if index > size {
		return nil
	} else if size == 1 {
		node := list.getHead()
		list.Head = nil
		list.Tail = nil
		list.Size = 0
		return node.Data
	} else if index == 0 {
		node := list.getHead()
		list.Head = node.Next
		list.Size--
		return node.Data
	} else if index == size-1 {
		preNode := list.Head
		for i := uint64(2); i < size; i++ {
			preNode = preNode.Next
		}
		tail := list.getTail()
		list.Tail = preNode
		preNode.Next = nil
		list.Size--
		return tail.Data
	} else {
		preNode := list.Head
		for i := uint64(2); i < index; i++ {
			preNode = preNode.Next
		}
		node := preNode.Next
		nxtNode := node.Next
		node.Next = nxtNode
		list.Size--
		return node.Data
	}
}
func (list *List) Remove(data Object) bool {
	if data == nil || list.IsEmpty() {
		return false
	}
	head := list.getHead()
	if list.match(head.getData(), data) == 0 {
		list.Head = nextNode(head)
	} else {
		cur := head
		nxt := nextNode(head)
		for ; nxt != nil; nxt = nextNode(nxt) {
			if list.match(data, nxt.getData()) == 0 {
				cur.Next = nextNode(nxt)
				break
			}
			cur = nxt
		}
		if nxt == nil {
			return false
		}
	}
	list.Size--
	return true
}
func (list *List) IsMember(data Object) bool {
	if list.IsEmpty() {
		return false
	}
	//get head
	head := list.getHead()
	for i := head; i != nil; i = nextNode(i) {
		if list.match(data, i.getData()) == 0 {
			return true
		}
	}
	return false
}
func (list *List) Init(yourMatch ...MatchFun) {
	list.Size = 0
	list.Head = nil
	list.Tail = nil
	if len(yourMatch) == 0 {
		list.myMatch = nil
	} else {
		list.myMatch = yourMatch[0]
	}
}
func (list *List) GetSize() uint64 {
	return list.Size
}
func (list *List) IsEmpty() bool {
	return list.Size == 0
}
func (list *List) Append(data Object) bool {
	newItem := new(Node)
	newItem.Data = data
	newItem.Next = nil
	if list.Size == 0 {
		list.Head = newItem
		list.Tail = list.Head
	} else {
		oldNode := list.Tail
		oldNode.Next = newItem
		list.Tail = newItem
	}
	list.Size++
	return true
}
func (list *List) InsertAtHead(data Object) bool {
	newNode := list.creatNode(data)
	//insert head
	newNode.Next = list.getHead()
	list.Head = newNode
	list.Size++
	return true
}

//ToDo:First
//get the first data
func (list *List) First() Object {
	if list.GetSize() == 0 {
		return nil
	} else {
		return list.getHead().Data
	}
}

//get the last data
func (list *List) Last() Object {
	if list.GetSize() == 0 {
		return nil
	} else {
		return list.getTail().Data
	}
}
func (list *List) Next(curData Object) Object {
	//get head
	head := list.getHead()
	for i := head; i != nil; i = nextNode(i) {
		if list.match(curData, i.getData()) == 0 {
			nxt := nextNode(i)
			if nxt == nil {
				return nil
			} else {
				return nxt.getData()
			}
		}
	}
	return nil
}
func (list *List) GetAt(index uint64) Object {
	size := list.GetSize()
	if index >= size {
		return nil
	} else if index == 0 {
		return list.First()
	} else if index == size-1 {
		return list.Last()
	} else {
		item := list.getHead()
		for i := uint64(0); i < size; i++ {
			if i == index {
				break
			}
			item = item.Next
		}
		return item.getData()
	}
}
func (list *List) InsertAt(index uint64, data Object) bool {
	size := list.GetSize()
	if index > size {
		return false
	} else if index == size {
		return list.Append(data)
	} else if index == 0 {
		return list.InsertAtHead(data)
	} else {
		newNode := list.creatNode(data)
		prevIndex := index - 1
		prevItem := list.getHead()
		for i := uint64(0); i < size; i++ {
			if i == prevIndex {
				break
			}
			prevItem = prevItem.Next
		}
		newNode.Next = prevItem.Next
		prevItem.Next = newNode
		list.Size++
		return true
	}
}
func (list *List) Clear() {
	list.Init()
}
