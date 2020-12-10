package main

import (
	"bytes"
	"container/list"
	"fmt"
	"strconv"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}
type BinaryTree struct {
	root *Node //根节点
	Size int   //数据的数量
}

//新建一个二叉树
func NewBinaryTree() *BinaryTree {
	bst := &BinaryTree{nil, 0}
	return bst
}

//获取二叉树的大小
func (bst *BinaryTree) getSize() int {
	return bst.Size
}

//判断二叉树是否为空
func (bst *BinaryTree) IsEmpty() bool {
	return bst.Size == 0
}

//根节点插入
func (bst *BinaryTree) Adds(data int) {
	bst.root = bst.Add(bst.root, data)
}

//插入节点
func (bst *BinaryTree) Add(n *Node, data int) *Node {
	if n == nil {
		bst.Size++
		return &Node{Data: data}
	} else {
		if data < n.Data {
			n.Left = bst.Add(n.Left, data) //比我小，左边
		} else if data > n.Data {
			n.Right = bst.Add(n.Right, data)
		}
		return n
	}
}
func (bst *BinaryTree) isIn(n *Node, data int) bool {
	if n == nil {
		return false
	}
	if data == n.Data {
		return true
	} else if data < n.Data {
		return bst.isIn(n.Left, data)
	} else {
		return bst.isIn(n.Right, data)
	}
}
func (bst *BinaryTree) IsIn(data int) bool {
	return bst.isIn(bst.root, data)
}
func (bst *BinaryTree) FindMax() int {
	if bst.Size == 0 {
		panic("二叉树为空")
	}
	node := bst.root
	for node.Right != nil {
		node = node.Right
	}
	return node.Data
}
func (bst *BinaryTree) FindsMax() int {
	if bst.Size == 0 {
		panic("二叉树为空")
	}
	//node:=bst.root
	//for node.Right!=nil{
	//	node=node.Right
	//}
	//return node.Data
	return bst.findMax(bst.root).Data
}
func (bst *BinaryTree) findMax(n *Node) *Node {
	if n.Right == nil {
		return n
	} else {
		return bst.findMax(n.Right)
	}
}
func (bst *BinaryTree) FindMin() int {
	if bst.Size == 0 {
		panic("二叉树为空")
	}
	return bst.findMin(bst.root).Data
}
func (bst *BinaryTree) findMin(n *Node) *Node {
	if n.Right == nil {
		return n
	} else {
		return bst.findMin(n.Left)
	}
}

//前序遍历
func (bst *BinaryTree) PreOrder() {
	bst.preOrder(bst.root)
}
func (bst *BinaryTree) preOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.Data)
	bst.preOrder(n.Left)
	bst.preOrder(n.Right)
}

//中序遍历
func (bst *BinaryTree) InOrder() {
	bst.inOrder(bst.root)
}
func (bst *BinaryTree) inOrder(n *Node) {
	if n == nil {
		return
	}
	bst.inOrder(n.Left)
	fmt.Println(n.Data)
	bst.inOrder(n.Right)
}

//后序遍历
func (bst *BinaryTree) PostOrder() {
	bst.postOrder(bst.root)
}
func (bst *BinaryTree) postOrder(n *Node) {
	if n == nil {
		return
	}
	bst.postOrder(n.Left)
	bst.postOrder(n.Right)
	fmt.Println(n.Data)
}
func (bst *BinaryTree) String() string {
	var buffer bytes.Buffer
	bst.GenerateBSTString(bst.root, 0, &buffer)
	return buffer.String()
}
func (bst *BinaryTree) GenerateBSTString(node *Node, depth int, buffer *bytes.Buffer) {
	if node == nil {
		buffer.WriteString(bst.GenerateDepthString(depth) + "nil\n") //空节点
		return
	}
	//写入字符串
	buffer.WriteString(bst.GenerateDepthString(depth) + strconv.Itoa(node.Data) + "\n")
	bst.GenerateBSTString(node.Left, depth+1, buffer)
	bst.GenerateBSTString(node.Right, depth+1, buffer)
}
func (bst *BinaryTree) GenerateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--") //深度为0，深度为1，深度为2
	}
	return buffer.String()
}
func (bst *BinaryTree) RemoveMin() int {
	ret := bst.FindMin()
	bst.root = bst.removeMin(bst.root)
	return ret
}
func (bst *BinaryTree) removeMin(node *Node) *Node {
	if node.Left == nil {
		//删除
		rightNode := node.Right //备份右边节点
		bst.Size--
		return rightNode
	}
	node.Left = bst.removeMin(node.Left)
	return node
}

