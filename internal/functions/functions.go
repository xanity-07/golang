package functions

import "fmt"

func PrintAll() {
	// b := add(1, 2)
	// fmt.Println(b)

	// operation := add
	// result := operation(1, 3)
	// fmt.Println(result)

	// a := applyOperation(1, 4, add)
	// fmt.Println(a)

	// multByTwo := createMultiplier(2)
	// product := multByTwo(6)
	// fmt.Println(product)
	//*Return multiple vals
	// returnVal1, returnVal2 := divide(10, 2)
	// fmt.Println(returnVal1, returnVal2)

	//* Omit one of the return vals
	// _, remainder := divide(10, 3)
	// fmt.Println(remainder)

	//* Errors With Multiple return values
	// result, err := compare(2, 2)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println(result)
	// }

	//* Variadic Function
	total := sumAll(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(total)

	//* Spread Slice as Params
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(sumAll(1, 2, s...))
}

// func add(a int, b int) int {
// 	return a + b
// }

// func applyOperation(x int, y int, operation func(int, int) int) int {
// 	return operation(x, y)
// }

//	func createMultiplier(x int) func(int) int {
//		return func(y int) int {
//			return x * y
//		}
//	}

// func divide(x int, y int) (int, int) {
// 	quotient := x / y
// 	remainder := x % y
// 	return quotient, remainder
// }

// func compare(a int, b int) (string, error) {
// 	if a > b {
// 		return "A is greater than B", nil
// 	} else if b > a {
// 		return "B is greater than A", nil
// 	} else {
// 		return "", errors.New("They are equal can't compare.")
// 	}

// }

func sumAll(a int, b int, nums ...int) int {
	sum := a + b
	for _, v := range nums {
		sum += v
	}
	return sum
}
