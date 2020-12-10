package main

import (
	"fmt"
	"math/rand"
)

//B树的节点
type BTreeNode struct {
	Leaf     bool         //是否叶子
	N        int          //分支的数量
	keys     []int        //存储数据
	Children []*BTreeNode //指向自己的多个分支节点
}

//新建一个节点
func NewBtreeNode(n int, branch int, leaf bool) *BTreeNode {
	return &BTreeNode{leaf,
		n,
		make([]int, branch*2-1),
		make([]*BTreeNode, branch*2)}
}

type BTree struct {
	Root   *BTreeNode //根节点
	branch int        //分支的数量
}

//搜索一个节点
func (btreenode *BTreeNode) Search(key int) (myNode *BTreeNode, idx int) {
	i := 0
	//找到合适的位置,最后一个小于key的,i之后的就是大于等于
	for i < btreenode.N && btreenode.keys[i] < key { //搜索B树的节点
		i += 1
	}
	if i < btreenode.N && btreenode.keys[i] == key {
		myNode, idx = btreenode, i //找到了
	} else if btreenode.Leaf == false {
		//进入孩子叶子继续搜索
		myNode, idx = btreenode.Children[i].Search(key)
	}
	idx = i
	return
}

//B树

func (parent *BTreeNode) Split(branch int, idx int) {
	full := parent.Children[idx]                         //孩子节点
	newnode := NewBtreeNode(branch-1, branch, full.Leaf) //新建一个节点,备份
	for i := 0; i < branch-1; i++ {
		newnode.keys[i] = full.keys[i+branch]         //数据的移动，跳过一个分支
		newnode.Children[i] = full.Children[i+branch] //
	}
	newnode.Children[branch-1] = full.Children[branch*2-1] //处理最后
	full.N = branch - 1                                    //新增一个key到children
	for i := parent.N; i > idx; i-- {
		parent.Children[i] = parent.Children[i-1]
		parent.keys[i+1] = parent.keys[i] //从后往前移动
	}
	parent.keys[idx] = full.keys[branch-1]
	parent.Children[idx+1] = newnode //插入数据，增加总量
	parent.N++
}
func (btreenode *BTreeNode) String() string {
	return fmt.Sprintf("{n=%d,leaf=%v,children=%v\n}", btreenode.N, btreenode.keys, btreenode.Children)
}

//节点插入数据
func (btreenode *BTreeNode) InsertNonFull(branch int, key int) {
	if btreenode == nil {
		return
	}
	i := btreenode.N    //记录叶子节点的总量
	if btreenode.Leaf { //是否为叶子
		for i > 0 && key < btreenode.keys[i-1] {
			btreenode.keys[i] = btreenode.keys[i-1] //从后往前移动
			i--
		}
		btreenode.keys[i] = key //插入数据
		btreenode.N++           //总量+1
	} else {
		for i > 0 && key < btreenode.keys[i-1] {
			i--
		}
		c := btreenode.Children[i] //找到下标
		if c != nil && c.N == 2*branch-1 {
			btreenode.Split(branch, i) //切割
			if key > btreenode.keys[i] {
				i++
			}
		}
		btreenode.Children[i].InsertNonFull(branch, key) //递归插入到孩子叶子
	}
}
func (tree *BTree) Insert(key int) {
	root := tree.Root
	if root.N == 2*tree.branch-1 {
		s := NewBtreeNode(0, tree.branch, false)
		tree.Root = s //新建一个节点，备份根节点
		s.Children[0] = root
		s.Split(tree.branch, 0) //拆分，整合
		root.InsertNonFull(tree.branch, key)
	} else {
		root.InsertNonFull(tree.branch, key)
	}
}

//查找
func (tree *BTree) Search(key int) (n *BTreeNode, idx int) {
	return tree.Root.Search(key)
}

//返回字符串
func (tree *BTree) String() string {
	return tree.Root.String() //返回树的字符串
}
func NewBTree(branch int) *BTree {
	node := NewBtreeNode(0, branch, true)
	return &BTree{node, branch}
}
func main() {
	mybtree := NewBTree(100)
	for i := 1000; i > 0; i-- {
		mybtree.Insert(rand.Int() % 300)
	}
	fmt.Println(mybtree.String())
}
