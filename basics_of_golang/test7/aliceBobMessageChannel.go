package main

import (
	"fmt"
	"strings"
	"sync"
)

/*
func main() {
	str := "helloBob$helloALice#howareyou#fine,how are you?$^"

	var aliceCount, bobCount = strings.Count(str, "$"), strings.Count(str, "#")

	var aliceChan, bobChan = make(chan string, aliceCount), make(chan string, bobCount)

	var message string

	var wg sync.WaitGroup

Loop:
	for _, ch := range str {
		switch ch {
		case '$':
			fmt.Println("$ : ", message)
			wg.Add(1)
			go func() {
				defer wg.Done()
				aliceChan <- message
				message = ""
			}()

		case '#':
			fmt.Println("# : ", message)
			wg.Add(1)
			go func() {
				defer wg.Done()
				bobChan <- message
				message = ""
			}()
		case '^':
			break Loop
		default:
			message = message + string(ch)
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
