package main

import (
	"fmt"
	"sync"
	"time"
)
//线程安全，多个线程访问一个资源，产生资源竞争，最终结果不正确
var money =0
var lock *sync.RWMutex
func main()  {
	lock=new(sync.RWMutex)
	for i:=0;i<1000;i++{
		go add(&money)
	}
	time.Sleep(time.Second*20)
	fmt.Println(money)
}

func add(pInt *int)  {
	lock.Lock()
	for i:=0;i<100000;i++{
		*pInt++
	}
lock.Unlock()

}
