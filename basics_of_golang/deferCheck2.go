package main

import "fmt"

func main(){
	fmt.Println("Start of main")
	defer func(){
		if r:=recover();r!=nil{
		fmt.Println(r)
	}}()
	a()
	fmt.Println("End of main")
}

func a(){
	fmt.Println("Start of A")
	defer fmt.Println("AAA")
	defer fmt.Println("A1")
	panic("XXXXXXXXXXXXXXXXXXXXXXXXX")
	defer fmt.Println("A2")
	fmt.Println("BB")
}
