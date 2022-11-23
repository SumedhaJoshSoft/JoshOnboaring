
package main



import "fmt"



type salaryAccount interface {

   account

   credit(money float64) (message string)

}



type account interface {

   deposit(money float64) (message string)

}



type hdfc struct{}



func (hdfc *hdfc) credit(money float64) (message string) {
	fmt.Println(hdfc)
   return fmt.Sprintf("%v INR has been created into the account", money)

}




func main() {

   acc := &hdfc{}

   msg := acc.credit(23.45)

   fmt.Println(msg)

}


