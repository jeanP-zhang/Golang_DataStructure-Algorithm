package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

//节点
type Node struct {
	key    interface{}
	val    interface{}
	Parent *Node //父亲节点
	Left   *Node //左孩子节点
	Right  *Node //右孩子节点
}

//接口
type SPlayTree interface {
	SetRoot(n *Node)                //设置根节点
	GetRoot() *Node                 //返回根节点
	Ord(key1, key2 interface{}) int //排序
}
type MySPlayTree struct {
	root *Node
}

func (ST *MySPlayTree) SetRoot(n *Node) {
	ST.root = n
}
func (ST *MySPlayTree) GetRoot() *Node {
	return ST.root
}
func (ST *MySPlayTree) Ord(key1, key2 interface{}) int {
	if key1.(int) < key2.(int) {
		return -1
	} else if key1.(int) == key2.(int) {
		return 0
	}
	return +1
}

//Find,二叉树
/*查找一般均为递归编程*/
//search二叉树，返回节点
func Search(ST SPlayTree, key interface{}) *Node {
	return SearchNode(ST, key, ST.GetRoot())
}
func SearchNode(ST SPlayTree, key interface{}, n *Node) *Node {
	if n == nil {
		return nil
	} else {
		switch ST.Ord(key, n.key) {
		case 0:
			return n //相等
		case 1: //大于
			return SearchNode(ST, key, n.Right) //跳转到左边
		case -1: //小于
			return SearchNode(ST, key, n.Left) //跳转到左边
		}
		return nil
	}
}

//查找，实现返回数据
func Find(ST SPlayTree, key interface{}) interface{} {
	return FindNode(ST, key, ST.GetRoot())
}
func FindNode(ST SPlayTree, key interface{}, n *Node) interface{} {
	if n == nil {
		return nil
	} else {
		switch ST.Ord(key, n.key) {
		case 0:
			return n.val //相等
		case 1: //大于
			return FindNode(ST, key, n.Right) //跳转到左边
		case -1: //小于
			return FindNode(ST, key, n.Left) //跳转到左边
		}
		return nil
	}
}

//插入数据
func Insert(ST SPlayTree, key interface{}, value interface{}) (*Node, error) {
	if Search(ST, key) != nil {
		return nil, errors.New("要插入的已经存在")
	}
	n := InsertNode(ST, key, value, ST.GetRoot()) //调用插入
	//伸展
	fmt.Println(n)
	return n, nil
}
func InsertNode(ST SPlayTree, key interface{}, value interface{}, n *Node) *Node {
	if n == nil {
		_n := new(Node)
		_n.key = key
		_n.val = value
		ST.SetRoot(_n)
		return ST.GetRoot() //伸展树为空，直接插入
	}
	switch ST.Ord(key, n.key) {
	case 0:
		return nil
	case 1: //数据已经存在
		if n.Right == nil {
			n.Right = new(Node)
			n.Right.key = key
			n.Right.val = value
			n.Right.Parent = n //设定父亲节点
			return n.Right
		} else {
			return InsertNode(ST, key, value, n.Right) //插入数据
		}
	case -1: //
		if n.Left == nil {
			n.Left = new(Node)
			n.Left.key = key
			n.Left.val = value
			n.Left.Parent = n //设定父亲节点
			return n.Left
		} else {
			return InsertNode(ST, key, value, n.Left) //插入数据
		}
	}
	return nil
}

//删除数据
func Delete(ST SPlayTree, key interface{}) error {

	if n := Search(ST, key); n == nil {
		return errors.New("要删除的不存在")
	} else {
		p := n.Parent //保存父亲节点
		if n.Left != nil {
			iop := InOrderPredecessor(n.Left) //取得左边的最大值
			Swap(n, iop)                      //交换节点
			Remove(ST, iop)                   //删除节点
		} else if n.Right != nil {
			ios := InOrderSuccessor(n.Right) //寻找右边的最小值
			Swap(n, ios)
			Remove(ST, ios)
		} else {
			Remove(ST, n)
		}
		if p != nil {
			//伸展
		}
		return nil
	}

}

