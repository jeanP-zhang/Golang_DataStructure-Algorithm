package main
type LinkQueue interface {
	IsEmpty() bool         //判断为空
	EnQueue(value interface{}) //入队
	DeQueue() interface{}      //出队
	Length() int           //长度
}
type QueueLink struct {
	rear *Node
	front*Node
}
func NewLinkQueue()*QueueLink{
	return &QueueLink{}
}
func (qlk *QueueLink)length()int{

}
func (qlk *QueueLink)EnQueue(value interface{}){
newnode:=&Node{
	data:value,
	pNext:nil,
}//新的节点
}
func (qlk *QueueLink)DeQueue() interface{}{

}
func (qlk *QueueLink)IsEmpty() bool{

}