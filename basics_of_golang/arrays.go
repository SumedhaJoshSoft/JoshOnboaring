package main

import "fmt"

func main(){

	var a=[] int{1,2,3,4,11,12}
	fmt.Println(a)
	var b = [3]int{}
	fmt.Println(b)
	var c = [...]int{}
	fmt.Println(c)
	var d = [...]int{1,2,3,3}
	fmt.Println(d)
	var e = [2][3]int{
		[3]int{1,2,3},
		[3]int{4,5,6}}
	fmt.Println(e)
	var f = [2][3]int{
		{1,2,3},{3,4,5}}
	fmt.Println(f)

	var g = [2][3]int{}
	fmt.Println(g)
}

