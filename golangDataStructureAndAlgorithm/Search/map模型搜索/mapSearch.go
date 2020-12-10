package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const N = 84331445

func main() {
	allStrs := make(map[int]string, N)
	path := "C:\\Users\\小狗子\\Documents\\WeChat Files\\wxid_aqasnyv71o4421\\FileStorage\\File\2020-06\\11.txt"
	QQFile, _ := os.Open(path) //打开文件
	defer QQFile.Close()       //最后关闭文件
	i := 0                     //统计一共多少行
	br := bufio.NewReader(QQFile)
	for {
		line, _, end := br.ReadLine() //读取一行数据
		if end == io.EOF {            //文件关闭，跳出循环
			break
		}
		lineStr := string(line)
		lines := strings.Split(lineStr, "----")
		if len(lines) == 2 {
			//allStrs[i].QQPass = lines[1]
			//allStrs[i].QQUser, _ = strconv.Atoi(lines[0])
			QQPass := lines[1]
			QQUser, _ := strconv.Atoi(lines[0])
			allStrs[QQUser] = QQPass
		}

		i++

	}
	fmt.Println("数据载入内存")
	time.Sleep(time.Second * 10)
	for {
		fmt.Println("请输入要查询的数据")
		var QQ int
		fmt.Scanf("%d", &QQ)
		startTime := time.Now()
		QQPass, ok := allStrs[QQ]
		if !ok {
			fmt.Println("不存在")
		} else {
			fmt.Println(QQ, QQPass)
		}
		for j := 0; j < N; j++ {
			if strings.Contains(allStrs[j].QQPass, QQ) {
				fmt.Println(j, allStrs[j].QQUser, allStrs[j].QQPass)
			}
		}
		fmt.Println("一共用了:", time.Since(startTime))
	}
}
