package main

import(
	"fmt"
	"strings"
)

func main(){
	a:= "sumedha sunil jagtap"
	a = a+string('C')
	fmt.Println(strings.SplitN(a," ",1))
	fmt.Println(strings.SplitAfter(a," "))
	fmt.Println(strings.SplitAfterN(a," ",1))
	//fmt.Println(strings.Cut(a," "))

	//index:=string.Index(a," ")
	//fmt.Println(index)
	for i :=range strings.SplitN(a," ",1){
		fmt.Println(i)
	}


	for _, c:= range a{
		switch c{
			case 's':
				fmt.Println("AAAAAAA :: ")
			default:
				fmt.Println("KKK")
		}
		fmt.Println(c)
	}
}
