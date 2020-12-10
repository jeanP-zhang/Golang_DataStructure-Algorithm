package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)
//整数的转码

//func  IntoBytes(n int)[]byte  {
//	data:=int64(n)
//	byteBuffer:=bytes.NewBuffer([]byte{})
//	binary.Write(byteBuffer,binary.BigEndian,data)
//	return byteBuffer.Bytes()
//}
//func BytesToInt(bts []byte) int {
//	byteBuffer:=bytes.NewBuffer(bts)
//	var datas int64
//	binary.Read(byteBuffer,binary.BigEndian,&datas)
//	return int(datas)
//}
func main()  {
	a:=0b10
fmt.Println(IntoBytes(1))
fmt.Printf("%d\n",a)
fmt.Println(string([]byte("123")))
myByte:=IntoBytes(1)
myByte=append(myByte,IntoBytes(2)...)
	fmt.Println(bytes.Join(IntoBytes(1),IntoBytes(2)))//拼接两个字符
}