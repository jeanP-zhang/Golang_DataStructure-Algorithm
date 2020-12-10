package main

import "fmt"

//1 9 2 8 3 7 4 6 5 10
/*1          7
    9          4
      2          6
        8          5
           3         10
  1 4 2 5  3 7 9 6 8 10
步长收缩排序*/
func ShellSortStep(arr []int, started, gap int) {
	length := len(arr)
	for i := started + gap; i < length; i += gap { //插入排序的变种
		backup := arr[i]               //备份插入的数据
		j := i - gap                   //上一个位置循环找到位置插入
		for j > 0 && backup < arr[j] { //从前往后移动
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = backup //插入
	}
}
func shellSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	gap := length / 2
	for gap > 0 {
		for i := 0; i < gap; i++ { //处理每个元素的步长
			ShellSortStep(arr, i, gap)
		}
		gap /= 2 //gap--
	}
	return arr
}
func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(shellSort(arr))
}
