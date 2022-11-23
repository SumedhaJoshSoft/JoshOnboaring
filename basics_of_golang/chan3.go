
package main



import (
	"fmt"
	"time"
)



func main() {

    var ch chan int

    var count int



    go func() {

        ch <- 1

    }()



    time.Sleep(2*time.Second)
    count = <-ch



    fmt.Println(count)

}
