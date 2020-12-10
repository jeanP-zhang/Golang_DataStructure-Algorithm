package Queue

import (
	"sync"
)

type Item interface {
	Less(than Item) bool
}
type Int int

func (x Int) Less(than Item) bool {
	return x < than.(Int)
}

//最小堆
type Heap struct {
	lock *sync.Mutex //线程安全
	data []Item      //数组
	min  bool        //是否最小堆
}

func NewHeap() *Heap {
	return &Heap{new(sync.Mutex), make([]Item, 0), true}
}
func NewMin() *Heap {
	return &Heap{new(sync.Mutex), make([]Item, 0), true}

}
func NewMax() *Heap {
	return &Heap{new(sync.Mutex), make([]Item, 0), false}
}
func (h *Heap) isEmpty() bool {
	return len(h.data) == 0
}
func (h *Heap) Len() int {
	return len(h.data)
}

//抓取数据
func (h *Heap) Get(index int) Item {
	return h.data[index]
}

//插入数据
func (h *Heap) Insert(it Item) bool {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.data = append(h.data, it)
	h.ShiftUp() //插入数据
	return true
}

func (h *Heap) Less(a, b Item) bool {
	if h.min {
		return a.Less(b)
	} else {
		return b.Less(a)
	}
}

//压缩
func (h *Heap) Extract() (el Item) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.Len() == 0 {
		return nil //长度为0，不需要处理
	}
	el = h.data[0]
	last := h.data[h.Len()-1]
	if h.Len() == 1 {
		h.data = nil //弹出一个数据
		return nil
	}
	h.data = append([]Item{last}, h.data[1:]...)

	h.ShiftDown()

	return el
}

//弹出一个极小值
func (h *Heap) ShiftDown() {
	//堆排序得循环过程
	for i, child := 0, 1; i <= h.Len() && (i*2+1) < h.Len(); i = child {
		child = i*2 + 1
		if child+1 < h.Len()-1 && h.Less(h.Get(child+1), h.Get(child)) {
			child++ //循环左右节点过程
		}
		if h.Less(h.Get(i), h.Get(child)) {
			break
		}
		h.data[i], h.data[child] = h.data[child], h.data[i] //处理数据交换
	}
}

//弹出一个极大值
func (h *Heap) ShiftUp() {
	for i, parent := h.Len()-1, h.Len()-1; i > 0; i = parent {
		parent = i / 2
		if h.Less(h.Get(i), h.Get(parent)) {
			h.data[parent], h.data[i] = h.data[i], h.data[parent]
		} else {
			break
		}
	}
}
