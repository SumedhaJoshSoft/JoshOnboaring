package main

import "fmt"

func main(){
	a:=[]int{1,2,3,4}
	b:=[]int{8,9}
	for i,j:=range a,b{
		fmt.Println(i,j)
	}

}
