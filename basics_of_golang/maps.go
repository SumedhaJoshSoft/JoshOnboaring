package main

import "fmt"

func main(){

timeZone:=map[string]string{
	"UTC":"AAA",
	"IST":"SSS",
}
var seconds string
var ok bool
seconds, ok = timeZone["UTC"]
fmt.Println(seconds,ok)
seconds, ok = timeZone["UTC1"]
fmt.Println(seconds,ok)
}
