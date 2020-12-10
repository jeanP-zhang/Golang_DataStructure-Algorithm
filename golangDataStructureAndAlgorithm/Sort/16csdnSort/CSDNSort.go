package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
fi,err:=os.Open("C:\\Users\\小狗子\\Desktop\\01-详细摘要-张国强.docx",)
if err!=nil{
	fmt.Println("文件读取失败",err)
	return
}
path:="C:\\Users\\小狗子\\Desktop\\111.docx"
saveFail,_:=os.Create(path)
defer saveFail.Close()
defer fi.Close()//延迟关闭文件
save:=bufio.NewWriter(saveFail)//用于对象写入
br:=bufio.NewReader(fi)
for{
	line,_,err:=br.ReadLine()
	if err==io.EOF{
		break//跳出循环
	}
	linStr:=string(line)//读取转化为字符串
	myString:=strings.Split(linStr,"#")//字符串切割
//	fmt.Println(string(line))
	fmt.Fprint(save,myString[1])
}
save.Flush()//刷新
}
