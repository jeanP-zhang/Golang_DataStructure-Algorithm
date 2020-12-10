package main

type Set struct {
	myList *List //内部实现基于链表
}
type SerIterator struct {
	index uint64 //索引
	mySet *Set   //集合
}

func (st *Set) GetAt(index uint64) Object {
	return (*st).myList.GetAt(index) //二次调用
}

//返回集合的大小
func (st *Set) GetSize() uint64 {
	return st.myList.GetSize()
}

//初始化
func (st *Set) Init(match ...MatchFun) {
	myLists := new(List) //新的链表
	st.myList = myLists
	if len(match) == 0 {
		myLists.Init() //初始化
	} else {
		myLists.Init(match[0])
	}
}

//判断数据是否存在
func (st *Set) IsIn(data Object) bool {
	return st.myLi
}

//插入
func (st *Set) Insert(data Object) bool {
	if !st.IsIn(data) {
		return st.myList.Append(data)
	}
	return false
}

//判断
func (st *Set) IsEmpty() bool {
	return st.myList.IsEmpty()
}

//删除
func (st *Set) Remove(data Object) bool {
	return st.myList.Remove(data)
}

//1.2.3
//1,3,4
func (st *Set) Union(set1 *Set) *Set {
	if set1 == nil {
		return st
	}
	if st == nil {
		return set1
	}
	//集合:新建的几何存储新的结构
	nSet := new(Set)
	nSet.Init(st.myList.myMatch)
	if st.IsEmpty() && set1.IsEmpty() {
		return nSet
	}
	for i := uint64(0); i < st.GetSize(); i++ {
		nSet.Insert(st.GetAt(i)) //插入数据
	}
	var data Object //判断set1的数据在st中是否存在，存在则跳过，不存在则插入
	for i := uint64(0); i < set1.GetSize(); i++ {
		data = set1.GetAt(i)
		if !nSet.IsIn(data) {
			nSet.Insert(data)
		}
	}
	return nSet
}

//123
//235 //23
func (st *Set) Share(set1 *Set) *Set {
	if set1 == nil {
		return nil
	}
	if st == nil {
		return nil
	}
	nSet := new(Set)
	nSet.Init(st.myList.myMatch)
	if st.IsEmpty() && set1.IsEmpty() {
		return nSet //空集合
	}
	largeSet := st   //保存最多元素
	smallSet := set1 //保存较小元素
	if set1.GetSize() > st.GetSize() {
	}
	largeSet, smallSet = set1, st
	var data Object
	for i := uint64(0); i < largeSet.GetSize(); i++ {
		data = largeSet.GetAt(i)
		if smallSet.IsIn(data) {
			nSet.Insert(data)
		}
	}
	return nSet
}

//差集
func (st *Set) Different(set1 *Set) *Set {
	if set1 == nil {
		return st
	}
	if st == nil {
		return set1
	}
	nSet := new(Set)
	nSet.Init(st.myList.myMatch)
	if st.IsEmpty() && set1.IsEmpty() {
		return nSet //空集合
	}
	var data Object
	for i := uint64(0); i < st.GetSize(); i++ {
		data = st.GetAt(i) //保存st有而set1没有
		if !set1.IsIn(data) {
			nSet.Insert(data)
		}
	}
	return nSet
}
func (st *Set) isSub(subSet *Set) bool {
	if subSet == nil {
		return true
	}
	if st == nil {
		return false
	}
	for i := uint64(0); i < subSet.GetSize(); i++ {
		if !st.IsIn(subSet.GetAt(i)) {
			return false //有一个不存在就不是子集
		}

	}
	return true
}
func match(data1, data2 Object) int {
	if data2 == data1 {
		return 0
	} else {
		return 1
	}
}
func (st *Set) IsEquals(subSet *Set) bool {
	if st == nil || subSet == nil {
		return false
	}
	nSet := st.Share(subSet)
	return nSet.GetSize() == st.GetSize()
}
func (st *Set) GetIterator() *SerIterator {
	it := new(SerIterator)
	it.index = 0
	it.mySet = st
	return it
}
func (it *SerIterator) HasNext() bool {
	set := it.mySet
	index := it.index
	return index < set.GetSize() //判断是否有下一个
}

func (it *SerIterator) Next() Object {
	set := it.mySet
	index := it.index
	if index < set.GetSize() { //取出数据
		data := set.GetAt(index)
		it.index++
		return data
	}
	return nil
}
func main() {
	mySet1 := new(Set)
	mySet1.Init(match)
	mySet2 := new(Set)
	mySet2.Init(match)
	for it := mySet1.GetIterator(); it.HasNext(); {
		it.Next()
	}
}
