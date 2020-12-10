package SingleLink

//单向链表
type SingleLinkNode struct {
	Value interface{}
	Next  *SingleLinkNode
}

//单链表的接口

func NewSingleLinkNode(data interface{}) *SingleLinkNode {
	return &SingleLinkNode{data, nil}
}

//返回数据
func (node *SingleLinkNode) Values() interface{} {
	return node.Value
}
func (node *SingleLinkNode) PNext() *SingleLinkNode {
	return node.Next
}
