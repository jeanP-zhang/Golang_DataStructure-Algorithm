package main

import (
	"bufio"
	"code.qiangqiang.com/studygo/golang-数据结构与算法/Code/分布式归并排序原理/pipeLineMiddleWare"
	"fmt"
	"os"
	"time"
)

func main1() {
	var fileName = "data.in" //文件写入
	var count = 1000000
	file, _ := os.Create(fileName)
	defer file.Close()                               //延迟关闭文件
	myPipe := pipeLineMiddleWare.RandomSource(count) //管道装随机数
	writer := bufio.NewWriter(file)
	pipeLineMiddleWare.WriterSlink(writer, myPipe)
	writer.Flush() //刷新
	file, _ = os.Open(fileName)
	defer file.Close() //延迟关闭文件
	myPipeRead := pipeLineMiddleWare.ReaderSource(bufio.NewReader(file), -1)
	counter := 0
	for v := range myPipeRead {
		fmt.Println(v)
		counter++
		if counter > 1000 {
			break
		}

	}
}

//没有开启任何一个goRuntine
func main12() {
	go func() {
		myp := pipeLineMiddleWare.Merge(
			pipeLineMiddleWare.InMemorySort(pipeLineMiddleWare.ArraySource(3, 9, 2, 1, 10)),
			pipeLineMiddleWare.InMemorySort(pipeLineMiddleWare.ArraySource(13, 19, 12, 11, 120)))
		for v := range myp {
			fmt.Println(v)
		}
	}()

	time.Sleep(time.Second * 10)
}
