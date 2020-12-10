package main

import (
	"code.qiangqiang.com/studygo/golang-数据结构与算法/KDTree/KDrange"
	"code.qiangqiang.com/studygo/golang-数据结构与算法/KDTree/points"
	pq "code.qiangqiang.com/studygo/golang-数据结构与算法/KDTree/priorityQueue"
	"fmt"
	"math"
	"sort"
)

type KDPoint points.Points
type node struct {
	KDPoint
	Left  *node
	Right *node
}
type KDTree struct {
	root *node
}

func NewKDTree(point []KDPoint) *KDTree {
	return &KDTree{root: newKDTree(point, 0)} //从0维度开始
}
func newKDTree(point []KDPoint, axis int) *node {
	if len(point) == 0 {
		return nil
	}
	if len(point) == 1 {
		return &node{point[0], nil, nil}
	}
	sort.Sort(&byDimension{axis, point})
	mid := len(point) / 2                     //取得中间值
	root := point[mid]                        //设定根节点
	nextDim := (axis + 1) % root.Dimensions() //1,2,3,4->1维度循环

	return &node{
		KDPoint: root,
		Left:    newKDTree(point[:mid], nextDim),
		Right:   newKDTree(point[mid:], nextDim),
	}
}

func (kt *KDTree) Strings() string {
	return fmt.Sprintf("[s]", PrintTreeNode(kt.root))
}

//打印树节点
func PrintTreeNode(n *node) string {
	if n != nil && (n.Left != nil || n.Right != nil) {
		return fmt.Sprintf("[%s %s %s]", PrintTreeNode(n.Left), n.Strings(), PrintTreeNode(n.Right))
	}
	return fmt.Sprintf("%s", n)
}
func (kt *KDTree) Inserts(p KDPoint) {
	if kt.root == nil {
		kt.root = &node{p, nil, nil}
	} else {
		kt.root.Insert(p, 0)
	}
}

//实现了删除节点
func (kt *KDTree) Remove(p KDPoint) KDPoint {
	if kt.root == nil || p == nil {
		return nil
	}
	n, sub := kt.root.Remove(p, 0)
	if n == kt.root {
		kt.root = sub
	}
	if n == nil {
		return nil
	}
	return n.KDPoint
}
func (kt *KDTree) Pointss() []KDPoint {
	if kt.root == nil {
		return []KDPoint{}
	}
	return kt.root.Points()
}
func (kt *KDTree) Balance() {
	kt.root = newKDTree(kt.Pointss(), 0) //平衡
}

func (kt *KDTree) RangeSearch(r KDrange.Range) []KDPoint {
	if kt.root == nil || kt == nil || len(r) != kt.root.Dimensions() {
		return []KDPoint{}
	}
	return kt.root.RangeSearch(r, 0)
}

func (kt *KDTree) KNN(p KDPoint, k int) []KDPoint {
	if kt.root == nil || p == nil || k == 0 {
		return []KDPoint{}
	}
	nearestPQ := pq.NewPriorityQueue(pq.WithMinPrioSize(k)) //k表示数据的核心
	knn(p, k, kt.root, 0, nearestPQ)                        //遍历所有KD树
	pointss := make([]KDPoint, 0, k)
	for i := 0; i < k && 0 < nearestPQ.Len(); i++ {
		o := nearestPQ.PopLowest().(*node).KDPoint
		pointss = append(pointss, o) //追加
	}
	return pointss
}
func knn(p KDPoint, k int, start *node, curAxis int, nearestPQ *pq.PriorityQueue) {
	if p == nil || k == 0 || start == nil {
		return
	}
	var path []*node //路径
	curNode := start //当前节点
	//	向下移动
	for curNode != nil {
		path = append(path, curNode) //记录路径
		if p.Dimension(curAxis) < curNode.Dimension(curAxis) {
			curNode = curNode.Left
		} else {
			curNode = curNode.Right
		}
		curAxis = (curAxis + 1) % p.Dimensions()
	}
	//向上移动,维度倒退
	curAxis = (curAxis - 1 + p.Dimensions()) % p.Dimensions()
	for path, curNode := popLast(path); curNode != nil; path, curNode = popLast(path) {
		//计算当前距离
		curDistance := distance(p, curNode)               //计算距离
		checkedDistance := GetKthDistance(nearestPQ, k-1) //取出第k-1个距离
		if curDistance < checkedDistance {
			nearestPQ.Insert(curNode, curDistance)           //插入当前节点，距离
			checkedDistance = GetKthDistance(nearestPQ, k-1) //逐步求精
		} //淘汰长距离
		if planeDistance(p, curNode.Dimension(curAxis), curAxis) < checkedDistance {
			var next *node
			if p.Dimension(curAxis) < curNode.Dimension(curAxis) {
				next = curNode.Right
			} else {
				next = curNode.Left
			}
			knn(p, k, next, (curAxis+1)%p.Dimensions(), nearestPQ) //knn算法
		}
		curAxis = (curAxis - 1 + p.Dimensions()) % p.Dimensions()
	}
}

//定义数据结构
type byDimension struct {
	dimension int
	pointed   []KDPoint
}

func (bd *byDimension) Len() int {
	return len(bd.pointed)
}

