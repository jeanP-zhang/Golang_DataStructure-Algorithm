package main

import "fmt"

var pos,b[80]int//10*8,pos[i]=j，第i行皇后在j位置,b[i]，第i行有没有位置
var c,d[150]int//c[i+j] i ->0--7 j >-0--7 i+j 0-14   d[j-i+7]反对角线
//填充数组0，1
func putN(i,j ,n int)  {
	pos[i],b[j],c[j-i+7],d[i+j]=j,n,n,n
}
//检查皇后
func checkPos(i,j int)bool {
	if b[j]==1||c[j-i+7]==1||d[i+j]==1{
return false
	}
	return true
}
func show(n int)  {

	for i:=0;i<n ;i++  {
		for j:=0;j<n ;j++  {
			if pos[i]==j{
				fmt.Printf("%4s","O")
			}else {
				fmt.Printf("%4s", "X")
			}
		}
		fmt.Println()

	}
}
func Queue(i,n int ,count *int)  {
	if i>7{
		fmt.Println("-----------------------------------------")
*count++
show(n)
		return
	}else{
		for j:=0;j<n;j++{
			if checkPos(i,j){
				putN(i,j,1)
				Queue(i+1,n,count)
				putN(i,j,0)
			}
		}
	}
}
func main()  {
	n,count:=8,0
	Queue(0,n,&count)
}