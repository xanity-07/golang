package basics

import "fmt"

func IntroToMaps() {
	//* Declaring Map
	// var x map[string]string //* Nil map
	// fmt.Println(x)

	//* Declaring Map Using Make
	// z := make(map[string]int)
	// z["age"] = 25
	// z["code"] = 7
	// z["lol"] = 3

	// fmt.Println(z)

	//* Declaring Map Literal
	// x = map[string]string{
	// 	"name": "Xanity",
	// }

	//* Delete Key
	// delete(z, "lol")
	// fmt.Println(z)

	//* Delete all
	// clear(z)
	// fmt.Println(z)

	//* Inline declaration/asignment
	y := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println(y)

	//* Iteration of keys
	for k := range y {
		fmt.Println(k)
	}

	//* Iteration of value
	for _, v := range y {
		fmt.Println(v)
	}

	//* Iteration of key and value
	for k, v := range y {
		fmt.Printf("Key: %s Value: %d\n", k, v)
	}

	//* Does a value exist?
	_, ok := y["a"]
	fmt.Println(ok)

	fmt.Println("Map length:", len(y))

	//* Nested Maps
	nestedMap := make(map[string]map[string]int)
	nestedMap["m1"] = y
	fmt.Println(nestedMap)
}
