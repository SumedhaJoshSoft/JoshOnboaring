

 package main

 import (
 "fmt"
 "math"
 )

 func calculateArea(r float64) (area float64, err error) {
 if r < 0 {
 err = fmt.Errorf("radius can not be negative")
 return
 }

 area = r * r * math.Pi
 return
 }
 func main() {
 area := calculateArea(12)
 fmt.Println("Area of circle is ", area)
 }
