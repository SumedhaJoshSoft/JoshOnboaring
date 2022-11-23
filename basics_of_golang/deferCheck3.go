
package main



import (

        "fmt"

)



func main() {

        i := 22

        defer fmt.Println("Value of i : ", i)

        i = 33

        i = 44

        return

        i = 55

}
