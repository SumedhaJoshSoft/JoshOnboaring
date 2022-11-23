
package main



import (

    "fmt"

)



func test(n int, c chan int) {

    x, y := 0, 1

    for i := 0; i < n; i++ {

        c <- x
	x,y = y,x+y
    }



}



func main() {

    c := make(chan int, 2)

    go test(cap(c), c)

    for i := range c {

        fmt.Println(i)

    }
    close(c)

}
