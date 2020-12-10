package main

import "net"

func main()  {
	serverListener,err:=net.Listen("tcp","127.0.0.1:7001")
	if err!=nil{
		panic(err)
	}
	defer serverListener.Close()
	for{
		newConn,err:=serverListener.Accept()
		if err!=nil{
			panic(err)
		}
	}
}