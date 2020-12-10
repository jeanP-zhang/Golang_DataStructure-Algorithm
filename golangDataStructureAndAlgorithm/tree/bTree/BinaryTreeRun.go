package main

import "fmt"

func main() {
	bst := NewBinaryTree()
	for i := 1; i <= 7; i++ {
		bst.Adds(i)
	}
	bst.InOrder()
	fmt.Println("------------------------")
	bst.PreOrder()
	fmt.Println("------------------------")
	bst.PostOrder()
	fmt.Println("------------------------")

}
