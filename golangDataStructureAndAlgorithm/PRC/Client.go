package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

//RPC（远程协议调用）协议标准：
/*1.字母必须是大写
2.必须有返回值
3.函数还要有一个返回值err
4.第一个参数是接受的参数，第二个参数是返回给客户端的参数，而且第二个参数是指针类型*/
//通过网络通信进行调用

type Args struct {
	A, B int //两个数据
}
type Query struct {
	X, Y int //两个数据
}
type Last int

func (l *Last) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B //乘法
	fmt.Println(reply, "乘法执行了")
	return nil
}
func (l *Last) Divide(args *Args, query *Query) error {
	if args.B == 0 {
		return errors.New("除数为0")
	}
	query.X = args.A / args.B
	query.Y = args.A % args.B
	fmt.Println(query, "除法执行了")
	return nil
}
func main() {
	la := new(Last)
	fmt.Println(la, "=la")
	rpc.Register(la) //注册类型
	rpc.HandleHTTP() //设定http类型，开启服务
	//err := http.ListenAndServe(":7001", nil)
	list, err := net.Listen("tcp", "127.0.0.1:7001")
	if err != nil {
		panic(err)
	}
	http.Serve(list, nil)

}
