package main

import (
	"code.qiangqiang.com/studygo/golang-数据结构与算法/Hash/HashTableArray"
	"fmt"
)

//func main()  {
//fmt.Println(HashTableArray.MySHA("abcd",100))
//	fmt.Println(HashTableArray.MySHA("abce",100))
//}
func main()  {
	MyTable,_:=HashTableArray.NewHashTable(100,HashTableArray.MySHA)
	MyTable.Insert("abcde1")
	MyTable.Insert("abcdf3")
	MyTable.Insert("abcde2")
	pos:=MyTable.Find("abcde1")
	fmt.Println(MyTable.GetValue(pos))
	pos=MyTable.Find("abcde2")
	fmt.Println(MyTable.GetValue(pos))
}