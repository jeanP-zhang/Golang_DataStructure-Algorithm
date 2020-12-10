package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

type rIndex []uint32
type ring struct {
	rMap map[uint32]string//环形结构体
	rIndexArr rIndex//索引
lock	 *sync.RWMutex//线程安全
}

func (r *rIndex)Less(i,j int)bool  {
	return (*r)[i]<(*r)[j]
}
func (r *rIndex)Len()int  {
	return len(*r)
}
func (r *rIndex)Swap(i,j int)  {
(*r)[i],(*r)[j]=(*r)[j],(*r)[i]
}
func (ri *ring)AddNode(nodeName string)  {
ri.lock.Lock()
defer ri.lock.Unlock()
index:=crc32.ChecksumIEEE([]byte(nodeName))//sha256,一个字节集合返回一个byte
if _,ok:=ri.rMap[index];ok{
	return//已经存在，返回
}
ri.rMap[index]=nodeName//完成赋值
ri.rIndexArr=append(ri.rIndexArr,index)//加载索引
sort.Sort(&ri.rIndexArr)

}

func (ri *ring)RemoveNode(nodeName string)  {
	ri.lock.Lock()
	defer ri.lock.Unlock()
	index:=crc32.ChecksumIEEE([]byte(nodeName))//sha256,一个字节集合返回一个byte
	if _,ok:=ri.rMap[index];!ok{
		return//已经存在，返回
	}
	delete(ri.rMap,index)//删除map内置结构
	ri.rIndexArr=rIndex{}
	for k:=range ri.rMap{
ri.rIndexArr=append(ri.rIndexArr,k)
	}
}
func (ri *ring)GetNode(nodeName string)  string{
	ri.lock.RLock()//其他线程可以读取，不可修改
	defer ri.lock.RUnlock()
	hash:=crc32.ChecksumIEEE([]byte(nodeName))
	i:=sort.Search(len(ri.rIndexArr), func(i int) bool {
		return ri.rIndexArr[i]==hash
	})
	if i<0||i>len(ri.rIndexArr)-1{
		return ""
	}
	node:=ri.rMap[ri.rIndexArr[i]]
	return node
}

func main()  {
fileList:=[]string{"123","456","789"}
hashMap:=&ring{
	map[uint32]string{},
	rIndex{},
	new(sync.RWMutex),
}
for _,v:=range fileList{
	index:=crc32.ChecksumIEEE([]byte(v))//循环索引
	hashMap.rMap[index]=v
	hashMap.rIndexArr=append(hashMap.rIndexArr,index)
}
sort.Sort(&hashMap.rIndexArr)
fmt.Println(hashMap)
}