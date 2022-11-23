// package main

// import (
// 	"fmt"
// 	"strings"
// 	"sync"
// )

/*
func main() {
	str := "helloBob$helloALice#howareyou#fine,howareyou?$^"

	var aliceCount, bobCount = strings.Count(str, "$"), strings.Count(str, "#")

	var aliceChan, bobChan = make(chan string, aliceCount), make(chan string, bobCount)

	var message string

	var wg sync.WaitGroup

	var mu sync.Mutex

Loop:
	for _, ch := range str {
		switch ch {
		case '$':
			wg.Add(1)
			go func() {
				defer wg.Done()
				mu.Lock()
				fmt.Println("$ : ", message)
				aliceChan <- message
				// message = ""
				mu.Unlock()
			}()

		case '#':
			wg.Add(1)
			go func() {
				defer wg.Done()
				mu.Lock()
				fmt.Println("# : ", message)

				bobChan <- message
				message = ""
				mu.Unlock()
			}()
		case '^':
			break Loop
		default:
			mu.Lock()
			message = message + string(ch)
			mu.Unlock()
		}

	}
	wg.Wait()
	close(aliceChan)
	close(bobChan)
	// wg.Wait()
	fmt.Println(message)
	for i := 0; i < aliceCount+bobCount; i++ {
		// fmt.Print(i)
		select {
		case aliceMsg := <-aliceChan:
			fmt.Println(i, aliceMsg)

		case bobMsg := <-bobChan:
			fmt.Println(i, bobMsg)
		}
	}

}

/*
func main() {
	str := "helloBob$helloALice#howareyou#fine,how are you$^"

	var aliceCount, bobCount = strings.Count(str, "$"), strings.Count(str, "#")

	var aliceChan, bobChan = make(chan string, aliceCount), make(chan string, bobCount)

	var message string

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

	Loop:
		for _, ch := range str {
			switch ch {
			case '$':
				fmt.Println("$ : ", message)
				aliceChan <- message
				message = ""
			case '#':
				fmt.Println("# : ", message)
				bobChan <- message
				message = ""
			case '^':
				break Loop
			default:
				message = message + string(ch)
			}

		}
	}()
	wg.Wait()
	close(aliceChan)
	close(bobChan)
	// wg.Wait()
	fmt.Println(message)
	for i := 0; i < aliceCount+bobCount; i++ {
		// fmt.Print(i)
		select {
		case aliceMsg := <-aliceChan:
			fmt.Println(i, aliceMsg)

		case bobMsg := <-bobChan:
			fmt.Println(i, bobMsg)
		}
	}
	// wg.Wait()
}
*/
/*
func main() {
	str := "helloBob$helloALice#howareyou#fine,howareyou?$^"

	var aliceCount, bobCount = strings.Count(str, "$"), strings.Count(str, "#")

	var aliceChan, bobChan = make(chan string, aliceCount), make(chan string, bobCount)

	var message string

	var wg sync.WaitGroup

	var mu sync.Mutex
	// Loop:
	// for i:=0;i<(aliceCount+bobCount);i++{
	// 	aliceIndex:=strings.Index(str,"$")
	// 	bobIndex:=strings.Index(str,"#")
	// 	endIndex:=strings.Index(str,"^")
	// 	if aliceIndex<bobIndex{

	// 	}
	// }
Loop:
	for _, ch := range str {
		switch ch {
		case '$':
			wg.Add(1)
			go func() {
				defer wg.Done()
				mu.Lock()
				fmt.Println("$ : ", message)
				aliceChan <- message
				// message = ""
				mu.Unlock()
			}()

		case '#':
			wg.Add(1)
			go func() {
				defer wg.Done()
				mu.Lock()
				fmt.Println("# : ", message)

				bobChan <- message
				message = ""
				mu.Unlock()
			}()
		case '^':
			break Loop
		default:
			mu.Lock()
			message = message + string(ch)
			mu.Unlock()
		}

	}
	wg.Wait()
	close(aliceChan)
	close(bobChan)
	// wg.Wait()
	fmt.Println(message)
	for i := 0; i < aliceCount+bobCount; i++ {
		// fmt.Print(i)
		select {
		case aliceMsg := <-aliceChan:
			fmt.Println(i, aliceMsg)

		case bobMsg := <-bobChan:
			fmt.Println(i, bobMsg)
		}
	}

}
*/
/*
package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	str := "helloBob$helloALice#howareyou#fine,how are you$^"

	var aliceCount, bobCount = strings.Count(str, "$"), strings.Count(str, "#")

	var aliceChan, bobChan = make(chan string), make(chan string)

	var message string

	w := sync.WaitGroup{}

	w.Add(1)
	go func() {
		count := aliceCount + bobCount

		for {
			select {
			case aliceVal := <-aliceChan:
				fmt.Println(aliceVal)
				count = count - 1
			case bobVal := <-bobChan:
				fmt.Println(bobVal)
				count = count - 1

			}

			if count == 0 {
				break
			}
		}
		w.Done()
	}()

	for _, ch := range str {
		switch ch {
		case '$':
			// fmt.Println("$ : ", message)
			aliceChan <- message
			message = ""
		case '#':
			// fmt.Println("# : ", message)
			bobChan <- message
			message = ""
		case '^':
			break
		default:
			message = message + string(ch)
		}

	}

	w.Wait()
}
*/
// package main

