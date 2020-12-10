package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"sort"
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
	buf := make([]byte, 16)
	defer conNet.Close()
	arr := make([]int, 0) //切片保存数据
	for {
		n, err := conNet.Read(buf)
		if err != nil {
			fmt.Println("conn close", conNet)
			panic(err)
		}

		if n == 16 {
			data1 := BytesToInt(buf[:len(buf)/2]) //取出的第一个数据
			data2 := BytesToInt(buf[len(buf)/2:]) //取出的第二个数据
			if data1 == 0 && data2 == 0 {         //重新接收数组
				arr = make([]int, 0) //开辟数据
			}
			if data1 == 1 {
				arr = append(arr, data2)
			}
			if data2 == 0 && data2 == 1 {
				fmt.Println("数组接收完成", arr)
				sort.Ints(arr)
				arr = make([]int, 0)
				//写入
				myBStart := IntoBytes(0)
				myBStart = append(myBStart, IntoBytes(n)...)
				conNet.Write(myBStart)
				for i := 0; i < len(arr); i++ {
					myBData := IntoBytes(1)
					myBData = append(myBData, IntoBytes(arr[i])...)
					conNet.Write(myBData)
				}
			}
		}

		clientIP := conNet.RemoteAddr()
		conNet.Write([]byte("hello" + clientIP.String() + "\n"))
	}
}
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
func main() {
	addr := "127.0.0.1:8848"
	serverListener, err := net.Listen("tcp", addr) //监听地址
	if err != nil {
		panic(err) //处理错误
	}

	//length:=len(arr)

	//0 0代表开始传输
	//1 1
	//1 9
	//1 2
	//1 8
	//......
	//0 1代表结束传输
	defer serverListener.Close() //延迟关闭
	for {
		newConn, err := serverListener.Accept()
		if err != nil {
			panic(err) //处理错误
		}

		go MsgHandler(newConn) //处理客户端消息
	} //接受消息
}
