package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"time"
)

//处理心跳的channel
func HeartChanHandler(n []byte, beatch chan byte) {
	defer close(beatch) //关闭管道
	for _, v := range n {
		beatch <- v
	}
}

//判断30秒内有没有产生通信
//超过30秒就退出
func HeartBeat(conn net.Conn, heartChan chan byte, timeout int) {
	fmt.Println("HeartBeat")
	select {
	case hc := <-heartChan:
		{
			fmt.Println(string(hc))
			log.Println(string(hc))
			conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
			//计数器归零
		}
	case <-time.After(time.Second * 30):
		fmt.Println("time out", conn.RemoteAddr()) //客户端超时
		conn.Close()
	}

}

func MsgHandler(conNet net.Conn) {
	fmt.Println("HeartHandler")
	buf := make([]byte, 1<<10)
	defer conNet.Close()
	for {
		n, err := conNet.Read(buf)
		if err != nil {
			fmt.Println("conn close", conNet)
			panic(err)
		}
		msg := buf[1:n]
		if n != 0 {
			if string(buf[0]) == "0" {
				fmt.Println("client data", string(buf[1:n]))
				//clientIP := conNet.RemoteAddr()
				conNet.Write([]byte("收到数据" + string(buf[1:n]) + "\n"))
			} else {
				fmt.Println("client cmd", string(buf[1:n]))
				cmd := exec.Command(string(buf[1:n])) //执行命令
				cmd.Run()
				conNet.Write([]byte("收到命令" + string(buf[1:n]) + "\n"))
			}
			beatch := make(chan byte)
			go HeartBeat(conNet, beatch, 30)
			go HeartChanHandler(msg, beatch)
		}

		clientIP := conNet.RemoteAddr()
		conNet.Write([]byte("hello" + clientIP.String() + "\n"))
	}
}
func main() {
	addr := "127.0.0.1:8848"
	serverListener, err := net.Listen("tcp", addr) //监听地址
	if err != nil {
		panic(err) //处理错误
	}
	defer serverListener.Close() //延迟关闭
	for {
		newConn, err := serverListener.Accept()
		if err != nil {
			panic(err) //处理错误
		}

		go MsgHandler(newConn) //处理客户端消息
	} //接受消息
}
