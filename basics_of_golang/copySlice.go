

package main



import (

   "fmt"

)



func main() {

   a1 := [5]string{"English", "Japanese", "Spanish", "French", "Hindi"}

   a2 := a1

   a2[1] = "German"

   fmt.Println("a1 = ", a1)

   fmt.Println("a2 = ", a2)

}
