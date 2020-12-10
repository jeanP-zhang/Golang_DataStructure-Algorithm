package main

import "fmt"

//mergeable可合并的
type MergerAble interface {
	Merge(m2 MergerAble) MergerAble //用于合并的函数
}

//可对比的
type Comparable interface {
	Compare(c2 Comparable) int
}
type ArraySegmentTree struct {
	data []MergerAble
	tree []MergerAble //兼顾合并
}
type Integer int
type Integers []int

func (i Integer) Merge(i2 MergerAble) MergerAble {

	return Integer(i) + Integer(i2.(Integer))
}
func (i Integers) Merge(i2 MergerAble) MergerAble {
	newi2 := i2.(Integers)

	for j := 0; j < len(i); j++ {
		i[j] += newi2[j]
	}
	return i
}

func (i Integer) Compare(c2 Comparable) int {
	return int(i) - int(c2.(Integer))
}

//合并的接口
func TransInts(original []int) []MergerAble {
	ret := make([]MergerAble, len(original)) //分配内存
	for key, value := range original {
		ret[key] = Integer(value)
	}
	return ret
}

//线段树
func CreateSegmentTree(arr []MergerAble) *ArraySegmentTree {
	tree := make([]MergerAble, len(arr)*4) //开辟四倍内存
	st := &ArraySegmentTree{arr, tree}
	st.BuildSegmentTree(0, 0, len(arr)-1)
	return st
}
func (ast *ArraySegmentTree) LeftChild(index int) int {
	return index*2 + 1
}

func (ast *ArraySegmentTree) RightChild(index int) int {
	return index*2 + 2
}
func (ast *ArraySegmentTree) String() string {
	return fmt.Sprintln(ast.tree)
}

func (ast *ArraySegmentTree) Size() int {
	return len(ast.data)
}

//构造树
func (ast *ArraySegmentTree) BuildSegmentTree(index, left, right int) {
	if left == right {
		ast.tree[index] = ast.data[left] //插入第一个元素
	} else {
		leftchild := ast.LeftChild(index)
		rightchild := ast.RightChild(index)
		midchild := left + (right-left)/2 //取得中间数据
		ast.BuildSegmentTree(leftchild, left, midchild)
		ast.BuildSegmentTree(rightchild, midchild+1, right)               //反复构造
		ast.tree[index] = ast.tree[leftchild].Merge(ast.tree[rightchild]) //合并数据
	}
}

//查询
func (ast *ArraySegmentTree) Query(qleft, qright int) MergerAble {
	if qleft < 0 || qright < 0 || qleft > len(ast.data) || qright > len(ast.data) {
		panic("index out")
	}
	return ast.query(0, 0, len(ast.data)-1, qleft, qright)
}
func (ast *ArraySegmentTree) query(index, left, right, qleft, qright int) MergerAble {
	if left == right {
		return ast.tree[index] //返回，找到了
	} else {
		leftchild := ast.LeftChild(index)
		rightchild := ast.RightChild(index)
		midchild := left + (right-left)/2 //取得中间数据
		if qleft >= midchild+1 {

			return ast.query(rightchild, midchild+1, right, qleft, qright)
		} else {

			return ast.query(leftchild, left, midchild, qleft, qright)
		}
		leftRes := ast.query(leftchild, left, midchild, qleft, qright)
		rightRes := ast.query(rightchild, midchild+1, right, qleft, qright)
		return leftRes.Merge(rightRes)
	}
}

//设置
func (ast *ArraySegmentTree) Set(index int, e MergerAble) {
	if index < 0 || index >= len(ast.data) {
		panic("index out")
	} else {
		ast.data[index] = e
		ast.set(0, 0, len(ast.data)-1, index, e)
	}
}
func (ast *ArraySegmentTree) set(tree, left, right, index int, e MergerAble) {
	if left == right {
		ast.tree[tree] = e //恰好要插入的位置
	} else {
		leftchild := ast.LeftChild(index)
		rightchild := ast.RightChild(index)
		midchild := left + (right-left)/2 //取得中间数据
		if index >= midchild+1 {
			ast.set(rightchild, midchild+1, right, index, e)
		} else {
			ast.set(leftchild, left, midchild, index, e)
		}
		ast.tree[tree] = ast.tree[left].Merge(ast.tree[right])
	}
}
func main() {
	var data = []MergerAble{Integer(1), Integer(7), Integer(21), Integer(11), Integer(6)}
	fmt.Println(data)
	myTree := CreateSegmentTree(data)
	myTree.Set(0,Integer(999))
	fmt.Println(myTree.data)
	fmt.Println(myTree.tree)
}
