package main

import "fmt"

func search(arr[] int,elem int)int{
	fmt.Println(arr)
	if len(arr)>0{

	if arr[len(arr)/2]==elem{
		return len(arr)/2
	}else if elem<arr[len(arr)/2] && elem>arr[len(arr)/2-1]{
		return len(arr)/2
	}else if elem>arr[len(arr)/2] && elem<arr[len(arr)/2+1]{
		return len(arr)/2+1
	}else if elem<arr[len(arr)/2]{
		return search(arr[:len(arr)/2],elem)
	}else {
		return  search(arr[len(arr)/2:],elem)
	}
	}
	return -1
}


func main(){
	a:=[]int{1,3,5,7,9}
	fmt.Println(a)
	var input int
	fmt.Scanln(&input)
	fmt.Println(search(a,input))
}
