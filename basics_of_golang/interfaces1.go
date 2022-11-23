
package main



import "fmt"



type animal interface {

   talk()

}

type human interface{
	walk()
}

type girl struct{}

type cat struct{}

func (g girl)walk(){
	fmt.Println("Girl can walk")	
}

func NewCat() animal {

   return cat{}

}



func (c cat) talk() {

   fmt.Println("cat can talk")

}



type dog struct{}



func NewDog() animal {

   return dog{}

}



func (c dog) talk() {

   fmt.Println("dog can talk")

}




func main() {

   a := NewCat()
   g:=girl{}
   a.talk()
   g.walk()

   a=NewDog()
   a.talk()
}
