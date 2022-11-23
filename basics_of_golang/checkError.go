/*
package main



import "fmt"



func main() {

   s := sum(10, 20)

   if s != nil {

       fmt.Println(*s)

   }

}



func sum(num1, num2 int) (sum *int) {

   s := num1 + num2

   return &s

}

*/
/*
package main



import "fmt"



func main() {

   str := "Hello"

   for s := range str {

       fmt.Printf("%v", s)

   }

}

*/

/*
package main

 

import (

   "fmt"

)

 

func main() {

   s := "GO"

   for _, val := range s {

       fmt.Println(val)

   }

}
*/

package main

 

import "fmt"

 

func main() {

   done := false

   if done := true; done {

       fmt.Println(done)

   }

 

   fmt.Println(done)

}
