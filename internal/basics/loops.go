package basics

import "fmt"

func Loops() {
	//* Iteration
	// for i := 0; i < 6; i++ {
	// 	fmt.Println(i)
	// }

	//* Iteration over a collection
	// nums := []int{1, 2, 3}
	// for i, v := range nums {
	// 	fmt.Printf("Index: %d Value: %d\n", i, v)
	// }

	//* Iteration excluding index
	// for _, v := range nums {
	// 	fmt.Printf("Value: %d\n", v)
	// }

	//* Break and Continue Statements
	//? Break Statement
	// for i := 0; i < 6; i++ {
	// 	if i == 3 {
	// 		fmt.Printf("Reached the max steps: %d\n", 3)
	// 		break
	// 	}
	// 	fmt.Printf("Current index: %d. %d Steps remaining\n", i, 3-i)
	// }

	//? Continue Statement
	// for i := 0; i < 6; i++ {
	// 	if i == 3 {
	// 		fmt.Println("Skipped")
	// 		continue
	// 	}
	// 	fmt.Printf("Current index: %d\n", i)
	// }

	//* For Loop Exercise + Nested For Loops
	rows := 5
	for i := 1; i <= rows; i++ {
		//? Spaces
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		//? Stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		//? New Line
		fmt.Println()
	}
}
