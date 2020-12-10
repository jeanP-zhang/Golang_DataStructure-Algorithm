package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func IntToBytes(n int) []byte {
	data := int64(n)
	bytesBuf := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuf, binary.BigEndian, data) //写入数据
	return bytesBuf.Bytes()
}
func BytesToInt(b []byte) int {
	bytesBuf := bytes.NewBuffer(b) //空的字节数组
	var data int64
	binary.Read(bytesBuf, binary.BigEndian, &data)
	return int(data)
}
func sendArray(arr []int, conn net.Conn) {
	length := len(arr)
	myBStart := IntToBytes(0)
	myBStart = append(myBStart, IntToBytes(0)...)
	myBEnd := IntToBytes(0)
	conn.Write(myBStart)
	for i := 0; i < length; i++ {
		myBData := IntToBytes(1)
		myBData = append(myBData, IntToBytes(arr[i])...)
		conn.Write(myBData)
	}
	myBEnd = append(myBEnd, IntToBytes(1)...)
	conn.Write(myBEnd)
}
func ServerMsgHandler(conn net.Conn) <-chan int {
	out := make(chan int, 1<<10)
	buf := make([]byte, 1<<4)
	defer conn.Close()
	arr := []int{}
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Client 关闭", conn.RemoteAddr())
			return nil
		}
		if n == 16 {
			data1 := BytesToInt(buf[:len(buf)/2])
			data2 := BytesToInt(buf[len(buf)/2:])
			if data1 == 0 && data2 == 2 {
				arr = make([]int, 0)
			}
			if data1 == 1 {
				arr = append(arr, data2)
			}
			if data1 == 0 && data2 == 1 {
				fmt.Println("数组接收完成", arr)
				for _, v := range arr {
					out <- v //数组压入管道
				}
				close(out)
				arr = make([]int, 0)
				return out
			}
		}
	}
	return nil
}
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	v1, ok1 := <-in1
	v2, ok2 := <-in2
	for ok1 || ok2 {
		if !ok2 || (ok1 && v1 <= v2) {
			out <- v1
			v1, ok1 = <-in1
		} else {
			out <- v2
			v2, ok2 = <-in2
		}
	}
	close(out)
	return out

}
func main() {
	arrList := [][]int{{1, 9, 2, 7}, {100, 103, 101, 102}}
	sortResult := []<-chan int{}
	for i := 0; i < 2; i++ {
		tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:700"+strconv.Itoa(i))
		CheckError(err)
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		CheckError(err)
		sendArray(arrList[i], conn)
		sortResult = append(sortResult, ServerMsgHandler(conn)) //处理数据接收
	}
	last := Merge(sortResult[0], sortResult[1])
	for v := range last {
		fmt.Printf("%d ", v)
	}
	time.Sleep(30 * time.Second)
}
