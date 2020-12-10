package main

import "fmt"

func GnomeSort(arr []int) []int {
	i := 1
	for i < len(arr) {
		if arr[i] >= arr[i-1] {
			i++
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			if i > 1 {
				i--
			}
		}
	}
	return arr
}

func main() {
	arr := []int{11, 2, 1, 3, 4, 5, 1, 6, 6, 2, 4, 6, 1, 2, 3, 4, 5, 61, 2, 4, 3, 65}
	GnomeSort(arr)
	fmt.Println(arr)
}
