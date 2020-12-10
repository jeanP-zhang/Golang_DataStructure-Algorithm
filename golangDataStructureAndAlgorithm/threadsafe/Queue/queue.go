package Queue

import "sync"

type Queue struct {
	queue []interface{}
	len   int
	lock  *sync.Mutex //互斥锁
}

//新建一个队列
func NewQueue() *Queue {
	q := &Queue{}
	q.queue = make([]interface{}, 0)
	q.len = 0
	q.lock = new(sync.Mutex)
	return q
}

//解决了线程安全
func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len
}

func (q *Queue) isEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len == 0
}
func (q *Queue) Shift() interface{} {

	q.lock.Lock()
	defer q.lock.Unlock()
	el := q.queue[0]
	q.queue = q.queue[1:]
	return el
}
func (q *Queue) Push(el interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, el)
	q.len++
	return
}
func (q *Queue) Peek() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue[0]
}
