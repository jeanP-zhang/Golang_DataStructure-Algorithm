package main

import "fmt"

func main() {
	i, x, y := 100, 0, 0
	for i > 0 {
		i--
		x = i % 8
		if x == 1 {
			y++
		}
	}
	fmt.Println(y)
}
