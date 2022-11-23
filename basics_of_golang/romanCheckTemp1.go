
package main



import (
	"fmt"
	//"strings"
)



func main() {
	str:=readStr()
	number:=romanToInt(str)
	fmt.Println(number)
}

func readStr()(str string){
	fmt.Scanln(&str)
	return
}

func romanToInt(s string) int{

	var RomanNumerals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
	}
	sum := 0
	greatest := 0
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]
		num := RomanNumerals[rune(letter)]
		fmt.Printf("%c %d %d %d\n",letter,num,greatest,sum)
		if num >= greatest {
			greatest = num
			sum = sum + num
			continue
		}
		sum = sum - num
	}
	return sum

}
