
package main



import "fmt"



type person struct {

}


type employee struct {

   p person

}



type human interface {

   describe()

}


func(e *employee) describe(){
	fmt.Println("AAA")
}


func (p *person) describe() {

   fmt.Println("I am a person")

}


func common(h human)(){
	fmt.Println("Common")
}


func main() {

   h := &employee{}

   h.describe()
   p1:=person{}
   p1.describe()
	p:=&person{}
	p.describe()

	common(h)
	common(p)
}
