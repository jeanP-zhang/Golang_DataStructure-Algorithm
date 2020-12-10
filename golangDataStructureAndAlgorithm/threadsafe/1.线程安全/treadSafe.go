package main

import (
	"fmt"
	"time"
)
//线程安全，多个线程访问一个资源，产生资源竞争，最终结果不正确
var money =0
func main()  {
for i:=0;i<1000;i++{
	go add(&money)
}
	time.Sleep(time.Second*20)
fmt.Println(money)
}

func add(pInt *int)  {
	for i:=0;i<100000;i++{
		*pInt++
	}


}
