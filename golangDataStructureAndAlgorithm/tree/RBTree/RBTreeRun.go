package main

//定义常量红黑
const (
	RED   = true
	BLACK = false
)

//红黑树结构
type RBNode struct {
	Left   *RBNode //左节点
	Right  *RBNode //右节点
	Parent *RBNode //父亲节点
	Color  bool    //颜色
	//DataItem interface{}//数据
	Item //数据接口
}
type Item interface {
	Less(than Item) bool
}

type RBTree struct {
	NIL   *RBNode
	Root  *RBNode
	count uint
}

//比大小
func less(x, y Item) bool {
	return x.Less(y)
}

//初始化内存
func NewRBTree() *RBTree {
	return new(RBTree).Init()
}

//初始化红黑树
func (rbt *RBTree) Init() *RBTree {
	node := &RBNode{nil, nil, nil, BLACK, nil}
	return &RBTree{node, node, 0}
}

//获取红黑树长度
func (rbt *RBTree) Len() uint {
	return rbt.count
}

//取得红黑树的极大值
func (rbt *RBTree) GetMax() {

}

func (rbt *RBTree) max(x *RBNode) *RBNode {
	if x == rbt.NIL {
		return rbt.NIL
	}
	for x.Right != rbt.NIL {
		x = x.Right
	}
	return x
}

//取得红黑树的极小值

func (rbt *RBTree) GetMin() {

}
func (rbt *RBTree) min(x *RBNode) *RBNode {
	if x == rbt.NIL {
		return rbt.NIL
	}
	for x.Left != rbt.NIL {
		x = x.Left
	}
	return x
}

//插入一条数据
func (rbt *RBTree) Insert(item Item) *RBNode {
	if item == nil {
		return nil
	}
	return rbt.insert(&RBNode{rbt.NIL, rbt.NIL, rbt.NIL, RED, item})
}

//搜索红黑树
func (rbt *RBTree) search(x *RBNode) *RBNode {
	pNode := rbt.Root //根节点
	for pNode != rbt.NIL {
		if less(pNode.Item, x.Item) {
			pNode = pNode.Right
		} else if less(x.Item, pNode.Item) {
			pNode = pNode.Left
		} else {
			break //找到
		}
	}
	return pNode
}
func (rbt *RBTree) leftRotate(x *RBNode) {
	if x.Right == rbt.NIL {
		return //左旋转，逆时针，右孩子不可为0
	}
	y := x.Right
	x.Right = y.Left
	if y.Left != rbt.NIL {
		y.Left.Parent = x //设定父亲节点
	}
	y.Parent = x.Parent //交换父节点
	if x.Parent == rbt.NIL {
		//根节点
		rbt.Root = y
	} else if x == x.Parent.Left { //x在根节点的左边
		x.Parent.Left = y
	} else { //x在根节点的右边
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}
func (rbt *RBTree) rightRotate(x *RBNode) {
	if x.Left == nil {
		return //右边旋转,所以左子树不能为空
	}
	y := x.Left
	x.Left = y.Right
	if y.Right != rbt.NIL {
		y.Right.Parent = x //设置祖先
	}
	y.Parent = x.Parent //y保存x的父节点
	if x.Parent == rbt.NIL {
		rbt.Root = y
	} else if x == x.Parent.Left { //x小于根节点
		x.Parent.Left = y //父亲节点的孩子是x，父亲节点孩子y
	} else { //x大于根节点
		x.Parent.Right = y
	}
	y.Right = x
	x.Parent = y
}

//插入
func (rbt *RBTree) insert(z *RBNode) *RBNode {
	//寻找插入位置
	x := rbt.Root
	y := rbt.NIL
	for x != rbt.NIL {
		y = x                     //备份位置,数据插入下x，y之间
		if less(z.Item, x.Item) { //小于
			x = x.Left
		} else if less(x.Item, z.Item) { //大于
			x = x.Right
		} else { //相等
			return x //数据已经存在无法插入
		}
	}
	z.Parent = y
	if y == rbt.NIL {
		rbt.Root = z
	} else if less(z.Item, y.Item) {
		y.Left = z //小于，左边插入
	} else {
		y.Right = z //大于，右边插入
	}
	rbt.count++
	rbt.insertFixUp(z) //调整平衡
	return z
}

//插入之后调整平衡
func (rbt *RBTree) insertFixUp(z *RBNode) {
	for z.Parent.Color == RED { //一直循环下去,直到根节点
		if z.Parent == z.Parent.Parent.Left { //父亲节点在爷爷节点左边
			y := z.Parent.Parent.Right
			if y.Color == RED { //判断大伯节点红色，黑色
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent //循环 前进
			} else {
				if z == z.Parent.Right { //z比父亲小
					z = z.Parent
					rbt.leftRotate(z) //实现左旋
				} else { //z比父亲小
					z.Parent.Color = BLACK
					z.Parent.Parent.Color = RED
					rbt.rightRotate(z.Parent.Parent)
				}
			}
		} else { //父亲节点在爷爷节点右边
			y := z.Parent.Parent.Left
			if y.Color == RED { //判断大伯节点红色，黑色
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Left { //在左边
					z = z.Parent
					rbt.rightRotate(z)
				} else {
					z.Parent.Color = BLACK
					z.Parent.Parent.Color = RED
					rbt.leftRotate(z.Parent.Parent)
				}
			}
		}
	}
	rbt.Root.Color = BLACK
}

//函数初始化，实现代码的包含
func (rbt *RBTree) GetDepth() int { //函数
	var getDepth func(node *RBNode) int
	getDepth = func(node *RBNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			return 1
		}
		var leftDepth = getDepth(node.Left)
		var rightDepth = getDepth(node.Right)
		if leftDepth > rightDepth {
			return leftDepth + 1
		} else {
			return rightDepth + 1
		}

	}
	return getDepth(rbt.Root)
}

