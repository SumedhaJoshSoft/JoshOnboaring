package main

import (
	"fmt"
	"strings"
)

func main(){
	str := "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^dadad#dsd$"
	str = "heyalice#heybob$$$howareyou#"
	fmt.Println(str[:strings.Index(str,"^")])

}
