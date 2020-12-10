package main

import (
	"errors"
	"math"
)

type HashTable struct {
	Table map[int]*DoubleLinkList
	Size int
	Cap int
}
//元素
type Item struct {
	Key string
Value interface{}
}

func NewHashTable(cap int) *HashTable {
tab:=make(map[int]*DoubleLinkList,cap)
return &HashTable{tab,0,cap}
}
func (h *HashTable)Put(key ,value string)  {
index:=h.Pos(key)//索引
if h.Table[index]==nil{
	h.Table[index]=NewDoubleLinkList()
}
item:=&Item{key,value}//新建插入对象
data,err:=h.Find(index,key)
if err!=nil{
	h.Table[index].Append(item)
h.Size++
}else{
	data.Value=value//替换
}
}
func (h *HashTable)Del(key string)error  {
index:=h.Pos(key)
myList:=h.Table[index]
var val *Item
myList.Each(func(node DoubleLinkNode) {
	if node.Value.(Item).Key==key{
		val=node.Value.(*Item)//取出数据
	}
})
if val==nil{
	return nil//返回数据
}
h.Size--
retrurn myList.Remove(val)
}
//循环hash表的各个链表，循环每个链表的元素
func (h *HashTable)ForEach(item *Item)  {
for k:=range h.Table{
	if h.Table[k]!=nil{
		h.Table[k].Each(func(node DoubleLinkNode) {
			f(node.Value.(*Item))
		})
	}
}
}
func (h *HashTable)Pos(s string)int  {
return hashCode(s)%h.Cap//根据hash值计算
}
func (h *HashTable)Find(i  int,key string)(*Item,error)  {
myList:=h.Table[i]//每一个hash值对应一个链表
var val *Item
myList.Each(func(node DoubleLinkNode) {
	if node.Value.(*Item).Key==key{
		val=node.Value.(*Item)//取出数据
	}
})
if val==nil{
	return nil,errors.New("not find")
}
return val,nil
}
func (h *HashTable)Get(key string)  (interface{},error){
	index:=h.Pos(key)
	item,err:=h.Find(index,key)
	if err!=nil{
		return "",errors.New("not find" )
	}
	return item.Value,nil
}
//根据字符串计算hash
func hashCode(str string)int  {
hash:=int32(0)
for i:=0;i<len(str);i++{
	hash=int32(hash<<5-hash)+int32(str[i])
	hash&=hash
}
return int(math.Abs(float64(hash)))
}
func main()  {

}