//设定自身大小规则
func (bd *byDimension) Less(i, j int) bool {
	return bd.pointed[i].Dimension(bd.dimension) < bd.pointed[j].Dimension(bd.dimension)
}
func (bd *byDimension) Swap(i, j int) {
	bd.pointed[i], bd.pointed[j] = bd.pointed[j], bd.pointed[i]
}

//显示节点
func (n *node) Strings() string {
	return fmt.Sprintf("%v", n.KDPoint)
}

//返回所有节点
//中序遍历
func (n *node) Points() []KDPoint {
	var pointss []KDPoint
	if n.Left != nil {
		pointss = n.Left.Points()
	}
	pointss = append(pointss, n.KDPoint) //中序遍历
	if n.Right != nil {
		pointss = append(pointss, n.Right.Points()...)
	}
	return pointss
}
func (n *node) Insert(p KDPoint, axis int) {
	if p.Dimension(axis) < n.KDPoint.Dimension(axis) {
		if n.Left == nil {
			n.Left = &node{KDPoint: p, Left: nil, Right: nil}
		} else {
			n.Left.Insert(p, (axis+1)%n.KDPoint.Dimensions())
		}
	} else {
		if n.Right == nil {
			n.Right = &node{KDPoint: p, Left: nil, Right: nil}
		} else {
			n.Right.Insert(p, (axis+1)%n.KDPoint.Dimensions())
		}
	}
}

//按照维度查找最小
func (n *node) FindMin(axis int, smallest *node) *node {
	if smallest == nil || n.Dimension(axis) < smallest.Dimension(axis) {
		smallest = n
	}
	if n.Left != nil {
		smallest = n.Left.FindMin(axis, smallest)
	}

	if n.Right != nil {
		smallest = n.Right.FindMin(axis, smallest)
	}
	return smallest
}

func (n *node) FindMax(axis int, biggest *node) *node {

	if biggest == nil || n.Dimension(axis) > biggest.Dimension(axis) {
		biggest = n
	}
	if n.Left != nil {
		biggest = n.Left.FindMin(axis, biggest)
	}

	if n.Right != nil {
		biggest = n.Right.FindMin(axis, biggest)
	}
	return biggest
}

//返回节点，替换节点
func (n *node) Remove(p KDPoint, axis int) (*node, *node) {
	for i := 0; i < n.Dimensions(); i++ {
		if n.Dimension(i) != p.Dimension(i) {
			if n.Left != nil {
				//左子树维度循环
				returnNode, subNode := n.Left.Remove(p, (axis+1)%n.Dimensions())
				if returnNode != nil {
					if returnNode == n.Left {
						n.Left = subNode
					}
					return returnNode, nil
				}
			}
			if n.Right != nil {
				returnNode, subNode := n.Right.Remove(p, (axis+1)%n.Dimensions())
				if returnNode != nil {
					if returnNode == n.Right {
						n.Right = subNode
					}
					return returnNode, nil
				}
			}

			return nil, nil //不等，无需删除
		}
	}
	if n.Left != nil {
		biggest := n.Left.FindMax(axis, nil)
		removed, sub := n.Left.Remove(biggest, axis)
		removed.Left = n.Left
		removed.Right = n.Right
		if n.Left == removed {
			removed.Left = sub
		}
		return n, removed

	}
	if n.Right != nil {
		smallest := n.Right.FindMin(axis, nil)
		removed, sub := n.Right.Remove(smallest, axis)
		removed.Right = n.Right
		removed.Left = n.Left
		if n.Right == removed {
			removed.Right = sub
		}
		return n, removed
	}
	//left,right=nil,nil
	return n, nil
}

//按照维度搜索范围内的数据
func (n *node) RangeSearch(r KDrange.Range, axis int) []KDPoint {
	pointss := []KDPoint{} //节点集合
	for dim, limit := range r {
		if limit[0] > n.Dimension(dim) || limit[1] < n.Dimension(dim) { //节点在范围之内
			goto ChildCheck
		}
	}
	pointss = append(pointss, n.KDPoint)
ChildCheck:
	if n.Left != nil && n.Dimension(axis) >= r[axis][0] {
		pointss = append(pointss, n.Left.RangeSearch(r, (axis+1)%n.Dimensions())...)
	}
	if n.Right != nil && n.Dimension(axis) <= r[axis][0] {
		pointss = append(pointss, n.Right.RangeSearch(r, (axis+1)%n.Dimensions())...)
	}
	return pointss
}

//计算距离
func distance(p1, p2 KDPoint) float64 {
	sum := 0.0
	for i := 0; i < p1.Dimensions(); i++ {
		sum += math.Pow(p1.Dimension(i)-p2.Dimension(i), 2.0)
	}
	return math.Sqrt(sum)
}

func planeDistance(p KDPoint, plane float64, dim int) float64 {
	return math.Abs(plane - p.Dimension(dim)) //求距离的绝对值
}
func popLast(arr []*node) ([]*node, *node) {
	length := len(arr) - 1
	if length < 0 {
		return arr, nil
	}
	return arr[:length], arr[length]
}
func GetKthDistance(nearest *pq.PriorityQueue, i int) float64 {
	if nearest.Len() <= i {
		return math.MaxFloat64
	}
	_, prio := nearest.Get(i)
	return prio
}
