
package main



import (

   "fmt"

)



func main() {
	 on := 1
	  fmt.Println(on == true)

   switch {

   case true:

       fmt.Println("In true case")

   case 1 < 2:

       fmt.Println("In conditional operation check case")

   case false:

       fmt.Println("In false case")

   default:

       fmt.Println("In default case")

   }

}
