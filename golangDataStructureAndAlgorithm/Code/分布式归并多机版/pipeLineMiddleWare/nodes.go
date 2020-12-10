package pipeLineMiddleWare

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"net"
	"sort"
	"time"
)

var startTime time.Time

func Init() {
	startTime = time.Now()
}
func UseTime() {
	fmt.Println(time.Since(startTime))
}

//内存排序
func InMemorySort(in <-chan int) <-chan int {
	out := make(chan int, 1<<10)
	go func() {
		data := make([]int, 0) //创建一个数组，存储数据并排序
		for v := range in {
			data = append(data, v)
		}
		fmt.Println("数据读取完成", time.Since(startTime))
		sort.Ints(data) //排序
		for _, v := range data {
			out <- v //压入数据
		}
		close(out) //关闭管道
	}()
	return out
}

//合并,两个有序管道的 数据，归并有序的数据压入到另外 一个管道
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1<<10)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		//归并排序
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1 //取出v1,压入，再读取v1
				v1, ok1 = <-in1
			} else {
				out <- v2 //取出v2,压入，再读取v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}

//写入
func WriterSlink(writer io.Writer, in <-chan int) {

	for v := range in {
		//64位，8字节
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(v)) //字节转换
		writer.Write(buf)                          //写入缓存
	}
}

//随机数数组
func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int() //压入随机数
		}
		close(out)
	}()
	return out
}

//多路合并
func MergeN(inputs ...<-chan int) <-chan int {
	m := len(inputs)
	if m == 1 {
		return inputs[0]
	} else {

		return Merge(MergeN(inputs[:m]...), MergeN(inputs[:m]...))
	}
}

//读取数据
func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1<<10)
	go func() {
		buf := make([]byte, 8) //64位系统
		readSize := 0
		for {
			n, err := reader.Read(buf)
			readSize += n
			if n > 0 {
				out <- int(binary.BigEndian.Uint64(buf)) //数据压入
			}
			if err != nil || chunkSize != -1 && readSize >= chunkSize {
				break //跳出循环
			}
		}
		close(out)
	}()
	return out
}
func ArraySource(num ...int) <-chan int {
	var out = make(chan int)
	go func() {
		for _, v := range num {
			out <- v //数组的数据压入进去
		}
		close(out)
	}()
	return out
}

//给我一个ip端口，127.0.0.1：8090
func NetWorkWrite(addr string, in <-chan int) {
	listens, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	go func() {
		defer listens.Close()
		conn, err := listens.Accept() //接受信息
		if err != nil {
			panic(err)
		}
		defer conn.Close()              //关闭连接
		writer := bufio.NewWriter(conn) //写入数据
		defer writer.Flush()
		WriterSlink(writer, in)
	}()
}

//给我一个端口，读取数据
func NetWorkRead(addr string) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		conn, err := net.Dial("tcp", addr)
		defer close(out)
		if err != nil {
			panic(err)
		}
		r := ReaderSource(bufio.NewReader(conn), -1)
		for v := range r {
			out <- v //压入数据
		}
	}()
	return out
}
