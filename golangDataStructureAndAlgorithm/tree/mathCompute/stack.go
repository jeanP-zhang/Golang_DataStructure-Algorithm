package main

import  "container/list"

type Stack struct {
	lists *list.List
}

func NewStack()*Stack  {
	list:=list.New()//新建内存
	return &Stack{lists:list}//新建一个栈
}
func (stack *Stack)Push(value interface{})  {
	stack.lists.PushBack(value)
}
func (stack *Stack)Pop()interface{}  {
element:=stack.lists.Back()//取得最后一个数据
if element!=nil{
	stack.lists.Remove(element)
	return element.Value
}
return nil
}
//取得数据不删除
func (stack *Stack)Peak()interface{}  {
	element:=stack.lists.Back()//取得最后一个数据
	if element!=nil{
		return element.Value
	}
	return nil
}

func (stack *Stack)Len()int  {
return stack.lists.Len()
}
func (stack *Stack)IsEmpty()bool  {
	return stack.Len()==0
}