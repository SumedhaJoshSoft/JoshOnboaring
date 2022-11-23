package main

import "fmt"


func main(){
	days := map[int]string{
		1:"MM",
		2:"TT",
		3:"WE",
		4:"TH",
		5:"FR",
		6:"SA",
		7:"SU",
	}
	fmt.Println(days)
	var index int
	fmt.Scanln(&index)
	if day,ok:=days[index];ok{
		fmt.Println(day)
	}else{
		fmt.Println("NO")
	}
}
