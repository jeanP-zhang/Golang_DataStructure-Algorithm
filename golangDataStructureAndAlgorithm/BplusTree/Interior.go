package main

import "sort"

//存储数据
type kc struct {
	key   int  //数据
	child node //接口类型
}

type kcs [MaxKC + 1]kc //数组，存储一个数组

//中间节点数据结构
type interiorNode struct {
	kcs    kcs           //数组存储数据
	count  int           //存储元素的数量
	parent *interiorNode //指向父亲节点
}

//若要创建一个节点，就要需要其父节点
func NewInteriorNode(p *interiorNode, largestChild node) *interiorNode {
	in := &interiorNode{parent: p, count: 1}
	if largestChild != nil {
		in.kcs[0].child = largestChild
	}
	return in
}

func (in *interiorNode) find(key int) (int, bool) { //一定会给予个位置
	//myfunc是一个函数，主要进行数据对比
	myfunc := func(i int) bool {
		return in.kcs[i].key > key
	}
	i := sort.Search(in.count-1, myfunc) //实现查询，跳过第一个节点
	return i, true
}

//判断长度
func (kvx *kcs) Len() int {
	return len(kvx)
}

//交换数据
func (kvx *kcs) Swap(i, j int) {
	kvx[i], kvx[j] = kvx[j], kvx[i]
}

//判断大小
func (kvx *kcs) Less(i, j int) bool {
	if kvx[i].key == 0 { //中间节点的，数组第一个空着
		return false
	}
	if kvx[j].key == 0 {
		return true
	}
	return kvx[i].key < kvx[j].key
}
func (in *interiorNode) full() bool {
	return in.count == MaxKC //判断是否满了
}

func (in *interiorNode) Parent() *interiorNode {
	return in.parent
}

func (in *interiorNode) SetParent(p *interiorNode) {
	in.parent = p
}
func (in *interiorNode) CountNum() int {
	return in.count
}

//初始化数组
func (l *interiorNode) InitArray(num int) {
	for i := num; i < len(l.kcs); i++ {
		l.kcs[i] = kc{}
	}
}

//插入节点
func (in *interiorNode) insert(key int, child node) (int, *interiorNode, bool) {
	//确定位置
	i, _ := in.find(key)
	if !in.full() { //b+树中间节没有满
		copy(in.kcs[i+1:], in.kcs[i:in.count]) //整体往后移动一位
		//设置子节点分裂以后的元素，设置key
		in.kcs[i].key = key
		in.kcs[i].child = child //设定
		child.SetParent(in)     //设定父亲节点
		in.count++
		return 0, nil, false
	} else { //b+树中间节点满了
		in.kcs[MaxKC].key = key //在最后的空间进行交换
		in.kcs[MaxKC].child = child
		child.SetParent(in)        //设定父亲节点
		next, midKey := in.split() //切割
		return midKey, next, true  //返回终点

	}
}

//12345 _____   -> 123 _45_
func (in *interiorNode) split() (*interiorNode, int) {
	//节点分裂前，节点插入正确的位置
	sort.Sort(&in.kcs) //确保有序
	//取得中间元素
	midIndex := MaxKC / 2
	midChild := in.kcs[midIndex].child //取得中间节点
	midKey := in.kcs[midIndex].key     //取得键值
	//新建一个中间节点
	next := NewInteriorNode(nil, nil)
	copy(next.kcs[0:], in.kcs[midIndex+1:]) //拷贝数据
	in.InitArray(midIndex + 1)              //数据的初始化
	next.count = MaxKC - midIndex           //下一个节点的数量
	//新开辟节点的每个叶子节点的祖先设置位next
	for i := 0; i < next.count; i++ {
		next.kcs[i].child.SetParent(next)
	}
	in.count = midIndex + 1
	in.kcs[in.count-1].key = 0 //设置为0,预留一个
	in.kcs[in.count-1].child = midChild
	midChild.SetParent(in) //设置父亲节点
	return next, midKey
}
