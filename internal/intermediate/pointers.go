package intermediate

import "fmt"

func IntroPointers() {
	var ptr *int //? Nil value
	var a int = 10
	ptr = &a //? Now its int because its initialized to the refrence of a
	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr) //? Derefrencing to get the value
}
