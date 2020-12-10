package main

import (
	"bytes"
	"encoding/binary"
	"net"
	"time"
)

func IntoBytes(n int) []byte {
	data := int64(n)
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, data)
	return byteBuffer.Bytes()
}
func SeverMsgHandler(conn net.Conn) {

}
func main() {
	addrs := "127.0.0.1:8848"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addrs)
	if err != nil {
		panic(err)
	}

	arr := []int{1, 9, 2, 8, 7, 3, 5, 6, 10, 4}
	length := len(arr)
	//-1  "1"开始
	//8 abcdefgh
	//4 abcd
	//-1 "0" 结束

	conn, err := net.DialTCP("tcp", nil, tcpAddr) //链接
	go SeverMsgHandler(conn)                      //收消息
	if err != nil {
		panic(err)
	}

	//0 0代表开始传输
	//1 1
	//1 9
	//1 2
	//1 8
	//......

	//0 1代表结束传输
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
	time.Sleep(time.Second * 30)
	//for {
	//	var inputStr string
	//	fmt.Scanln(&inputStr)
	//	conn.Write([]byte(inputStr))
	//	buf := make([]byte, 1<<10)
	//	n, _ := conn.Read(buf) //读取数据
	//	fmt.Println(string(buf[:n]))
	//}
}

//0 helloWodrd//数据
//1 cale 命令
