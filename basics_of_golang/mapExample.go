package main

 

import "fmt"

 

func main() {

   var employeeSalary map[string]int

   AddEmployeeSalary("John", 40000, employeeSalary)

   fmt.Printf("salary of john is %v \n", employeeSalary["john"])

}

 

func AddEmployeeSalary(name string, salary int, employeeSalary map[string]int) {

   employeeSalary[name] = salary

}


