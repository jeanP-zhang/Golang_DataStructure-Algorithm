package main

import "fmt"

type Set struct {
	buf  []interface{}        //存储数据
	num  int                  //数量
	hash map[interface{}]bool //借助map实现映射
}

//新建一个可以变长的Set
func NewSet() *Set {
	return &Set{make([]interface{}, 0), 0, make(map[interface{}]bool)}
}
func (s *Set) Add(value interface{}) bool {
	if s.IsExist(value) {
		return false
	}
	s.buf = append(s.buf, value) //追加
	s.hash[value] = true
	s.num++
	return s.hash[value]
}
func (s *Set) IsExist(value interface{}) bool {
	return s.hash[value]
}
func main() {
	set := NewSet()
	set.Add(1)
	set.Add("哈哈")
	set.Add(0x80)
	fmt.Println(set)
}
