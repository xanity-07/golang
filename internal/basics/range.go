package basics

import (
	"fmt"
)

func RangeKeyword() {
	message := "Hello world!"

	for _, v := range message {
		fmt.Printf("%c\n", v)
	}
}
