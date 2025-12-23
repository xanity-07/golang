package advanced

import (
	"fmt"
	"time"
)

func IntroChannels() {
	greeting := make(chan string)
	// str := "Hello"
	alphabet := "abcdefg"
	// go func() {
	// 	greeting <- str
	// 	greeting <- "World!"
	// }()
	// go func() {
	// 	receiver := <-greeting
	// 	fmt.Println(receiver)
	// 	receiver = <-greeting
	// 	fmt.Println(receiver)
	// }()

	go func() {
		for _, e := range alphabet {
			greeting <- "Alphabet: " + string(e)
		}
	}()

	go func() {
		for range 7 {
			receiver := <-greeting
			fmt.Println(receiver)
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Program end")
}
