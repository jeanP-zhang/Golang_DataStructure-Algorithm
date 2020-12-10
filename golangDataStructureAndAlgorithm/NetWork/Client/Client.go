package main

import (
	"fmt"
	"net"
)

func main() {
	addrs := "127.0.0.1:8848"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addrs)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr) //链接
	if err != nil {
		panic(err)
	}

	for {
		var inputStr string
		fmt.Scanln(&inputStr)
		conn.Write([]byte(inputStr))
		buf := make([]byte, 1<<10)
		n, _ := conn.Read(buf) //读取数据
		fmt.Println(string(buf[:n]))
	}
}

//0 helloWodrd//数据
//1 cale 命令
