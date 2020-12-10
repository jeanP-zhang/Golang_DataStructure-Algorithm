package main

import (
	"errors"
	"fmt"
	"sort"
)

const (
	Inf    = int(^uint(0) >> 1) //正无穷大
	NegInf = -Inf - 1           //负数无穷大
)

//线段的左右
type segment struct {
	from int
	to   int
}

//区间
type interval struct {
	segment
	element interface{}
}
type node struct {
	segments    segment     //线段
	left, right *node       //二叉树左右孩子
	intervals   []*interval //指向指针
}
type Tree struct {
	base     []interval
	elements map[interface{}]struct{} //元素索引
	root     *node
}

//压入数据
func (t *Tree) Push(from, to int, element interface{}) {
	if to < from {
		from, to = to, from
	}
	if t.elements == nil {
		t.elements = make(map[interface{}]struct{})
	}
	//开辟内存
	t.elements[element] = struct{}{}
	t.base = append(t.base, interval{segment{from, to}, element})
}

func (t *Tree) Clear() {
	t.base = nil
	t.root = nil //清空
}
func (t *Tree) BuildTree() error {
	if len(t.base) == 0 {
		return errors.New("已经构造")
	}
	//插入数据
	leaves := elementArrayIntervals(t.endPoints())
	t.root = t.InsertNodes(leaves)
	for i := range t.base {
		t.root.InsertInterval(&t.base[i])
	}
	return nil
}
func (t *Tree) endPoints() []int {
	baselen := len(t.base) //取得长度
	endpointes := make([]int, baselen*2)
	//划分区间
	for i, intervals := range t.base {
		endpointes[i] = intervals.from
		endpointes[i+baselen] = intervals.to
	}
	sort.Sort(sort.IntSlice(endpointes)) //排序
	return removeDups(endpointes)
}

func (t *Tree) InsertNodes(leaves []segment) *node {
	var n *node
	if len(leaves) == 1 { //插入一个直接插入
		n = &node{segments: leaves[0]}
		n.left = nil
		n.right = nil
	}
	n = &node{segments: segment{leaves[0].from, leaves[0].to}}
	center := len(leaves) / 2 //取得中间数据
	n.left = t.InsertNodes(leaves[:center])
	n.right = t.InsertNodes(leaves[center:])
	return n

}

func (s *segment) contains(index int) bool {
	return s.from <= index && s.to >= index
}

//
func (t *Tree) Query(index int) (<-chan interface{}, error) {
	if t.root == nil {
		return nil, errors.New("树为空")
	}
	intervals := make(chan *interval) //构造管道
	//并发调用
	go func(t *Tree, index int, intervals chan *interval) {
		query(t.root, index, intervals)
		close(intervals)
	}(t, index, intervals)
	elements := make(chan interface{})
	go func(intervals chan *interval, elements chan interface{}) {
		defer close(elements)
		results := make(map[interface{}]struct{})
		for v := range intervals {
			_, alreadyFound := results[v.element] //找到
			if !alreadyFound {
				results[v.element] = struct{}{}
				elements <- v.element //压入
				if len(results) >= len(t.elements) {
					return //溢出
				}
			}
		}
	}(intervals, elements)

	return elements, nil
}

//查询数据 ，可以用并发实现
func query(node *node, index int, results chan<- *interval) {
	if node.segments.contains(index) { //判断数据是否在区间内
		for _, intervalx := range node.intervals {
			results <- intervalx
		}
		if node.left != nil {
			query(node.left, index, results)
		}
		if node.right != nil {
			query(node.right, index, results)

		}
	}
}

//区间判断
func (s *segment) subSetOf(other *segment) bool {
	return other.from <= s.from && other.to >= s.to
}

//区间判断，要么other包含s，要么s包含other
func (s *segment) intersectsWith(other *segment) bool {
	return (other.from <= s.from && other.to >= s.to) || (other.from >= s.from && other.to <= s.to)
}
func (n *node) InsertInterval(i *interval) {
	if n.segments.subSetOf(&i.segment) {
		if n.intervals == nil {
			n.intervals = make([]*interval, 0, 1) //开辟内存
		}
		n.intervals = append(n.intervals, i) //只有一个
	} else { //插入数据
		if n.left != nil && n.left.segments.intersectsWith(&i.segment) {
			n.left.InsertInterval(i)
		}
		if n.right != nil && n.right.segments.intersectsWith(&i.segment) {
			n.right.InsertInterval(i)
		}
	}
}

func (t *Tree) QueryIndex(index int) {

}

//区间划分
//[p1,p2,p3,----,pn]
//[p1,p2],[p3,p4]
//n,2n+1，2n+2
func elementArrayIntervals(endpoint []int) []segment {
	if len(endpoint) == 1 {
		return []segment{{endpoint[0], endpoint[0]}}
	} else {
		intervals := make([]segment, len(endpoint)*2-1)
		for i := 0; i < len(endpoint); i++ {
			intervals[i*2] = segment{endpoint[i], endpoint[i]}
			if i < len(endpoint)-1 {
				intervals[2*i+1] = segment{endpoint[i], endpoint[i]}
			}
		}
		return intervals
	}
}

func removeDups(sorted []int) (unqique []int) {
	unqique = make([]int, 0, len(sorted))
	unqique = append(unqique, sorted[0])
	prev := sorted[0] //取得前置数据
	for _, val := range sorted[1:] {
		if val != prev {
			unqique = append(unqique, val)
			prev = val
		}
	}
	return
}
func Traverse(node *node, depth int, enter, leave func(*node, int)) {
	if node == nil {
		return
	}
	if enter != nil {
		enter(node, depth) //递归函数
	}
	Traverse(node.left, depth+1, enter, leave)
	if leave != nil {
		leave(node, depth)
	}
}
func (n *node) print() {
	from := fmt.Sprintf("%d", n.segments.from) //打印区间
	switch n.segments.from {
	case Inf:
		from = "+00"
	case NegInf:
		from = "-00"
	}
	to := fmt.Sprintf("%d", n.segments.to) //打印区间
	switch n.segments.to {
	case Inf:
		to = "+00"
	case NegInf:
		to = "-00"
	}
	fmt.Printf("%s,%s", from, to) //打印数据
	if n.intervals != nil {
		fmt.Print("->[")
		for _, intervl := range n.intervals {
			fmt.Printf("(%v,%v)=[%v]", intervl.from, intervl.to, intervl.element)
		}
		fmt.Print("]")
	}
}
func (t *Tree) Print() {
	endPoints := len(t.base)*2 + 2
	leaves := endPoints*2 - 3
	height := 1 + log2(leaves)
	levels := make([][]*node, height+1) //层级
	fmt.Println("height:", height, "leaves:", leaves)
	Traverse(t.root, 0, func(n *node, i int) {
		levels[i] = append(levels[i], n)
	}, nil)
	for i, level := range levels {
		for j, n := range level {
			space(12 * (len(levels) - 1 - i))
			n.print()
			space(1 * (height - i))
			if j-1&2 == 0 {
				space(2)
			}
		}
		fmt.Println()
	}
}

//2^n -1
func log2(num int) int {
	if num == 0 {
		return NegInf
	}
	i := -1
	for num > 0 {
		num = num >> 1
		i++
	}
	return i //求指数
}
func space(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(" ") //打印空格显示层级
	}
}
func main() {
	mytree := new(Tree)
	mytree.Push(1, 10, "HELLO,WORD")
	mytree.BuildTree()
	//fmt.Println(mytree.QueryIndex(4))
	mytree.Print()
}
