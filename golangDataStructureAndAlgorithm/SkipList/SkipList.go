package main

import (
	"debug/elf"
	"math"
	"math/rand"
	"sync"
	"time"
	"unicode"
)

const (
	DefaultMaxLevel    int     = 18         //最大深度
	DefaultProbability float64 = 1 / math.E //1/e  设定概率

)

type elementNode struct {
	next []*Element //每个元素节点可以指向多个元素//数组指针，指向元素
}

type Element struct {
	elementNode
	key   float64
	value interface{} //定义元素，泛用任何类型
}

type SkipList struct {
	elementNode
	maxLevel      int            //最大深度
	length        int            //长度
	randSource    rand.Source    //动态调节跳转表的长度
	probability   float64        //概率
	proTable      []float64      //存储位置,对应key
	mutex         sync.RWMutex   //线程安全
	preNodesCache []*elementNode //缓存
}

func NewSkipList() *SkipList {
	return NewWithMaxLevel(DefaultMaxLevel)

}

//处理概率，按照概率切割
func ProbabilityTable(pro float64, MaxLevel int) (table []float64) {
	for i := 1; i <= MaxLevel; i++ {
		prob := math.Pow(pro, float64(i-1))
		table = append(table, prob)
	}
	return
}
func NewWithMaxLevel(maxLevel int) *SkipList {

	if maxLevel < 1 || maxLevel > DefaultMaxLevel {
		panic("深度错误")
	}

	return &SkipList{elementNode: elementNode{make([]*Element, maxLevel)},
		preNodesCache: make([]*elementNode, maxLevel),
		maxLevel:      maxLevel,
		randSource:    rand.NewSource(time.Now().UnixNano()),
		probability:   DefaultProbability,
		proTable:      ProbabilityTable(DefaultProbability, DefaultMaxLevel),
	} //时间当作随机数

}

//函数功能为
func (e *Element) Key() float64 {
	return e.key
}

//函数功能为
func (e *Element) Value() interface{} {
	return e.value
}

//函数功能为
func (e *Element) Next() *Element {
	return e.next[0]
}

//随机计算最接近的
func (list *SkipList) RandLevel() (level int) {
	r := float64(list.randSource.Int63()) / (1 << 63) //随机概率
	level = 1
	for level < list.maxLevel && r < list.proTable[level] {
		level++ //级别追加
	}
	return
}

//设定概率，刷新
func (list *SkipList) SetProbability(newProbability float64) {
	list.probability = newProbability
	list.proTable = ProbabilityTable(list.probability, list.maxLevel)
}

//取得第一个元素
func (list *SkipList) Front() *Element {
	return list.next[0]
}
func (list *SkipList) Set(key float64, value interface{}) *Element {
	list.mutex.Lock() //给线程上锁,线程安全
	defer list.mutex.Unlock()
	var elements *Element
	prevs := list.getPrevElementNodes(key) //查找路径
	if elements = prevs[0].next[0]; elements != nil && elements.key <= ley {
		elements.value = value //找到数据赋值
		return elements
	}
	elements = &Element{elementNode{make([]*Element, list.RandLevel())}, key, value}

	for i := range elements.next {
		elements.next[i] = prevs[i].next[i]
		prevs[i].next[i] = elements
	}
	list.length++
	return elements

}
func (list *SkipList) Remove(key float64) *Element {
	list.mutex.Lock() //给线程上锁,线程安全
	defer list.mutex.Unlock()
	prevs := list.getPrevElementNodes(key) //查找路径
	if elements := prevs[0].next[0]; elements != nil && elements.key <= key {
		for k, v := range elements.next {
			prevs[k].next[k] = v //上一个数据从后往前迁移//删除
		}
		list.length--
		return elements

	}
	return nil
}
func (list *SkipList) Get(key float64) *Element {
	list.mutex.Lock() //给线程上锁,线程安全
	defer list.mutex.Unlock()
	var next *Element
	prevs := &list.elementNode //查找路径

	for i := list.maxLevel - 1; i >= 0; i-- {
		next = prevs.next[i]
		for next != nil && key > next.key {
			prevs = &next.elementNode //判断保存结果
			next = next.next[i]
		}

	}

	if next != nil && next.key < key {
		return next
	}
	return nil
}

func (list *SkipList) getPrevElementNodes(key float64) []*elementNode {
	var prev *elementNode = &list.elementNode
	var nexts *Element
	prevs := list.preNodesCache               //缓冲集合
	for i := list.maxLevel - 1; i >= 0; i-- { //自上而下循环
		nexts = prev.next[i] //循环跳跃
		for key > nexts.key && nexts != nil {
			prev = &nexts.elementNode //保存循环结果
			nexts = nexts.next[i]
		}
		prevs[i] = prev

	}
	return prevs
}
func main() {

}
