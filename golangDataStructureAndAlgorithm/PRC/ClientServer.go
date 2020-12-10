package main

import (
	"fmt"
	"net/rpc"
)

//RPC（远程协议调用）协议标准：
/*1.字母必须是大写
2.必须有返回值*/
type ArgsX struct {
	A, B int //两个数据
}
type QueryX struct {
	X, Y int //两个数据
}

func main() {
	severIp := "127.0.0.1:7001"
	client, err := rpc.DialHTTP("tcp", severIp)
	if err != nil {
		fmt.Println(err)
	}
	i1 := 19
	i2 := 13
	args := ArgsX{i1, i2}
	var replay int
	err = client.Call("Last.Multiply", args, &replay)
	if err != nil {
		panic(err)
	}
	var qu QueryX
	err = client.Call("Last.Divide", args, &qu)
	fmt.Println(qu.X, qu.Y) //乘法
}
