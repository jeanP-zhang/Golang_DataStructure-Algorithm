package main

import "fmt"
const N=10
func HanIo(n int,a,b,c string)  {
	if n<1{
		return
	}
	if n==1{
		fmt.Printf("%s -> %s\n",a,c)
		move(a,c)
		show()
	}else{
		HanIo(n-1,a,c,b)
		fmt.Printf("%s -> %s\n",a,c)
		move(a,c)
		show()
		HanIo(n-1,b,a,c)
	}
}
var arr=[3][N]int{{0,0,0,0,0,0,0,0,0,0},{0,0,0,0,0,0,0,0,0,0},{0,0,0,0,0,0,0,0,0,0}}

func show()  {
	fmt.Printf("%5s%5s%5s\n","A","B","C")
	for i:=0;i<10 ; i++ {
		for j:=0;j<3 ;j++  {
			fmt.Printf("%5d",arr[j][i])
		}
		fmt.Println()
	}
}
func move(X,Y string)  {
	var m =int(X[0])-65//转化成数组0，1，2方便计算
	var n=int(Y[0])-65
	var iMove=-1//保存第一个不等于零的索引
	for i:=1;i<10 ;i++  {
		if arr[m][i]!=0{
			iMove=i
			break
		}
	}

	var jMove int
	if arr[n][N-1]==0{
		jMove=N-1
	}else{
		jMove=N
		for i:=0;i<10;i++{
			if arr[n][i]!=0{
				jMove=i
				break
			}
		}
		jMove-=1
	}
	arr[m][iMove],arr[n][jMove]=arr[n][jMove],arr[m][iMove]
}
func DataInit(n int)  {
	for i:=0;i<n;i++{
		arr[0][N-1-i]=n-i
		}

}
func main(){
	DataInit(3)
	HanIo(3,"A","B","C")
}