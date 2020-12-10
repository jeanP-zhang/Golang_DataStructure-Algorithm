package main

const (
	//叶子节点的最大存储数量 2^n-1
	MaxKV = 255
	//中间节点的最大存储数量
	MaxKC = 511
)

//接口设计
type node interface {
	find(key int) (int, bool)      //查找key
	Parent() *interiorNode         //返回父节点
	SetParent(node2 *interiorNode) //
	full() bool                    //p判断是否满了
	CountNum() int                 //统计元素数量
}
