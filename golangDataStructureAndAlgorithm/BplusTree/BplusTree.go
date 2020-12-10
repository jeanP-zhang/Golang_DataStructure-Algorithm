package main

import "fmt"

type BPlusTree map[int]node //（内存中的b+树）定义，通过map来存储，一个数据对应一个节点

//建立一个新树
func NewBplusTree() *BPlusTree {
	bt := BPlusTree{}               //初始化
	leaf := NewLeafNode(nil)        //叶子节点
	r := NewInteriorNode(nil, leaf) //中间节点
	leaf.parent = r                 //设定父亲节点
	//设置根节点
	bt[-1] = r   //被当作根节点
	bt[0] = leaf //一个叶子
	return &bt
}

//返回根节点
func (bpt *BPlusTree) Root() node {
	return (*bpt)[-1]
}

//处理第一个节点
func (bpt *BPlusTree) First() node {
	return (*bpt)[0]
}

//统计数量
func (bpt *BPlusTree) Count() int {
	count := 0
	leaf := (*bpt)[0].(*leafNode)
	for {
		count += leaf.CountNum() //数量的叠加
		if leaf.next == nil {
			break
		}
		leaf = leaf.next
	}
	return count
}
func (bpt *BPlusTree) Values() []*leafNode {
	nodes := make([]*leafNode, 0) //开辟节点
	leaf := (*bpt)[0].(*leafNode) //
	for {
		nodes = append(nodes, leaf) //数据节点的叠加
		if leaf.next == nil {
			break
		}
		leaf = leaf.next
	}
	return nodes
}
func (bpt *BPlusTree) Insert(key int, value string) {
	//插入前搜索下是否存在,并确定插入位置
	_, oldIndex, leafs := search((*bpt)[-1], key)
	p := leafs.Parent()                             //保存父亲节点
	mid, nextLeaf, bump := leafs.insert(key, value) //插入叶子节点判断是否分裂
	if !bump {                                      //没有分裂，直接返回
		return
	} else { //分裂的节点插入B+树
		(*bpt)[mid] = nextLeaf
		var midNode node
		midNode = leafs
		p.kcs[oldIndex].child = leafs.next   //设置父亲节点
		leafs.next.SetParent(p)              //分裂的节点设置父亲节点
		interior, interiorP := p, p.Parent() //获取中间节点，父亲节点
		//平衡过程，迭代向上判断是否需要平衡

		for {
			//	var oldIndex int //保存老的索引
			var newInterior *interiorNode
			//判断是否到达根节点
			isRoot := interiorP == nil
			if !isRoot {
				oldIndex, _ = interiorP.find(key) //查找
			}
			//叶子节点分裂后的中间节点传入父亲的中间节点，传入分裂
			mid, newInterior, bump = interior.insert(mid, midNode)
			if !bump {
				return
			}
			(*bpt)[newInterior.kcs[0].key] = newInterior //插入填充好了的map
			if !isRoot {
				interiorP.kcs[oldIndex].child = newInterior //没有到根节点，直接插入到父亲节点
				newInterior.SetParent(interiorP)
				midNode = interior
			} else {
				//更新节点
				(*bpt)[interior.kcs[0].key] = (*bpt)[-1]       //备份根节点
				(*bpt)[-1] = NewInteriorNode(nil, newInterior) //根节点插入新的根节点
				node := (*bpt)[-1].(*interiorNode)
				node.insert(mid, interior) //重新插入
				(*bpt)[-1] = node
				newInterior.SetParent(node) //设定父亲节点
			}
			interior, interiorP = interiorP, interior.Parent()

		}
	}

}

func (bpt *BPlusTree) Search(key int) (string, bool) {
	kvss, _, _ := search((*bpt)[-1], key)
	if kvss == nil {
		return " ", false
	} else {
		return kvss.value, true
	}
}

//搜索数据
func search(n node, key int) (*kv, int, *leafNode) {
	curr := n ///查找
	oldIndex := -1
	for {
		switch t := curr.(type) {
		case *leafNode:
			i, ok := t.find(key)
			if !ok {
				return nil, oldIndex, t //没有找到
			} else {
				return &t.kvs[i], oldIndex, t
			}
		case *interiorNode:
			i, _ := t.find(key) //中间节点进行查找
			curr = t.kcs[i].child
			oldIndex = i
		default:
			panic("异常节点")
		}
	}
}

//item 项目
func main() {
	bpt := NewBplusTree()
	for i := 0; i < 100; i++ {
		bpt.Insert(i, " ")
	}
	fmt.Println(bpt.Count())
	fmt.Println(bpt.Search(34))
}
