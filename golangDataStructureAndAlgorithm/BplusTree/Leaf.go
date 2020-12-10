package main

import "sort"

//继续存储数据
type kv struct {
	key   int    //数据
	value string //
}

type kvs [MaxKC]kv

//求数组的长度
func (kvx *kvs) Len() int {
	return len(kvx)
}

//交换数据
func (kvx *kvs) Swap(i, j int) {
	kvx[i], kvx[j] = kvx[j], kvx[i]
}

//判断大小
func (kvx *kvs) Less(i, j int) bool {
	return kvx[i].key < kvx[j].key
}

//判断大小
type leafNode struct {
	kvs    kvs //数据
	count  int
	next   *leafNode     //下一个节点
	parent *interiorNode //父亲节点
}

//创建叶子节点
func NewLeafNode(parent *interiorNode) *leafNode {
	return &leafNode{parent: parent}
}

//
func (l *leafNode) find(key int) (int, bool) {
	//myfunc是一个函数，主要进行数据对比
	myfunc := func(i int) bool {
		return l.kvs[i].key >= key
	}
	i := sort.Search(l.count, myfunc) //实现查询
	if i < l.count && l.kvs[i].key == key {
		return i, true
	}
	return i, false
}

//数组插入叶子节点

func (l *leafNode) insert(key int, value string) (int, *leafNode, bool) {
	i, ok := l.find(key)
	if ok {
		l.kvs[i].value = value
		return 0, nil, false
	}
	//判断叶子节点是否满了
	if !l.full() {
		copy(l.kvs[i+1:], l.kvs[i:l.count]) //数组的删除，需要整体往后移动(费劲)
		l.kvs[i].key = key
		l.kvs[i].value = value
		l.count++
		return 0, nil, false
	} else {
		next := l.split() //分裂叶子节点
		if key < next.kvs[0].key {
			l.insert(key, value)
		} else {
			next.insert(key, value)
		}
		return next.kvs[0].key, next, true
	}

}

func (l *leafNode) full() bool {
	return l.count == MaxKV //判断是否满了
}

func (l *leafNode) Parent() *interiorNode {
	return l.parent
}

func (l *leafNode) SetParent(p *interiorNode) {
	l.parent = p
}
func (l *leafNode) CountNum() int {
	return l.count
}

//初始化数组
func (l *leafNode) InitArray(num int) {
	for i := num; i < len(l.kvs); i++ {
		l.kvs[i] = kv{}
	}
}

//叶子节点，分裂//123    456   7
//123456 ---------99999
//123    456  (123指向456)
func (l *leafNode) split() *leafNode {
	next := NewLeafNode(nil)                //新建一个右边节点
	copy(next.kvs[0:], l.kvs[l.count/2+1:]) //复制数据到右边节点
	l.InitArray(l.count/2 + 1)              //hou
	next.count = MaxKV - l.count/2 - 1
	next.next = l.next
	l.count = l.count/2 + 1 //取得中间节点
	l.next = next
	return next
}
