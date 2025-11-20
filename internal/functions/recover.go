package functions

import "fmt"

func IntroToRecover() {
	processRecover()
	fmt.Println("Returned from processRecover.")
}

func processRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

	fmt.Println("Start Process.")
	panic("Something went wrong.")
	//! fmt.Println("End Process.")
}
