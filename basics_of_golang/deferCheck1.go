package main

import "fmt"

func first(s string) string{
	fmt.Println("First "+s)
	//return "First "+s
	return ""
}

func second(s string) string{
	fmt.Println("Second "+s)
	//return "Second "+s
	return ""
}

func third(s string) string{
	fmt.Println("Third "+s)
	//return "third "+s
	return ""
}

func dummy(){
	fmt.Println("Dummy...!")
}

func main(){
	fmt.Println("Start of main")
	defer first(second(third("A")))

	defer dummy()
	fmt.Println("End of of main")
	defer first(">>First again")

}


