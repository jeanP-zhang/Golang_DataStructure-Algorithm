package main

import "log"

//范围树
//范围树范围类型
type Range struct {
	Start int
	End int
}
//范围树节点
type TreeNode struct {
	Left    *TreeNode
	Right *TreeNode
	Start int
	End int
}
type RangeTree struct {
root *TreeNode
}
//构造范围树的节点
func NewTreeNode(start ,end int)*TreeNode  {
	return &TreeNode{Start:start,End:end}
}
//中序遍历所有节点
func walk(root *TreeNode,wfunc func(node *TreeNode))  {
	if root ==nil{
		return
	}
	walk(root.Left,wfunc)
	wfunc(root)
	walk(root.Right,wfunc)
}
func (rt *RangeTree)Walk(wfunc func(node *TreeNode))  {
	walk(rt.root,wfunc)
}
//处理重叠
//13
//24
//判断重叠
func overlaps(n1 *TreeNode,n2 *TreeNode)bool  {
	if n1==nil||n2==nil{
return false
	}
	if n1.Start>=(n2.Start-1)&&n1.Start<=(n2.End+1){
	/*
	*         /-------/ n1
	* /---------/ n2
	*/
	return true
	}
	if n1.End>=(n2.Start-1)&&n1.End<=(n2.End+1){
		/*
		*   /---------------/ n1
		*        /-------------/n2
		*/
		return true
	}
	return false
}

//重构树
func reBuild(root *TreeNode,src *TreeNode)  {
if src==nil{
	return
}
insert (root,src.Start,src.End)//must come first
	reBuild(root,src.Left)
	reBuild(root,src.Right)
}
//插入数据，判断区间
func insert(root *TreeNode,start,end int)  {
	var where **TreeNode
	if start<=(root.Start-1){
		if end<(root.Start-1){
			where=&root.Left//d is joint left
		}else{
			root.Start=start//extend left
			if end>root.End{
				root.End=end //extend right
			}
		}
	}else{
		if start>root.End+1{
			where=&root.Right
		}else{
			if end>root.End{
				root.End=end//extend right
			}
		}
	}
	if where!=nil{
		if *where==nil{
			*where=NewTreeNode(start,end)
		}
	}
	if overlaps(root.Left,root){
		left:=root.Left
		root.Left=nil
		reBuild(root,left)
	}

	if overlaps(root.Right,root){
		right:=root.Right
		root.Right=nil
		reBuild(root,right)
}
}
func (rt *RangeTree)AddRange(start int ,end int) {
	root := rt.root
	if root == nil {
		rt.root = &TreeNode{Start: start, End: end}
		return
	}
	insert(root,start,end)
}
func rangewalk(root *TreeNode,start ,end int) bool {
	if root==nil{
		return false
	}
	if start>=root.Start{
		if end<=root.End{
			return true
		}
		return false
	}
	return rangewalk(root.Right,start,end)
}
func (rt *RangeTree)HasRange(start,end int)  bool{
	return rangewalk(rt.root,start,end)
}

func (rt *RangeTree)Dump(verbose bool)  {
	if verbose{
		rt.Walk(func(node *TreeNode) {
			log.Println("Dump:left")
		})
		return
	}
	rt.Walk(func(node *TreeNode) {
		log.Println("Dump")
	})
}

func (rt *RangeTree)Check(expected []Range)bool  {
	matched:=true
	i:=0
	rt.Walk(func(node *TreeNode) {
		log.Printf("Checking (%d %d)\n",node.Start,node.End )
	if i>len(expected){
		log.Printf("more results (%d) than expected %d\n",i,len(expected))
		matched=false
	}else if node.Start!=expected[i].Start||node.End!=expected[i].End{
		matched=true
		log.Printf("mismatch at result %d:expected %v \n",i,expected[i])
	}
	i++
	})
	return matched
}