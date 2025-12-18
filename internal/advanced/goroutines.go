package advanced

import (
	"fmt"
	"time"
)

//* Goroutines are basically like async await and promises in JS

func printNumbers(input int) {
	for i := range input {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetter(letters string) {
	for _, letter := range letters {
		fmt.Println(string(letter))
		time.Sleep(200 * time.Millisecond)
	}
}

func IntroGoroutines() {
	fmt.Println("Hello from xanity")
	go printNumbers(8)
	go printLetter("hello")
	time.Sleep(2 * time.Second)
}
