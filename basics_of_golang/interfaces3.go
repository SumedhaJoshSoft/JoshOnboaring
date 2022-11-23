
package main



import "fmt"



type bar interface {

   talk()

}



type foo struct {

   bar

}



func (f foo) talk() {

   fmt.Println("foo can talk...")

}



func main() {

   b := foo{}

   var f foo

   f.bar = b

   f.bar.talk()

}