//
func Remove(ST SPlayTree, n *Node) {
	var isRoot bool
	var isLeft bool
	isRoot = n == ST.GetRoot()
	if isRoot != true {
		isLeft = n == n.Parent.Left //判断是否为左边节点
	}
	if isRoot != true {
		if isLeft == true {
			if n.Left != nil {
				n.Parent.Left = n.Left
				n.Left.Parent = n.Parent
			} else if n.Right != nil {
				n.Parent.Left = n.Right
				n.Right.Parent = n.Parent
			} else {
				n.Parent.Left = nil //叶子节点
			}
		} else {
			if n.Left != nil {
				n.Parent.Right = n.Left
				n.Left.Parent = n.Parent
			} else if n.Right != nil {
				n.Parent.Right = n.Right
				n.Right.Parent = n.Parent
			} else {
				n.Parent.Right = nil //叶子节点
			}
		}
	}
	n = nil
}

//交换,数据交换
func Swap(n1, n2 *Node) {
	n1.key, n2.key = n2.key, n1.key
	n1.val, n2.val = n2.val, n1.val
}

//取得最大值
func InOrderPredecessor(n *Node) *Node {
	if n.Right == nil {
		return n
	} else {
		return InOrderPredecessor(n.Right)
	}
}

//取得最小

func InOrderSuccessor(n *Node) *Node {
	if n.Left == nil {
		return n
	} else {
		return InOrderSuccessor(n.Left)
	}
}

//伸展
//树的左旋和右旋
func Splay(ST SPlayTree, n *Node) {
	for n != ST.GetRoot() {
		if n.Parent == ST.GetRoot() && n.Parent.Left == n {
			ZigL(ST, n) //根节点，只需要左旋
		} else if n.Parent == ST.GetRoot() && n.Parent.Right == n {
			ZigR(ST, n) //根节点，只需要右旋
		} else if n.Parent.Left == n && n.Parent.Parent.Left == n.Parent {
			ZigZigL(ST, n)
		} else if n.Parent.Right == n && n.Parent.Parent.Right == n.Parent {
			ZigZigR(ST, n)
		} else if n.Parent.Right == n && n.Parent.Parent.Left == n.Parent {
			ZigLZigR(ST, n)
		} else {
			ZigRZigL(ST, n)
		}
	}
}

//左旋
func ZigL(ST SPlayTree, n *Node) {
	n.Parent.Left = n.Right //存储左边的数据,n到根节点
	if n.Right != nil {
		n.Right.Parent = n.Parent
	}
	n.Parent.Parent = n
	n.Right = n.Parent
	n.Parent = nil
	ST.SetRoot(n)
}

//右旋
func ZigR(ST SPlayTree, n *Node) {
	n.Parent.Right = n.Left //存储左边的数据,n到根节点
	if n.Left != nil {
		n.Left.Parent = n.Parent
	}
	n.Parent.Parent = n
	n.Left = n.Parent
	n.Parent = nil
	ST.SetRoot(n)
}

//双层左旋
func ZigZigL(ST SPlayTree, n *Node) {
	gg := n.Parent.Parent.Parent //访问太爷
	var isRoot bool
	var isLeft bool
	if gg == nil {
		isRoot = true
	} else {
		isRoot = false
		isLeft = gg.Left == n.Parent.Parent
	}
	n.Parent.Parent.Left = n.Parent.Right //备份left
	if n.Parent.Right != nil {
		n.Parent.Right.Parent = n.Parent.Parent
	}
	n.Parent.Left = n.Right
	if n.Right != nil {
		n.Right.Parent = n.Parent
	}
	n.Parent.Right = n.Parent.Parent
	n.Parent.Parent.Parent = n.Parent
	n.Right = n.Parent
	n.Parent.Parent = n
	n.Parent = gg
	//判断树，

	if isRoot {
		ST.SetRoot(n)
	} else if isLeft {
		gg.Left = n
	} else {
		gg.Right = n
	}

}

