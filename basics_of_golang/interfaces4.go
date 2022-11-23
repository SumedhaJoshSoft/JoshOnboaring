
package main



import "fmt"



type parent struct {

}



type child struct {

   p parent

}



type foo interface {

   hello()

}



func (p parent) hello() {

   fmt.Println("In child")

}



func (p child) hello() {

   fmt.Println("In parent")

}



func main() {

   var f foo

   f = child{}

   f.p.hello()

}
