package main

import (
	"bufio"
	"code.qiangqiang.com/studygo/golang-数据结构与算法/Code/分布式归并排序原理/pipeLineMiddleWare"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func creatNetworkPipeline(fileName string, fileSize int, chunkCount int) <-chan int {
	out := make(chan int)
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	myPipe := pipeLineMiddleWare.RandomSource(fileSize / 8) //管道装随机数
	writer := bufio.NewWriter(file)                         //写入
	writer.Flush()                                          //刷新
	chunkSize := fileSize / chunkCount                      //大小
	var sortAddr = []string{}
	pipeLineMiddleWare.Init() //初始化
	files, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}
	for i := 0; i < chunkCount; i++ {
		files.Seek(int64(i*chunkSize), 0)                                           //移动文件指针位置
		source := pipeLineMiddleWare.ReaderSource(bufio.NewReader(file), chunkSize) //读取
		addr := ":" + strconv.Itoa(7000+i)                                          //开辟地址
		NetWorkWrite(addr, pipeLineMiddleWare.InMemorySort(source))
		sortAddr = append(sortAddr, addr) //地址复制
	}
	sortResults := []<-chan int{}
	for _, addr := range sortAddr {
		sortResults = append(sortResults, NetWorkRead(addr))
	}
	return pipeLineMiddleWare.MergeN(sortResults...)

}

//给我一个ip端口，127.0.0.1：8090
func NetWorkWrite(addr string, in <-chan int) {
	listens, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	go func() {
		defer listens.Close()
		conn, err := listens.Accept() //接受信息
		if err != nil {
			panic(err)
		}
		defer conn.Close()              //关闭连接
		writer := bufio.NewWriter(conn) //写入数据
		defer writer.Flush()
		pipeLineMiddleWare.WriterSlink(writer, in)
	}()
}

//给我一个端口，读取数据
func NetWorkRead(addr string) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		conn, err := net.Dial("tcp", addr)
		defer close(out)
		if err != nil {
			panic(err)
		}
		r := pipeLineMiddleWare.ReaderSource(bufio.NewReader(conn), -1)
		for v := range r {
			out <- v //压入数据
		}
	}()
	return out
}

//1.本地 归并排序2.多线程 3.分布式
func merge(leftArr, rightArr []int) []int {
	leftIndex := 0  //左边索引
	rightIndex := 0 //右边索引
	lastArr := make([]int, 0)
	for leftIndex < len(leftArr) && rightIndex < len(rightArr) {
		if leftArr[leftIndex] < rightArr[rightIndex] {
			lastArr = append(lastArr, leftArr[leftIndex])
			leftIndex++
		} else if leftArr[leftIndex] > rightArr[rightIndex] {
			lastArr = append(lastArr, rightArr[rightIndex])
			rightIndex++
		}
		lastArr = append(lastArr, leftArr[leftIndex])
		lastArr = append(lastArr, rightArr[rightIndex])
		leftIndex++
		rightIndex++
	}
	for leftIndex < len(leftArr) {
		lastArr = append(lastArr, leftArr[leftIndex])
		leftIndex++
	}
	for rightIndex < len(rightArr) {
		lastArr = append(lastArr, rightArr[rightIndex])
		rightIndex++
	}
	return lastArr
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr //小于10改用插入排序
	}
	mid := length / 2
	leftArr := mergeSort(arr[:mid])
	rightArr := mergeSort(arr[mid+1:])
	return merge(leftArr, rightArr)
}

//中间件  有的程序在中间完成配合步骤的
//生产随机数据
//func main() {
//
//}

//多线程
func creatPipeLine(fileName string, fileSize int, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	sortResults := make([]<-chan int, 0)
	pipeLineMiddleWare.Init() //初始化
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkCount), 0)                                           //跳到文件指针
		source := pipeLineMiddleWare.ReaderSource(bufio.NewReader(file), chunkSize) //读取
		sortResults = append(sortResults, pipeLineMiddleWare.InMemorySort(source))
	}
	return pipeLineMiddleWare.MergeN(sortResults...)
}

//写入文件
func writeToFile(in <-chan int, fileName string) {
	file, err := os.Open(fileName) //打开文件
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()                       //刷新
	pipeLineMiddleWare.WriterSlink(writer, in) //写入数据
}

//显示文件
func showFile(fileName string) {
	file, err := os.Create(fileName) //打开文件
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := pipeLineMiddleWare.ReaderSource(bufio.NewReader(file), -1)
	counter := 0
	for v := range p {
		fmt.Println(v)
		counter++
		if counter > 1000 {
			break
		}
	}
}
func main1X() {
	go func() {

		p := creatPipeLine("big.in", 80, 2) //80数据切割成4段
		writeToFile(p, "big.out")
		showFile("big.out")
	}()
	time.Sleep(time.Second * 10)
}
