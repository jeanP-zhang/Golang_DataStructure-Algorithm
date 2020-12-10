package main

import (
	"container/heap"
	"fmt"
)

//哈夫曼树，最优二叉树

//堆：快速取出极大值
type HuffmanTree interface {
	Freq() int //哈夫曼树的接口
}

//哈夫曼叶子类型
type HuffmanLeaf struct {
	freq  int  //频率
	value rune //int32
}

//哈夫曼树的类型
type HuffmanNode struct {
	freq        int //频率
	left, right HuffmanTree
}

func (self HuffmanLeaf) Freq() int { //频率
	return self.freq
}
func (self HuffmanNode) Freq() int {
	return self.freq
}

type treeHeap []HuffmanTree

func (th treeHeap) Len() int {
	return len(th)
}

//比较函数
func (th treeHeap) Less(i int, j int) bool {
	return th[i].Freq() < th[j].Freq()
}

//压入
func (th *treeHeap) Push(ele interface{}) {
	*th = append(*th, ele.(HuffmanTree))
}

//弹出
func (th *treeHeap) Pop() interface{} {
	po := (*th)[len(*th)-1]
	*th = (*th)[:len(*th)-1]
	return po
}
func (th treeHeap) Swap(i, j int) {
	th[i], th[j] = th[j], th[i]
}
func BuildTree(symFreq map[rune]int) HuffmanTree {
	var trees treeHeap
	for c, f := range symFreq {
		trees = append(trees, HuffmanLeaf{f, c}) //叠加数据
	}
	heap.Init(&trees) //开始使用堆
	for trees.Len() > 1 {
		a := heap.Pop(&trees).(HuffmanTree)
		b := heap.Pop(&trees).(HuffmanTree)
		heap.Push(&trees, HuffmanNode{a.Freq() + b.Freq(), a, b})
	} //构造哈夫曼树
	return heap.Pop(&trees).(HuffmanTree)
}
func showTimes(tree HuffmanTree, prefix []byte) {
	switch i := tree.(type) {
	case HuffmanLeaf:
		fmt.Printf("%c\t%d\n", i.value, i.freq) //打印数据与频率
	case HuffmanNode:
		prefix = append(prefix, '0')
		showTimes(i.left, prefix)       //递归到左子树
		prefix = prefix[:len(prefix)-1] //删除最后一个
		prefix = append(prefix, '1')
		showTimes(i.right, prefix)
	}
}

func showCodes(tree HuffmanTree, prefix []byte) {
	switch i := tree.(type) {
	case HuffmanLeaf:
		fmt.Printf("%s\n", string(prefix)) //打印数据与频率
	case HuffmanNode:
		prefix = append(prefix, '0')
		showCodes(i.left, prefix)       //递归到左子树
		prefix = prefix[:len(prefix)-1] //删除最后一个
		prefix = append(prefix, '1')
		showCodes(i.right, prefix)
	}
}
func main() {
	stringcode := "aaaaaaabbbffffeeeesssggaattt"
	fmt.Println("stringcode", stringcode)
	symFreqs := make(map[rune]int)
	for _, c := range stringcode {
		symFreqs[c]++ //统计频率
	}
	trees := BuildTree(symFreqs)
	showTimes(trees, []byte{})
}
