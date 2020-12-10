package main

import (
	"bytes"
	"code.qiangqiang.com/studygo/golang-数据结构与算法/Code/分布式归并排序节点/pipeLineMiddleWare"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"time"
)

//func ()  {
//
//}
func BytesToInt(bts []byte) int {
	byteBuffer := bytes.NewBuffer(bts)
	var datas int64
	binary.Read(byteBuffer, binary.BigEndian, &datas)
	return int(datas)
}
func IntoBytes(n int) []byte {
	data := int64(n)
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, data)
	return byteBuffer.Bytes()
}
func ServerMsgHandler(conn net.Conn) <-chan int {
	buf := make([]byte, 16)
	arr := []int{}
	out := make(chan int, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Server close")
			return nil
		}
		if n == 16 {
			data1 := BytesToInt(buf[:len(buf)/2]) //取出第一个数据
			data2 := BytesToInt(buf[len(buf)/2:]) //取出第二个数据
			if data1 == 0 && data2 == 0 {
				arr = make([]int, 0) //开辟数据
			}
			if data1 == 1 {
				arr = append(arr, data2)
			}
			if data1 == 0 && data2 == 1 {
				fmt.Println("数组接收完成", arr)
				for i := 0; i < len(arr); i++ {
					out <- arr[i] //数组压入队列
				}
				close(out)
				return out

				arr = make([]int, 0)
			}
		}
	}
	return nil
}
func SendArray(arr []int, conn net.Conn) {

	length := len(arr)
	myBStart := IntoBytes(0)
	myBStart = append(myBStart, IntoBytes(0)...)
	conn.Write(myBStart)
	for i := 0; i < length; i++ {
		myBData := IntoBytes(1)
		myBData = append(myBData, IntoBytes(arr[i])...)
		conn.Write(myBData)
	}
	myBEnd := IntoBytes(0)
	myBEnd = append(myBEnd, IntoBytes(1)...)
	conn.Write(myBStart)
}

var sortResult []<-chan int

func main() {
	arrList := [][]int{{1, 9, 2, 8, 7, 3, 5, 6, 10, 4, 23, 24}, {11, 19, 12, 18, 17, 13, 15, 16, 101, 14, 123, 124}}
	//新的管道
	sortResult = make([]<-chan int, 128)
	for i := 0; i < 2; i++ {
		tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:700"+strconv.Itoa(1+i))
		if err != nil {
			panic(err)
		}
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			panic(err)
		}
		SendArray(arrList[i], conn)
		sortResult = append(sortResult, ServerMsgHandler(conn))
	}
	last := pipeLineMiddleWare.Merge(sortResult[0], sortResult[1])
	for v := range last {
		fmt.Printf("%d ", v)

	}
	time.Sleep(time.Second * 30)
}
