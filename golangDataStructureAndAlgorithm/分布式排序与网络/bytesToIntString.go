package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func IntToBytes(n int) []byte {
	data := int64(n)
	bytesBuf := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuf, binary.BigEndian, data) //写入数据
	return bytesBuf.Bytes()
}
func BytesToInt(b []byte) int {
	bytesBuf := bytes.NewBuffer(b) //空的字节数组
	var data int64
	binary.Read(bytesBuf, binary.BigEndian, &data)
	return int(data)
}
func main() {
	fmt.Println([]byte("1234"))
	fmt.Println(string([]byte("12345")))
}
