package main

import "errors"

type Node struct {
	data  interface{}
	pNext *Node
}
type LinkStack interface {
	IsEmpty() bool         //判断为空
	Push(data interface{}) //入栈
	Pop() interface{}      //出栈
	Length() int           //长度
}
func NewStack()*Node{
return &Node{}
}
func (n *Node)IsEmpty() bool{
	if n.pNext==nil{
		return true
	}else{
		return false
	}
}         //判断为空
func (n *Node)Push(data interface{}){
newNode:=&Node{data:data}
newNode.pNext=n.pNext
n.pNext=newNode
} //入栈
func (n *Node)Pop() (interface{} ,error){
if n.IsEmpty(){
	return nil,errors.New("bug")
}
value:=n.pNext.data//要弹出得数据
n.pNext=n.pNext.pNext
return value,nil
}      //出栈
func (n *Node)Length() int{
pointNext:= n
	length:=0
	for pointNext.pNext!=nil{
		pointNext=pointNext.pNext
		length++
}
return length
}          //长度