// import (
// 	"fmt"
// 	"strings"
// 	"sync"
// )

// func main() {
// 	str := "helloBob$helloALice#howareyou#fine,how are you$^"

// 	var aliceCount, bobCount = strings.Count(str, "$"), strings.Count(str, "#")

// 	var aliceChan, bobChan = make(chan string), make(chan string)

// 	var message string

// 	w := sync.WaitGroup{}

// 	w.Add(1)
// 	go func() {

// 		for {
// 			count := aliceCount + bobCount
// 			select {
// 			case aliceVal := <-aliceChan:
// 				fmt.Println(aliceVal)
// 				count = count - 1
// 			case bobVal := <-bobChan:
// 				fmt.Println(bobVal)
// 				count = count - 1

// 			}
// 			if count == 0 {
// 				break
// 			}
// 		}
// 		w.Done()
// 	}()

// 	for _, ch := range str {
// 		switch ch {
// 		case '$':
// 			// fmt.Println("$ : ", message)
// 			aliceChan <- message
// 			message = ""
// 		case '#':
// 			// fmt.Println("# : ", message)
// 			bobChan <- message
// 			message = ""
// 		case '^':
// 			break
// 		default:
// 			message = message + string(ch)
// 		}

// 	}

// 	w.Wait()
// }
package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	str := "helloBob$helloALice#howareyou#fine,how are you$^"
	str = "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^dadad#dsd$"
	str = "heyalice#heybob$$$howareyou#"

	var aliceCount, bobCount = strings.Count(str[:strings.Index(str, "^")], "$"), strings.Count(str[:strings.Index(str, "^")], "#")

	var aliceChan, bobChan = make(chan string), make(chan string)

	var message string

	w := sync.WaitGroup{}

	w.Add(1)
	go func() {
		count := aliceCount + bobCount

		for {
			select {
			case aliceVal := <-aliceChan:
				fmt.Println(aliceVal)
				count = count - 1
			case bobVal := <-bobChan:
				fmt.Println(bobVal)
				count = count - 1

			}

			if count == 0 {
				break
			}
		}
		w.Done()
	}()

Loop:
	for _, ch := range str {
		switch ch {
		case '$':
			// fmt.Println("$ : ", message)
			aliceChan <- message
			message = ""
		case '#':
			// fmt.Println("# : ", message)
			bobChan <- message
			message = ""
		case '^':
			break Loop
		default:
			message = message + string(ch)
		}

	}

	w.Wait()
}