//双层右旋
func ZigZigR(ST SPlayTree, n *Node) {
	gg := n.Parent.Parent.Parent //访问太爷
	var isRoot bool
	var isLeft bool
	if gg == nil {
		isRoot = true
	} else {
		isRoot = false
		isLeft = gg.Left == n.Parent.Parent
	}
	n.Parent.Parent.Left = n.Parent.Right //备份left
	if n.Parent.Right != nil {
		n.Parent.Right.Parent = n.Parent.Parent
	}
	n.Parent.Left = n.Right
	if n.Right != nil {
		n.Right.Parent = n.Parent
	}
	n.Parent.Right = n.Parent.Parent
	n.Parent.Parent.Parent = n.Parent
	n.Right = n.Parent
	n.Parent.Parent = n
	n.Parent = gg
	//判断树，

	if isRoot {
		ST.SetRoot(n)
	} else if isLeft {
		gg.Left = n
	} else {
		gg.Right = n
	}

}

//先左旋再右旋
func ZigLZigR(ST SPlayTree, n *Node) {
	gg := n.Parent.Parent.Parent //访问太爷
	var isRoot bool
	var isLeft bool
	if gg == nil {
		isRoot = true
	} else {
		isRoot = false
		isLeft = gg.Left == n.Parent.Parent
	}
	//先左再右
	n.Parent.Parent.Left = n.Parent.Right //备份left
	if n.Right != nil {
		n.Right.Parent = n.Parent.Parent

	}
	n.Parent.Right = n.Left
	if n.Right != nil {
		n.Right.Parent = n.Parent
	}
	n.Left = n.Parent
	n.Right = n.Parent.Parent
	n.Parent.Parent.Parent = n
	n.Parent.Parent = n
	n.Parent = gg
	//判断树，
	if isRoot {
		ST.SetRoot(n)
	} else if isLeft {
		gg.Left = n
	} else {
		gg.Right = n
	}

}

//先右旋再左旋
func ZigRZigL(ST SPlayTree, n *Node) {
	gg := n.Parent.Parent.Parent //访问太爷
	var isRoot bool
	var isLeft bool
	if gg == nil {
		isRoot = true
	} else {
		isRoot = false
		isLeft = gg.Left == n.Parent.Parent
	}
	//先左再右
	n.Parent.Parent.Right = n.Parent.Left //备份left
	if n.Left != nil {
		n.Left.Parent = n.Parent.Parent

	}
	n.Parent.Left = n.Right
	if n.Right != nil {
		n.Right.Parent = n.Parent
	}
	n.Right = n.Parent
	n.Left = n.Parent.Parent
	n.Parent.Parent.Parent = n
	n.Parent.Parent = n
	n.Parent = gg
	//判断树，

	if isRoot {
		ST.SetRoot(n)
	} else if isLeft {
		gg.Left = n
	} else {
		gg.Right = n
	}

}

//显示树
func Print(ST SPlayTree) {
	PrintNode(ST.GetRoot(), 0) //打印树枝
}
func PrintNode(node *Node, level int) {
	if node == nil {
		return
	}
	//先序遍历
	fmt.Println(strings.Repeat("-", 3*level), node.key, node.val) //打印数据
	PrintNode(node.Left, level+1)
	PrintNode(node.Right, level+1)
}
func main() {
	St := new(MySPlayTree)
	for i := 0; i < 36; i++ {
		_, err := Insert(St, rand.Int()%100, "hello ,qiangqiang QQ 123456789")
		if err != nil {
			fmt.Println(err)
		} else {
			Print(St)
		}
	}
	for i := 0; i < 36; i++ {
		err := Delete(St, i)
		if err != nil {
			fmt.Println(err)
		} else {
			Print(St)
		}
	}
}
