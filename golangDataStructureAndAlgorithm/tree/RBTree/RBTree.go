package main

import "fmt"

func main() {
	rbtTree := NewRBTree()
	for i := 0; i < 10000; i++ {
		rbtTree.Insert(Int(i))
	}

	//for j := 0; j < 9000; j++ {
	//	rbtTree.Delete(Int(j))
	//}
	fmt.Println(rbtTree.GetDepth())
}
