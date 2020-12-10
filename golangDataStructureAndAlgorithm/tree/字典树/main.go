package main

import "fmt"

func rev(keyWord string) string {
	//length := len(keyWord)
	runes := []rune(keyWord)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func main() {
	/*	s := NewKeyWordTree()
		s.Put(1, "ba")
		s.Put(2, "ab")
		s.Put(3, "abc")
		s.Put(4, "abcd")
		s.Put(5, "bcd")
		s.DeBugOut()
		fmt.Println(s.Sugg("b", 4))*/
	fmt.Println(rev("abcd123"))
}
