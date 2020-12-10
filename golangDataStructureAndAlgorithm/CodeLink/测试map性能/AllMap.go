package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 99999

func main() {
	allmap := make(map[int]string, N)
	path := "C:\\USERs\\111"
	sqlFile, _ := os.Open(path)
	br := bufio.NewReader(sqlFile)
	s := br.Size()
	fmt.Println(s)
}
