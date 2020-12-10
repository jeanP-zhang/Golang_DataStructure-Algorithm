package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	times:=time.Now()
	path:="C:\\Users\\小狗子\\Documents\\WeChat Files\\wxid_aqasnyv71o4421\\FileStorage\\File\2020-06\\11.txt"
	QQFile,_:=os.Open(path)//打开文件

	defer QQFile.Close()//最后关闭文件
	br:=bufio.NewReader(QQFile)
	for{
		line,_,end:=br.ReadLine()//读取一行数据
	if end==io.EOF{
		break
	}
	fmt.Println(string(line))
		lineStr:=string(line)
		if strings.Contains(lineStr,"张国强"){
			fmt.Println(lineStr)
		}
	}
	fmt.Println(time.Since(times))
}
