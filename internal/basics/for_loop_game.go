package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func GuessNumber() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	target := random.Intn(10) + 1

	fmt.Println("Welcome to the guessing game.")
	fmt.Println("Guess a number between 1 and 10")
	fmt.Println("Can you guess what it is?")

	var guess int

	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Congrats!!! You guess the correct number!")
			break
		} else if guess < target {
			fmt.Println("Try a bigger number :)")
		} else {
			fmt.Println("Try a smaller number :)")
		}
	}
}
