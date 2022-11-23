package main

import "fmt"


func main(){
	a:=[]int{1,2,3,4,5,6}
 	defer func(){
		if r:=recover();r!=nil{
		fmt.Println(r)
	}}()
	fmt.Println(a[9])
}

