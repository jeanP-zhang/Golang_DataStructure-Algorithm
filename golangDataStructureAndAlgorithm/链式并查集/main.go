package main

import "fmt"

func main() {
	seq := []int{4, 8, offLineMinimumExtract,
		3,
		offLineMinimumExtract,
		9, 2, 6,
		offLineMinimumExtract,
		offLineMinimumExtract,
		offLineMinimumExtract,
		1, 7, offLineMinimumExtract,
		5}
	exp := []int{4, 3, 2, 6, 8, 1}
	last := OffLineMinimum(seq)

	fmt.Println(seq)
	fmt.Println(exp)
	fmt.Println(last)
}
