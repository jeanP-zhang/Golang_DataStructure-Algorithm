package main

import "fmt"

const Num  int=4
var count =1
var queues[Num][Num]int

func show()  {
	fmt.Printf("第%d种解法\n",count)
	for i:=0;i<Num ;i++  {
		for j:=0;j<Num ;j++  {
			if queues[i][j]==1{
				fmt.Printf("%s","O")
			}else {
				fmt.Printf("%s", "X")
			}
		}
		fmt.Println()
	}
}
func setQueue(row,col int)bool  {
	fmt.Println(row,col)
	if row==0{//第一个放入
		queues[row][col]=1//设置为1
		return true
	}
	show()
	queues[row][col]=1
	for i:=0;i<Num ;i++  {
		if queues[row][i]==1{//列有一个为1，无法放置
			return false
		}
	}
	for i:=0;i<Num ;i++  {
		if queues[i][row]==1{//行有一个为1，无法放置
			return false
		}
	}
	for i,j:=row,col;i<Num&&j<Num ;i,j=i+1,j+1  {
			if queues[i][j]==1{//对角线有一个为1，无法放置
				return false
			}
	}
	for i,j:=row,col;i>=0&&j>=0 ;i,j=i-1,j-1  {
		if queues[i][j]==1{//对角线有一个为1，无法放置
			return false
		}
	}
	for i,j:=row,col;i<Num&&j>=0 ;i,j=i+1,j-1  {
		if queues[i][j]==1{//对角线有一个为1，无法放置
			return false
		}
	}
	for i,j:=row,col;i>=0&&j<=Num ;i,j=i-1,j+1  {
		if queues[i][j]==1{//对角线有一个为1，无法放置
			return false
		}
	}
	queues[row][col]=1

	return true
}
func solveQueue(row int)  {
	if row==Num{
	show()
		count++
		return
	}
	for i:=0;i<Num ;i++  {
		if setQueue(row,i){

			solveQueue(row+1)//暴力产生循环
		}
		queues[row][i]=0//回退设置为0
	}
}
func main()  {
	solveQueue(0)
}