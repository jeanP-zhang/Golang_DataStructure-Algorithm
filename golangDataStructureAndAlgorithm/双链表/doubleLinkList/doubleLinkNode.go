package doubleLinkList

type DoubleLinkNode struct {
	Value interface{}
	Prev  *DoubleLinkNode //上一个节点
	Next  *DoubleLinkNode //下一个节点
}

func NewDoubleLinkNode(value interface{}) *DoubleLinkNode {
	return &DoubleLinkNode{value, nil, nil}
}
func (node *DoubleLinkNode) Values() interface{} {
	return node.Value
}

//返回上一个节点
func (node *DoubleLinkNode) PrevNode() *DoubleLinkNode {
	return node.Prev
}

//返回下一个节点
func (node *DoubleLinkNode) NextNode() *DoubleLinkNode {
	return node.Next
}
