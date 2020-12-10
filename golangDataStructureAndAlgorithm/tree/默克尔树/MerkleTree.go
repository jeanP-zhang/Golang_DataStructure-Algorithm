package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type MerkleTree struct {
	RootNode *MerkleNode
}

//默克尔树的节点
type MerkleNode struct {
	left, right *MerkleNode
	HashData    []byte
}

//叶子节点，非叶子节点
func NewMerkleNode(left, right *MerkleNode, hashdata []byte) *MerkleNode {
	myNode := new(MerkleNode)        //创建结构体
	if left == nil && right == nil { //叶子节点
		myNode.HashData = hashdata //赋值数据
	} else {
		preHashes := append(left.HashData, right.HashData...) //叠加数据
		hashLR := sha256.Sum256(preHashes)                    //计算hash
		hash := sha256.Sum256(hashLR[:])                      //截取数据
		myNode.HashData = hash[:]                             //哈希数据
	}
	myNode.left = left
	myNode.right = right
	return myNode
}

//数组反转
func ReverseByte(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func NewMerkleTree(dataXx [][]byte) *MerkleTree {
	var nodes []MerkleNode //叶子集合
	for _, datax := range dataXx {
		node := NewMerkleNode(nil, nil, datax)
		nodes = append(nodes, *node) //加入节点集合
	}
	j := 0 //每一层的第一个元素
	//每次折半处理
	for length := len(dataXx); length > 1; length = (length + 1) / 2 {
		for i := 0; i < length; i += 2 {
			half := min(i+1, length-1)
			node := NewMerkleNode(&nodes[j+i], &nodes[j+half], nil) //生成哈希
			nodes = append(nodes, *node)
		}
		j += length
	}
	myTree := &MerkleTree{&nodes[len(nodes)-1]} //最后一个节点，根节点
	return myTree
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func main() {
	data1, _ := hex.DecodeString("4b901a321420")
	fmt.Printf("%T", data1)
	data := [][]byte{}
	myroot := NewMerkleTree(data)
}
