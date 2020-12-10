package main

import "fmt"

type TreeNode struct {
	element interface{}
	left    *TreeNode
	right   *TreeNode
	npl     int
}
type PQ *TreeNode //优先队列
func NewLeftHeap(element interface{}) PQ {
	head := new(TreeNode)  //开辟内存
	head.element = element //数据初始化
	head.left = nil
	head.right = nil
	head.npl = 0
	return PQ(head)
}
func MergeSort(H1, H2 PQ) PQ {
	if H1.left == nil {
		H1.left = H2 //直接插入
	} else {
		H1.right = Merge(H1.right, H2)  //递归合并下一个节点
		if H1.left.npl < H1.right.npl { //处理层级的互换
			H1.left, H1.right = H1.right, H1.left
		}
		H1.npl = H1.right.npl + 1 //层级递增
	}
	return H1
}

//确保有序
func Merge(H1, H2 PQ) PQ {
	if H1 == nil {
		return H2
	}
	if H2 == nil {
		return H1
	}
	if H1.element.(int) < H2.element.(int) { //确保左边小于右边,如果改成大于号，则弹出最大值
		return MergeSort(H1, H2)
	} else {
		return MergeSort(H2, H1)
	}
}
func Insert(H PQ, data interface{}) PQ {
	insertNode := new(TreeNode) //新建一个节点
	insertNode.element = data
	insertNode.left = nil
	insertNode.right = nil
	insertNode.npl = 0
	H = Merge(insertNode, H)
	return H
}
func DeleteMin(H *TreeNode) (PQ, interface{}) {
	if H == nil {
		return nil, nil
	}
	leftHeap := H.left
	rigHtHeap := H.right
	value := H.element
	H = nil
	return Merge(leftHeap, rigHtHeap), value
}

//遍历大树叶
func PrintHQ(H PQ) {
	if H == nil {
		return
	}
	PrintHQ(H.left)
	PrintHQ(H.right)
	fmt.Println(H.element, "        ")
}
func main() {
	H := NewLeftHeap(3)
	H = Insert(H, 2)
	H = Insert(H, 1)
	H = Insert(H, 4)
	PrintHQ(H)
	H, data := DeleteMin(H)
	fmt.Println("min", data)
}
