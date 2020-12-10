package Queue

type PriorityItem struct {
	value    interface{} //数据
	Priority int         //优先级
}

//基于堆实现优先队列
type PriorityQueue struct {
	data *Heap
}

//队列中间的元素
func NewPriorityItem(value interface{}, Priority int) *PriorityItem {
	return &PriorityItem{value, Priority}
}

func (p PriorityItem) Less(than Item) bool {
	return p.Priority < than.(PriorityItem).Priority
}
func NewMaxPriorityQueue() *PriorityQueue {
	return &PriorityQueue{NewMax()}
}
func (pq *PriorityQueue) Len() int {
	return pq.data.Len()
}
func (pq *PriorityQueue) Insert(el PriorityItem) {
	pq.data.Insert(el)
}
func (pq *PriorityQueue) Extract() PriorityItem {
	return pq.data.Extract().(PriorityItem)

}
func (pq *PriorityQueue) ChangePriority(val interface{}, Priority int) {
	var storage = NewQueue() //使用队列备份数据
	popped := pq.Extract()   //拿出最小的数值
	for val != popped.value {
		if pq.Len() == 0 {
			return
		}
		storage.Push(popped) //Push函数
		popped = pq.Extract()
	}
	popped.Priority = Priority
	pq.data.Insert(popped)
	for storage.Len() > 0 {
		pq.data.Insert(storage.Shift().(Item)) //其余数据重新放入优先队列
	}
}
