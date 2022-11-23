
package main



import "fmt"



type person struct {

   name string

}



func (p *person) SetName(name string) {

fmt.Println(p)
   p.name = name

}



func main() {

   p := person{}

   p.SetName("john")

   fmt.Printf("Person name is %v\n", p.name)

}


