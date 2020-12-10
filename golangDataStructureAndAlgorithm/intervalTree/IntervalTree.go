package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Start int64
	End int64
	Index int
	Max int64
}
type Nodes []Node
func (ns *Nodes)Add(start ,end int64)int  {
	idx:=len(*ns)
	*ns=append(*ns,Node{Start:start,End:end,Index:idx})
return idx
}
func max(a,b int64)int64  {
	if a<b{
		return b
	}
	return a
}
//实现排序
func (ns Nodes)sort () {
var a= func(i,j int)bool {
	if i>j{
		return true
	}else{
		return false
	}
}
	sort.Slice(ns, a)
}

func (ns Nodes)FillMax(off int)int64  {
length:=len(ns)
mid:=length/2
v:=(ns)[mid].End
if mid>0{
	v=max(v,(ns)[:mid].FillMax(off))
}
	if mid<length-1{
		v=max(v,(ns)[mid+1:].FillMax(off))
	}
	(ns)[mid].Max=v
	return v
}

func (ns *Nodes)Build()  {
	ns.sort()
	ns.FillMax(0)
}
func (ns *Nodes)query(q int64) []int{
return ns.Query(q,nil)
}
func (ns Nodes)Query(q int64,res []int)[]int  {
	length:=len(ns)
	mid:=length/2
	if q>ns[mid].Max{
		return res
	}
	if q>=ns[mid].Start&&q<=ns[mid].End{
		res=append(res ,ns[mid].Index)//叠加
	}
	if mid>0{
		res=ns[:mid].Query(q,res)
	}
	if mid<length-1{
		res=ns[mid+1:].Query(q,res)
	}
	return res
}


func (ns Nodes)QueryNode(q int64,res []Node)[]Node  {
	length:=len(ns)
	mid:=length/2
	if q>ns[mid].Max{
		return res
	}
	if q>=ns[mid].Start&&q<=ns[mid].End{
		res=append(res ,ns[mid])//叠加
	}
	if mid>0{
		res=ns[:mid].QueryNode(q,res)
	}
	if mid<length-1{
		res=ns[mid+1:].QueryNode(q,res)
	}
	return res
}
func main()  {

	ns:=new(Nodes)
	ns.Add(8,9)
	ns.Add(16,21)
	ns.Add(25,30)
	ns.Add(5,8)
	ns.Add(15,23)
	ns.Add(17,19)
ns.Build()
for i,x:=range *ns{
	fmt.Printf("%2d %2d  %2d %2d %2d  \n",i,x.Start,x.Index,x.End,x.Max)
}
r:=(*ns).query(20)
fmt.Println("r:",r)
}