//近似查找
func (rbt *RBTree) searchle(x *RBNode) *RBNode {
	p := rbt.Root
	n := p
	for n != rbt.NIL {
		if less(n.Item, x.Item) {
			p = n
			n = n.Right
		} else if less(x.Item, n.Item) {
			p = n
			n = n.Left //小于
		} else {
			return n
			break //跳出循环
		}

	}
	if less(p.Item, x.Item) {
		return p
	}
	p = rbt.desuccessor(p) //近似查找
	return p
}

func (rbt *RBTree) successor(x *RBNode) *RBNode {
	if x == rbt.NIL {
		return rbt.NIL
	}
	if x.Right != rbt.NIL {
		return rbt.min(x.Right) //求左边的最大值
	}
	y := x.Parent
	for y != rbt.NIL && x == y.Right {
		x = y
		y = y.Parent
	}
	return y
}

func (rbt *RBTree) desuccessor(x *RBNode) *RBNode {
	if x == rbt.NIL {
		return rbt.NIL
	}
	if x.Left != rbt.NIL {
		return rbt.max(x.Left) //求左边的最大值
	}
	y := x.Parent
	for y != rbt.NIL && x == y.Left {
		x = y
		y = y.Parent
	}
	return y
}
func (rbt *RBTree) Delete(item Item) Item {
	if item == nil {
		return nil
	}
	return rbt.delete(&RBNode{rbt.NIL, rbt.NIL, rbt.NIL, RED, item}).Item
}
func (rbt *RBTree) delete(key *RBNode) *RBNode {
	z := rbt.search(key)
	if z == rbt.NIL {
		return rbt.NIL //无需删除
	}
	//新建节点下备份，夹逼
	var x, y *RBNode
	//节点
	ret := &RBNode{rbt.NIL, rbt.NIL, rbt.NIL, z.Color, z.Item} //新建节点，起到备份作用
	if z.Left == rbt.NIL || z.Right == rbt.NIL {
		y = z //单节点，y,z重合
	} else {
		y = rbt.successor(z) //找到最接近的右边最小
	}
	if y.Left != rbt.NIL {
		x = y.Left
	} else {
		x = y.Right
	}
	x.Parent = y.Parent
	if y.Parent == rbt.NIL {
		rbt.Root = x
	} else if y == y.Parent.Left {
		y.Parent.Left = x
	} else {
		y.Parent.Right = x
	}
	if y != z {
		z.Item = y.Item
	}
	if y.Color == BLACK {
		rbt.deleteFixUp(x)
	}
	rbt.count--
	return ret
}
func (rbt *RBTree) deleteFixUp(x *RBNode) {
	for x != rbt.Root && x.Color == BLACK {
		if x == x.Parent.Left { //x在左边
			w := x.Parent.Right //哥哥节点
			if w.Color == RED { //左边旋转
				w.Color = BLACK
				x.Parent.Color = RED
				rbt.leftRotate(x.Parent)
				w = x.Parent.Right //循环步骤
			}
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color = RED
				x = x.Parent //循环条件
			} else {
				if w.Right.Color == BLACK {
					w.Left.Color = BLACK
					w.Color = RED
					rbt.rightRotate(w) //右旋转
					w = x.Parent.Right //循环条件
				}
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Right.Color = BLACK
				rbt.leftRotate(x.Parent)
				x = rbt.Root
			}
		} else { //x在右边
			w := x.Parent.Left  //左边节点
			if w.Color == RED { //左旋
				w.Color = BLACK
				x.Parent.Color = RED
				rbt.rightRotate(x.Parent)
				w = x.Parent.Right //循环步骤
			}
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color = RED
				x = x.Parent
			} else {
				if w.Right.Color == BLACK {
					w.Left.Color = BLACK
					w.Color = RED
					rbt.leftRotate(w) //右旋转
					w = x.Parent.Left //循环条件
				}
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Right.Color = BLACK
				rbt.rightRotate(x.Parent)
				x = rbt.Root
			}
		}
	}
	x.Color = BLACK //循环到最后就是根节点
}
