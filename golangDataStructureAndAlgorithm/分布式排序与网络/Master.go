package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"

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
func doWork(conn net.Conn) error {
	ch := make(chan int, 100)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case stat := <-ch:
			if stat == 2 {
				return errors.New("服务器没有消息")
			}
		case <-ticker.C:
			{
				ch <- 1
				go ServerMsgHandler(conn)
			}
		case <-time.After(time.Second * 10):
			defer conn.Close()
			fmt.Println("关闭超时链接")

		}
	}
	return nil
}
func MasterHandler(conn net.Conn, ch chan int) {
	<-ch
	msg := time.Now().String() //消息
	sendArray([]int{2, 9, 7, 6, 4}, conn)
	fmt.Println("send over", msg)

}
func ServerMsgHandler(conn net.Conn) {
	buf := make([]byte, 1<<4)
	defer conn.Close()
	arr := []int{}
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Client 关闭", conn.RemoteAddr())
			panic(err)
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
				arr = make([]int, 0)
			}
		}
	}
}
func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:7000")
	CheckError(err)
	for {
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			fmt.Println(err)
		} else {
			sendArray([]int{2, 9, 7, 6, 4}, conn)
			fmt.Println("send over")
			defer conn.Close()
			doWork(conn)
		}
		time.Sleep(time.Second)
	}
}

//}
//	sendArray([]int{2, 9, 7, 6, 4}, conn)
//	fmt.Println("send over")
//	ch := make(chan int, 100) //交换消息
//	ticker := time.NewTicker(time.Second)
//	defer ticker.Stop()
//	for {
//		select {
//		case <-ticker.C:
//			ch <- 1
//			//go MasterHandler(conn, ch)
//			go ServerMsgHandler(conn)
//		case <-time.After(time.Second * 10):
//			defer conn.Close()
//			fmt.Println("time out")
//		}
//		//var inputStr string
//		//fmt.Scanln(&inputStr)
//		//conn.Write([]byte(inputStr))
//		//buf := make([]byte, 1024)
//		//n, _ := conn.Read(buf)
//		//fmt.Println(string(buf[:n]))
//	}
//
//}
