
package main



import "fmt"



type movie struct {

   title, genre string

   rating       float64

   released     bool

}



func main() {

   var lineKing movie

   fmt.Printf("title: %v, rating: %v, released: %v", lineKing.title, lineKing.rating, lineKing.released)

}
