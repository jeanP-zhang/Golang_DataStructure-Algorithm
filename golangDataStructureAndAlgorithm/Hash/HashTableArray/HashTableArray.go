package HashTableArray

import (
	"crypto/sha256"
	"errors"
)

const (
	Deleted=iota//数据已经被删除
	MinTableSize=100//
	legimatc=iota//已经存在的合法数据
	Empty=iota//数据为空
)
//自己的hash算法。可靠性不高
func MySHA(str interface{},tableSize int) int {
	var HashVar =0
	var chars []byte
	if strings,ok:=str.(string);ok{
		chars=[]byte(strings)//字符串转化为字节数组
	}
	for _,v:=range chars{
		HashVar=(HashVar<<17)|12&1235^+int(v)//Hash算法
		/*左移17位再异或一个值*/
	}
	return HashVar%MinTableSize
}
func MySha256(str string,tableSize int)int{
	shaoBj:=sha256.New()
	shaoBj.Write([]byte(str))//哈希
	myBytes:=shaoBj.Sum(nil)
	var HashVar =0
	for _,v:=range myBytes{
		HashVar=(HashVar<<17|123&1235^139)+int(v)
	}
	return HashVar*MinTableSize
}
type HashFunc func(data interface{},tableSize int )int//函数指针
type HashEntry struct {
	data interface{}
	kind int
}
type HashTable struct {
tableSize int//哈希表的大小
theCells []*HashEntry//数组，每一个元素是指针指向hash结构
hashFun HashFunc
}
type HashTableGo interface {
	Find(data interface{})int//查找数据
	Insert(data interface{})//插入数据
	Empty()//为空
	GetValue(index int)interface{}//抓取value
}

func NewHashTable(size int,hash HashFunc)(*HashTable,error)  {
	if size<MinTableSize{
		return nil,errors.New("哈希表太小")
	}
	if hash==nil{
		return nil,errors.New("没有Hash函数")
	}
	HashTable:=new(HashTable)//创建hash表
	HashTable.tableSize=size
	HashTable.theCells=make([]*HashEntry,size)
HashTable.hashFun=hash//设置哈希函数
for i:=0;i<HashTable.tableSize;i++{
	HashTable.theCells[i]=new(HashEntry)
	HashTable.theCells[i].data=nil
	HashTable.theCells[i].kind=Empty
}
return HashTable,nil

}
func (ht *HashTable)Find(data interface{})int  {
	var colLid =0
	curpos:=ht.hashFun(data,ht.tableSize)//计算哈希位置
if ht.theCells[curpos].kind!=Empty&&ht.theCells[curpos].data!=data{
	colLid+=1
	curpos=2*curpos-1//平方探测
	if curpos>ht.tableSize{
		curpos-=ht.tableSize//越界，返回
	}
}
return curpos
}
func (ht *HashTable)Insert(data interface{}) {
pos:=ht.Find(data)//查找数据位置
entry:=ht.theCells[pos]//插入数据记录状态
if entry.kind!=legimatc{
	entry.kind=legimatc
	entry.data=data
}
}
func (ht *HashTable)Empty() {
for i:=0;i<ht.tableSize;i++{
	if ht.theCells[i]==nil{
		continue//继续循环清空数据
	}
	ht.theCells[i].kind=Deleted//删除数据
}
}
func (ht *HashTable)GetValue(index int)interface{} {
if index>ht.tableSize{
	return nil
}
entry:=ht.theCells[index]
if entry.kind==legimatc{
	return entry.data
}else{
	return nil
}
}