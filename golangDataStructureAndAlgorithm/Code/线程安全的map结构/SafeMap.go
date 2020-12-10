package main

import (
	"fmt"
	"sync"
	"time"
)

//map映射
//map管理数据，瞬间查找
//多线程需要维护线程安全
type SyncMap struct {
	myMap map[string]string
	lock  *sync.RWMutex
}

var sMap SyncMap
var done chan bool

func Write1() {
	keys := []string{"1", "2", "3"}
	for _, k := range keys {
		sMap.Lock()
		sMap.myMap[k] = k
		sMap.Unlock()
		time.Sleep(time.Second)
	}
	done <- true
}
func Write2() {
	keys := []string{"11", "22", "33"}
	for _, k := range keys {
		sMap.Lock()
		sMap.myMap[k] = k
		sMap.Unlock()
		time.Sleep(time.Second)
	}
	done <- true
}
func read() {
	sMap.RLock() //读锁
	fmt.Println("readLock")
	for k, v := range sMap.myMap {
		fmt.Println(k, v)
	}
	sMap.RUnlock()
}
func main() {
	smap := SyncMap{make(map[string]string), new(sync.RWMutex)}
	done = make(chan bool, 2)
	go Write1()
	go Write2()
	for {
		read()
		if len(done) == 2 {
			break
		} else {
			time.Sleep(time.Second)
		}
	}
}
