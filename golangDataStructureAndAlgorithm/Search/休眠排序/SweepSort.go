package main

import (
	"fmt"
	"time"
)

//写入1ns
//1亿数据
//多线程，分布式
var flag bool
var container chan bool

//func ()  {
//
//}
func main() {
	arr := []int{1, 2, 3, 4, 5}
	flag = true
	container = make(chan bool, 5) //5个管道
	for i := 0; i < 5; i++ {
		go toSleep(arr[i])
		fmt.Println("lalala")

	}
	go listen(len(arr))
	for flag {
		time.Sleep(1 * time.Second)
	}
}
func toSleep(data int) {
	times := time.Duration(data) * time.Second
	time.Sleep(times)
	fmt.Println(data)
	container <- true //管道输入true
}

var count int

func listen(size int) {

	for flag {
		select {
		case <-container:
			count++
			if count >= size {
				flag = false
				break
			}
		}
	}
}
