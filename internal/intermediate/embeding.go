package intermediate

// import "fmt"

// // * Struct embeding allows a struct to inherit fields and methods
// // * From another struct type

// type person struct {
// 	name string
// 	age  int
// }

// type employee struct {
// 	person
// 	empID  string
// 	salary int
// }

// func (p person) introduce() {
// 	fmt.Printf("Hi my name is %s and im %d years old.", p.name, p.age)
// }

// func IntroEmbeding() {
// 	emp := employee{
// 		empID:  "emp-321",
// 		salary: 40_000,
// 		person: person{
// 			name: "xanity",
// 			age:  25,
// 		},
// 	}
// 	fmt.Println(emp)
// 	emp.introduce()
// }
