
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
	sum:=0
	/*Loop:
	for index,val:=range s{
		switch val{
		case 'I':	
			if index<=len(s)-1{
				switch s[index+1]{
					case 'V':
						sum = sum+4
						continue Loop
					case 'X':
						sum = sum+9
						continue Loop
				}
				fmt.Println("KK : ", index,val,len(s)-1)
			}
			sum = sum +1
		case 'V':
			sum = sum + 5 
		case 'X':
			sum = sum + 10 
		case 'L':
			sum = sum + 50
		case 'C':
			sum = sum +100
		case 'D':
			sum = sum+500
		case 'M':
			sum = sum + 1000
		}
		fmt.Printf("%d %c %d\n",index,val,len(s)-index-1)
	}*/
	index:=0
	Loop:
	for index<len(s){
		switch s[index]{
                case 'I':
                        if index<len(s)-1{
                                switch s[index+1]{
                                        case 'V':
                                                sum = sum+4
                                                index=index+2
						continue Loop
                                        case 'X':
                                                sum = sum+9
						index=index+2
                                                continue Loop
                                }
                        }
                        sum = sum +1
                case 'V':
                        sum = sum + 5
                case 'X':
			if index<len(s)-1{
                                switch s[index+1]{
                                        case 'L':
                                                sum = sum+40
                                                index=index+2
                                                continue Loop
                                        case 'C':
                                                sum = sum+90
                                                index=index+2
                                                continue Loop
                                }
                        }
                        sum = sum + 10
                case 'L':
                        sum = sum + 50
                case 'C':
			if index<len(s)-1{
                                switch s[index+1]{
                                        case 'D':
                                                sum = sum+400
                                                index=index+2
                                                continue Loop
                                        case 'M':
                                                sum = sum+900
                                                index=index+2
                                                continue Loop
                                }
                        }
                        sum = sum +100
                case 'D':
                        sum = sum+500
                case 'M':
                        sum = sum + 1000
                }
		fmt.Println(s[index:])
		index=index+1


	}
	return sum

}
