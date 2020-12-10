package main

import "fmt"

func HanIo(n int,a,b,c string)  {
	if n<1{
		return
	}
	if n==1{
		fmt.Printf("%s -> %s\n",a,c)
	}else{
		HanIo(n-1,a,c,b)
		fmt.Printf("%s -> %s\n",a,c)
		HanIo(n-1,b,a,c)
	}


}
func main(){
	HanIo(3,"A","B","C")
}