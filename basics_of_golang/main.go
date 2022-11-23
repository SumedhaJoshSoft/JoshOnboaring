package main

import (
	"fmt"
	"math"
)


func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return (productionRate*int(successRate)*60)/100
}


func sum(b int,a ...float64){
	fmt.Println(a,b)
}

func dummyFunc(a int)(c int){
	a=10
	fmt.Println(a)
	c =func(x int) int{
		x=x+10
		fmt.Println(x)
		return x
	}(a)
	fmt.Println(a,c)
	return
}

func dummy(a int)(a1 int){
	if a<10{
		a1=100
	}else{
		a1=-100
	}
	return
}

func dummyF(a float64)(float64){
	return a+10.0
}

func doNothing(a int){
	fmt.Println(a)
}

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }


// Use them like this:
func deferCheck() {
    trace("a")
    defer untrace("a")
    fmt.Println("Do something here")
    // do something....
}

func main(){
	fmt.Println("Hello world...!")
	knightIsAwake:=   false
	archerIsAwake:=   false
	prisonerIsAwake:= true
	petDogIsPresent:=    false
		//expected:        true,
	fmt.Println((prisonerIsAwake && (!knightIsAwake && !prisonerIsAwake)) || (petDogIsPresent && !archerIsAwake))
	fmt.Println(!knightIsAwake && !prisonerIsAwake)
	fmt.Println(petDogIsPresent && !archerIsAwake)
	fmt.Println(((1105*90)/100)/60)
	fmt.Println(1547*90/100)
	fmt.Println(1547/10)
	fmt.Println(1547%10)


    var carsCount uint = 100
    var groupedCarsCount uint = uint(carsCount)/10
    var singleCarsCount uint = uint(carsCount)/10
    fmt.Println((groupedCarsCount*95000)+(singleCarsCount*10000),groupedCarsCount,singleCarsCount)
    fmt.Println(CalculateWorkingCarsPerMinute(221,100))
    fmt.Println(CalculateWorkingCarsPerMinute(1105,90))
    fmt.Println(CalculateWorkingCarsPerMinute(221,100))
    fmt.Println(CalculateWorkingCarsPerMinute(221,100))
    fmt.Println(CalculateWorkingCarsPerMinute(221,100))

    fmt.Println(math.Floor(65.3333*100)/100)
    fmt.Println(math.Ceil(65.3000*100)/100)

    a:=10
    for a>=0{
    	fmt.Println(a)
	a--
    }

    sum(1,2,3)

    fmt.Println(dummyFunc(10))

b:=10
    for{
	fmt.Println(b)
	if b<20{
		b++
	}else{
		break
	}
    }


    var s1 [5]int
    for idx,val:=range s1{
	s1[idx]=3
    	fmt.Println(idx,val,s1[idx])
    }
    a1:=101
    fmt.Println(dummy(9))
    fmt.Println(dummy(10))
    fmt.Println(a1)
    doNothing(11)
    a1,b1 :=10,23
    b1,c1:=21,22
    fmt.Println(a1,b1,c1)
    //b1:=11
    //fmt.Println(a1,b1,c1)

    var t interface{}
//t = dummy(a1)
t = dummyF(1.1)
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
deferCheck()

}
