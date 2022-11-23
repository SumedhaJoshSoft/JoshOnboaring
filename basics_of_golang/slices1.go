
package main



import "fmt"



func main() {



   arr := [5]string{"mon", "tue", "wed", "thu", "fri"}

   slc1 := arr[:]

   slc2 := slc1



   slc1[2] = "sat"

   fmt.Println(slc1)



   slc2[4] = "sun"

   fmt.Println(slc2)



}
