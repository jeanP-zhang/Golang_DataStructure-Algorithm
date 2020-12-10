package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"net"
	"os"
	"sort"
)

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

//字节转化
func BytesToInt(b []byte) int {
	bytesBuf := bytes.NewBuffer(b) //空的字节数据
	var data int64
	binary.Read(bytesBuf, binary.BigEndian, &data)
	return int(data)
}
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
func MsgHandler(conn net.Conn) {
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
				sort.Ints(arr)
				fmt.Println("数组排序完成", arr)
				sendArray(arr, conn) //排序完的数组返回给服务器
				arr = make([]int, 0)
			}
		}
		fmt.Println("Client Send", string(buf))
		conn.Write([]byte("收到" + string(buf)))
		msg := buf[:1] //备份buf
		beactch := make(chan byte)
		go HeartBeat(conn, beactch, 30)
		go HeartChanHandler(msg, beactch)
	}
}

//心跳机制
func HeartBeat(conn net.Conn, heartChan chan byte, timeout int) {
	select {
	case hc := <-heartChan:
		fmt.Println("heartChan", string(hc))
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break
	case <-time.After(time.Second * 30):
		fmt.Println("time is out")
		conn.Close()
	}
}
func HeartChanHandler(n []byte, beatch chan byte) {
	for _, v := range n {
		beatch <- v //管道压入数据
	}
	close(beatch)
}
func main() {
	server, err := net.Listen("tcp", "localhost:7000") //创建服务器
	CheckError(err)
	defer server.Close()
	for {
		newConn, err := server.Accept()
		CheckError(err)
		go MsgHandler(newConn)
	}
}
