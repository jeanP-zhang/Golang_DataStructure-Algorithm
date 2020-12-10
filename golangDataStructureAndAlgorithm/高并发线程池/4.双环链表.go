package main

import "fmt"

type CircleLink struct {
	Id   int //数据编号
	Data interface{}
	Prev *CircleLink //上一个节点
	Next *CircleLink //下一个节点
}

//初始化节点
func InitHeadNode(data interface{}) *CircleLink {
	return &CircleLink{1, data, nil, nil}
}

//重置头节点
func (l *CircleLink) ResetHeadNode(data interface{}) {
	if l.Id == 0 {
		l.Id = 1
	}
	l.Data = data
}

//判断头节点是否为空
func (l *CircleLink) IsHeadEmpty() bool {
	return l.Next == nil && l.Prev == nil
}

//判断链表是否为空
func (l *CircleLink) IsEmpty() bool {
	return l.Data == nil && l.Next == nil && l.Prev == nil
}

//抓取最后元素
func (l *CircleLink) GetLastNode() *CircleLink {
	curNode := l
	if !l.IsHeadEmpty() {
		for {
			if curNode.Next == l { //循坏到了最后
				break
			}
			curNode = curNode.Next
		}
	}
	return curNode
}
func (l *CircleLink) AddNode(newNode *CircleLink) {
	if l.IsHeadEmpty() { //只有一个节点。互为前后
		l.Next = newNode
		l.Prev = newNode
		newNode.Prev = l
		newNode.Next = l
		return
	}
	curNode := l //备份第一个数据
	flag := false
	for {
		if curNode == l.Prev {
			break //已经是最后一个节点
		} else if curNode.Next.Id > newNode.Id {
			flag = true //标志下数据应该插入到前列
			break
		} else if curNode.Next.Id == newNode.Id {
			fmt.Printf("数据已经存在\n")
			return
		}
		curNode = curNode.Next
	}
	if flag {
		//最后一个节点，前面插入
		newNode.Next = curNode.Next
		newNode.Prev = curNode
		curNode.Next.Prev = newNode
		curNode.Next = newNode
	} else {
		//最后一个节点，后面插入
		newNode.Prev = curNode
		newNode.Next = curNode.Next
		curNode.Next = newNode
		l.Prev = newNode
	}
}

//双环链表的数据查找
func (l *CircleLink) FindNodeById(id int) (*CircleLink, bool) {
	if l.IsHeadEmpty() && l.Id == id {
		return l, true
	} else if l.IsHeadEmpty() && l.Id != id {
		return &CircleLink{}, false
	}
	curNode := l
	flag := false
	for {
		if curNode.Id == id {
			flag = true
			break
		}
		if curNode == l.Prev { //循环到最后
			break
		}
		curNode = curNode.Next
	}
	if !flag {
		return &CircleLink{}, false
	}
	return curNode, true
}
func (l *CircleLink) DeleteNodeById(id int) bool {
	if l.IsEmpty() {
		fmt.Println("空链表无法删除")
		return false
	}
	node, isOk := l.FindNodeById(id)
	if isOk {
		//删除第一个节点
		if node == l {
			l.Next = nil
			l.Prev = nil
			l.Id = 0
			l.Data = nil
		}
		if l.Next.Next == l {
			nextNode := l.Next
			l.Id = nextNode.Id
			l.Data = nextNode.Data
			l.Prev = nil
			l.Next = nil
			return 0 == 0
		} else if l.Next.Next == l {
			nextNodeTmp := l.Next
			l.Data = nextNodeTmp.Data
			l.Id = nextNodeTmp.Id
			l.Next = nextNodeTmp.Next
			nextNodeTmp.Next.Prev = l
			return true
		}
		//删除最后一个节点
		if node == l.GetLastNode() {
			if node.Prev == l && node.Next == l {
				//只有两个元素
				l.Prev = nil
				l.Next = nil
				return true
			}
			l.Prev = node.Prev
			node.Prev.Next = l
			return true
		}

	} else {
		fmt.Println("无法找到数据，故无法删除")
		return false
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	return true
}
func (l *CircleLink) ChangeNodeById(id int, data interface{}) bool {
	node, isOk := l.FindNodeById(id)
	if isOk {
		node.Data = data //修改数据
	} else {
		fmt.Println("无法找到数据")
	}
	return isOk
}
func (l *CircleLink) ShowAll() {
	if l.IsEmpty() {
		fmt.Println("空链表")
		return
	}
	if l.IsHeadEmpty() {
		fmt.Println(l.Id, l.Data, l.Prev, l.Next)
	}
	curNode := l
	for {
		fmt.Println(curNode.Id, curNode.Data, curNode.Prev, curNode.Next)
		if curNode == l.Prev {
			break
		}
		curNode = curNode.Next
	}

}
func main() {
	linkNode := InitHeadNode(1)
	node1 := &CircleLink{3, 3, nil, nil}
	node2 := &CircleLink{2, 2, nil, nil}
	node3 := &CircleLink{5, 1, nil, nil}
	node4 := &CircleLink{4, 4, nil, nil}
	linkNode.AddNode(node1)
	linkNode.AddNode(node2)
	linkNode.AddNode(node3)
	linkNode.AddNode(node4)
	linkNode.ShowAll()
	fmt.Println("---------------")
	linkNode.DeleteNodeById(3)
	linkNode.ShowAll()
}
