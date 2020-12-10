package main

import (
	"errors"
	"fmt"
)

type HashNode struct {
	Key    int
	Value  int
	Childs map[int]*HashNode //哈希树不是二叉树
}

//插入数据
func (c *HashNode) AddValueRecursize(keys []int) {
	c.Value += 1 //
	if len(keys) == 0 {
		return
	}
	childNode, ok := c.Childs[keys[0]]
	if !ok {
		childNode := &HashNode{Key: keys[0]}
		if childNode == nil {
			c.Childs = make(map[int]*HashNode)

		}
		c.Childs[keys[0]] = childNode
	}
	if len(keys) > 1 {
		childNode.AddValueRecursize(keys[1:])
	} else if len(keys) == 1 {
		childNode.Value += 1
	}
}

//插入数据无需创建
func (c *HashNode) AddValueWithoutCreate(keys []int) error {
	c.Value += 1 //
	if len(keys) == 0 {
		return nil
	}
	childNode, ok := c.Childs[keys[0]]
	if !ok {
		return errors.New("no key for node")
	}
	c.Childs[keys[0]] = childNode

	if len(keys) > 1 {
		childNode.AddValueRecursize(keys[1:])
	} else if len(keys) == 1 {
		childNode.Value += 1
	}
	return nil
}

//插入数据无需创建
func (c *HashNode) AddNodeWithoutValue(keys []int) error {

	if len(keys) == 0 {
		return nil
	}
	childNode, ok := c.Childs[keys[0]]
	if !ok {

		childNode = &HashNode{Key: keys[0]}
		if childNode == nil {
			c.Childs = make(map[int]*HashNode)
		}
		c.Childs[keys[0]] = childNode
	}
	if len(keys) > 1 {
		err := childNode.AddNodeWithoutValue(keys)
		if err != nil {
			errors.New("创建失败")
		}
	} else if len(keys) == 1 {

	}
	return nil
}

//提取数据
func (c *HashNode) GetValueRecursive(keys []int) (int, error) {
	if len(keys) == 0 {
		return c.Value, nil
	}
	childNode, ok := c.Childs[keys[0]]
	if !ok {
		return 0, errors.New("节点不存在")
	} else {
		return childNode.GetValueRecursive(keys[1:])
	}
}
func main() {
	root := &HashNode{}
	root.AddValueRecursize([]int{0, 1, 2, 3, 5})
	fmt.Println(root.GetValueRecursive([]int{0, 1, 2, 3, 5}))
}
