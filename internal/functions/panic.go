package functions

import "fmt"

func IntroToPanic() {
	//* Syntax
	//? panic(interface{})
	//* Valid input
	inputs(10)

	//* Non valid input
	inputs(-10)
}

func inputs(input int) {

	defer fmt.Println("Defered 1")
	defer fmt.Println("Defered 2")

	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be a non-negative type")
	}
	fmt.Println("Processing input:", input)

}
