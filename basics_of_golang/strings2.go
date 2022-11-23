package main

import (
	"fmt"
	"strings"
)

func main(){
	str := "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^wyquwyuqhuduhqwhdq"
	str ="heyalice#heybob$$$howareyou#"	
	str = "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^dadad#dsd$"
	var aliceCount, bobCount = strings.Count(str,"$"),strings.Count(str,"#")
	fmt.Println(aliceCount, bobCount)

	var arr  []string

	var aliceChan, bobChan = make(chan string,aliceCount),make(chan string,bobCount)

	var msg string

	Loop:
	for _, ch := range str{
	switch ch{
		case '$':
			arr = append(arr,msg)
			aliceChan <- msg
			msg = ""
		case '#':
			arr = append(arr,msg)
			bobChan <- msg
			msg = ""
		case '^':
			break Loop
		default:
			msg = msg + string(ch)
		}
	}
	fmt.Println(arr)
	close(aliceChan)
	close(bobChan)


	for i:=0;i<aliceCount+bobCount;i++{
		select{
			case aliceMsg := <-aliceChan:
				fmt.Println("A : ",aliceMsg)
			case bobMsg := <-bobChan:
				fmt.Println("B : ",bobMsg)
		}
	}
//	for i:=0;i<aliceCount;i++{
//		fmt.Println(<-aliceChan)
//	}
//
//	for i:=0;i<bobCount;i++{
//		fmt.Println(<-bobChan)
//	}

}
