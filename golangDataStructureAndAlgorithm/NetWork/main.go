package main

import (
	"fmt"
	"net"
)

func MsgHandler(conNet net.Conn )  {
	buf:=make([]byte,1<<10)
	defer conNet.Close()
	for ; ;  {
		n,err:=conNet.Read(buf)
		if err!=nil{
			fmt.Println("conn close",conNet)
			panic(err)
		}
		fmt.Println("client data",string(buf[:n]))
     clientIP:=conNet.RemoteAddr()
     conNet.Write([]byte("hello"+clientIP.String()+"\n"))
	}
}
func main()  {
	addr:="127.0.0.1:8848"
	serverListener,err:=net.Listen("tcp",addr)//监听地址
if err!=nil{
	panic(err)//处理错误
}
defer serverListener.Close()//延迟关闭
for{
	newConn,err:=serverListener.Accept()
	if err!=nil{
		panic(err)//处理错误
	}
	go MsgHandler(newConn)//处理客户端消息
}//接受消息
}
