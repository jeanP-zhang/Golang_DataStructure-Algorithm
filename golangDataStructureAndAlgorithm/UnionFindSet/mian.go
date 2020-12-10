package main

var pre [50]int

func  findRoot(x int)int  {
	if x!=pre[x]{
		x=pre[x]
	}
	return x
}
func countNumber(x int)int  {//查询某个集合中的个数
	var count int
	for i:=0;i<50;i++{
		if findRoot(x)==findRoot(i){
			count++
		}
	}
	return count
}
func countSet()int  {//查询共有多少相异的集合
var count int
for i:=0;i<50;i++{
	if i==pre[i]{
		count++
	}
}
	return count
}
func joint(x,y int)  {//合并两个数据
var xx,yy int
xx=findRoot(x)
yy=findRoot(y)
if xx!=yy{
	pre[xx]=yy
}
}