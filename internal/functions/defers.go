package functions

import "fmt"

func IntroToDefer() {
	process(1)
}

func process(i int) {

	defer fmt.Println("Defered i value:", i)
	defer fmt.Println("1st Defered Statement")
	defer fmt.Println("2nd Defered Statement")
	defer fmt.Println("3rd Defered Statement")
	i++
	fmt.Println("Normal Execution Statement")
	fmt.Println("Value of i:", i)

}