func (bst *BinaryTree) removeMax(node *Node) *Node {
	if node.Right == nil {
		//删除
		leftNode := node.Left //备份右边节点
		bst.Size--
		return leftNode
	}
	node.Right = bst.removeMin(node.Right)
	return node
}
func (bst *BinaryTree) RemoveMax() int {
	ret := bst.FindMax()
	bst.root = bst.removeMax(bst.root)
	return ret
}
func (bst *BinaryTree) Remove(data int) {
	bst.root = bst.remove(bst.root, data)
}
func (bst *BinaryTree) remove(n *Node, data int) *Node {
	if n == nil {
		return nil
	}
	if data < n.Data {
		n.Left = bst.remove(n.Left, data)
		return n
	} else if data > n.Data {
		n.Right = bst.remove(n.Right, data)
		return n
	} else {
		if n.Left == nil {
			rightNode := n.Right //备份右边节点
			n.Right = nil
			bst.Size--
			return rightNode
		}
		if n.Right == nil {
			leftNode := n.Left //备份右边节点
			n.Left = nil       //处理节点返回
			bst.Size--         //删除
			return leftNode
		}
		//左右节点都不为空
		okNode := bst.findMin(n.Right)
		okNode.Right = bst.removeMin(n.Right)
		okNode.Left = n.Left //删除
		n.Left, n.Right = nil, nil
		return okNode
	}
}
func (bst *BinaryTree) InOrderNoRecursion() []int {
	myBst := bst.root
	mystack := list.New() //生成一个栈
	res := make([]int, 0) //生成一个数组，容纳中序的数据
	for myBst != nil || mystack.Len() != 0 {
		for myBst != nil {

			mystack.PushBack(myBst)
			myBst = myBst.Left
		}
		if mystack.Len() != 0 {
			v := mystack.Back()
			myBst = v.Value.(*Node)
			res = append(res, myBst.Data) //压入数据
			myBst = myBst.Right
			mystack.Remove(v) //删除
		}
	}
	return res
}
func (bst *BinaryTree) preOrderNoRecursion() []int {
	myBst := bst.root
	mystack := list.New() //生成一个栈
	res := make([]int, 0) //生成一个数组，容纳中序的数据
	for myBst != nil || mystack.Len() != 0 {
		for myBst != nil {
			res = append(res, myBst.Data) //压入数据
			mystack.PushBack(myBst)
			myBst = myBst.Left
		}
		if mystack.Len() != 0 {
			v := mystack.Back()
			myBst = v.Value.(*Node)

			myBst = myBst.Right
			mystack.Remove(v) //删除
		}
	}
	return res
}
func (bst *BinaryTree) postOrderNoRecursion() []int {
	myBst := bst.root
	mystack := list.New() //生成一个栈
	res := make([]int, 0) //生成一个数组，容纳中序的数据
	var PreVisited *Node  //提前访问的节点
	for myBst != nil || mystack.Len() != 0 {
		for myBst != nil {
			//压入数据
			mystack.PushBack(myBst)
			myBst = myBst.Left
		}
		v := mystack.Back() //取出节点
		top := v.Value.(*Node)
		if (top.Left == nil && top.Right == nil) || (top.Right == nil && PreVisited == top.Left) || (PreVisited == top.Right) {
			res = append(res, top.Data)
			PreVisited = top
			mystack.Remove(v)
		} else {
			myBst = top.Right
		}
	}
	return res
}
func (bst *BinaryTree) LevelShow() {
	bst.levelShow(bst.root)
}

func (bst *BinaryTree) levelShow(n *Node) {
	myQueue := list.New() //新建一个list模拟队列
	myQueue.PushBack(n)
	for myQueue.Len() > 0 {
		left := myQueue.Front() //前面取出数据
		right := left.Value
		myQueue.Remove(left) //删除
		if v, ok := right.(*Node); ok && v != nil {
			fmt.Println(v.Data)
			myQueue.PushBack(v.Right)
		}
	}
}

func (bst *BinaryTree) stackShow(n *Node) {
	myQueue := list.New() //新建一个list模拟队列
	myQueue.PushBack(n)
	for myQueue.Len() > 0 {
		left := myQueue.Back() //前面取出数据
		right := left.Value
		myQueue.Remove(left) //删除
		if v, ok := right.(*Node); ok && v != nil {
			fmt.Println(v.Data)
			myQueue.PushBack(v.Left)
			myQueue.PushBack(v.Right)
		}
	}
}
func (bst *BinaryTree) FindLowestAncestor(root *Node, nodeA *Node, nodeB *Node) *Node {
	if bst.root == nil {
		return nil
	}
	if root == nodeA || root == nodeB {
		return root //有一个节点是根节点
	}
	left := bst.FindLowestAncestor(root.Left, nodeA, nodeB)
	right := bst.FindLowestAncestor(root.Right, nodeA, nodeB)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return right
	} else {
		return left
	}
}
func (bst *BinaryTree) GetDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	lengthLeft := bst.GetDepth(root.Left)
	rightLength := bst.GetDepth(root.Right)
	if lengthLeft > rightLength {
		return lengthLeft + 1
	} else {
		return rightLength + 1
	}

}
