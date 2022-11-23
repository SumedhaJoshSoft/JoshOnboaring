package main
import "fmt"

func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func dummy(){
	fmt.Println("Dummy Func")
}
func b() {
    fmt.Println("before b's differ")
    fmt.Println("in b")
    a()
    
    defer un(trace("b"))
    defer dummy()
}

func main() {
    b()
}
