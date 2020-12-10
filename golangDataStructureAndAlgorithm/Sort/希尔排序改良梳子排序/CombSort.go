package main

//func shellSortStep() {
//
//}
func CombSort(arr []int) []int {
	length := len(arr)
	gap := length
	for gap > 1 {
		gap = gap * 10 / 13
		for i := 0; i+gap < length; i++ {
			if arr[i] > arr[i+gap] { //收缩
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
			}
		}
	}
	return arr
}
func main() {

}
