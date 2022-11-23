package main

import "fmt"


type Human struct{
	gender string
}

func show(h *Human)(){
	h.gender ="SHOW : "+h.gender
	fmt.Println(h)
}

func printt(h Human)(){
	h.gender="PRINT : "+h.gender
	fmt.Println(h)
}


func main(){
	//h := Human{gender:"Female"}
	//h.show()

	//h := &Human{gender:"Female"}
	//h.show()

	//h := Human{gender:"Female"}
	//h.printt()


	//h := &Human{gender:"Female"}
	//h.printt()

	h := Human{gender:"Male"}
	//show(h)
	printt(h)
	fmt.Println(h)

	h1:= &Human{gender:"PMale"}
	show(h1)
	fmt.Println(h1)
}
