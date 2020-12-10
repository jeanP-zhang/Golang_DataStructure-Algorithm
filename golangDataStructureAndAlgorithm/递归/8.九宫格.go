//package main
//
//import "fmt"
//
//var pos [9]int
//var subNum = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//var zhiNum = []int{1, 2, 3, 5, 7, 11, 13, 17, 19}
//var down, up = 0, 9
//
////判断两数之和是否为复数
//func searchIsPrime(n int) bool {
//	for i := 0; i < 9; i++ {
//		if n == zhiNum[i] {
//			return true
//		}
//	}
//	return false
//}
//func check(i, n int) bool {
//	//纵
//	if i-3 >= 0 {
//		if !searchIsPrime(pos[i] + pos[i-3]) {
//			return false
//		}
//	}
//	if i%3 != 0 {
//		if !searchIsPrime(pos[i] + pos[i-1]) {
//			return false
//		}
//	}
//	//横
//	return true
//}
//func fillBox(i, n, r int, count *int) {
//	if i == n {
//		*count++
//		fmt.Printf("------------------------------------------------%d\n", *count)
//		for k := 0; k < r; k++ {
//			for j := 0; j < r; j++ {
//				fmt.Printf("%3d", pos[k*r+j])
//			}
//			fmt.Println()
//		}
//		return
//	} else {
//		for j := down; j <= up; j++ {
//			//放入
//			pos[i] = subNum[j]
//			if subNum[j] != -1 && check(i, n) {
//				subNum[j] = -1
//				fillBox(i+1, n, r, count)
//				subNum[j] = pos[i]
//			}
//		}
//	}
//}
//func main() {
//	count := 0
//	fillBox(0, 9, 3, &count)
//}
package main

import (
	"fmt"
)

var pos [9]int
var sub []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var num []int = []int{1, 2, 3, 5, 7, 11, 13, 17, 19}

/*从质数中查找，找到返回true*/
func searchFromNum(n int) bool {
	for i := 0; i < 9; i++ {
		if n == num[i] {
			return true
		}
	}
	return false
}

/*检验结果是否正确*/
func check(i, n int) bool {
	//纵向
	if i-3 >= 0 {
		if searchFromNum(pos[i]+pos[i-3]) == false {
			return false
		}
	}
	//横向
	if i%3 != 0 {
		if searchFromNum(pos[i]+pos[i-1]) == false {
			return false
		}
	}
	return true
}

var down, up = 0, 9

/*填入1~10到九宫格的解，回溯法*/
func fillBox(i, n, r int, count *int) {
	if i == n {
		*count++
		for i := 0; i < r; i++ {
			for j := 0; j < r; j++ {
				fmt.Printf("%3d", pos[i*r+j])
			}
			fmt.Println()
		}
		fmt.Println("============")
		return
	}
	for j := down; j <= up; j++ {
		//先放入
		pos[i] = sub[j]
		if sub[j] != -1 && check(i, n) {
			sub[j] = -1
			fillBox(i+1, n, r, count)
			sub[j] = pos[i]
		}

	}
}

func main() {
	count := 0
	fillBox(0, 9, 3, &count)
	fmt.Println(count)
}
