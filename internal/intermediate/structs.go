package intermediate

// import "fmt"

// //* Anonymous fields make you able to just access the property names
// //* instead of doing PhoneBook.home or cell you can just do xan.cell
// //*

// // * Struct Declaration
// type Person struct {
// 	name string
// 	age  int
// 	//* Struct Embeding
// 	address Address
// 	//? Anonymous Struct
// 	PhoneBook
// }

// type Address struct {
// 	city    string
// 	country string
// }

// // * Phone struct for anonymous struct
// type PhoneBook struct {
// 	home string
// 	cell string
// }

// // * Struct method
// func (p Person) printPerson() {
// 	fmt.Printf("Name: %s Age: %d\n", p.name, p.age)
// }

// // * Struct Pointer Method
// func (p *Person) birthday() int {
// 	p.age++
// 	fmt.Printf("Happy birth day!!! you are now %d years old\n", p.age)
// 	return p.age
// }

// func IntroStructs() {
// 	//* Struct initialization
// 	xan := Person{
// 		name: "Xanity",
// 		age:  25,
// 		//? Embedding the struct as a field
// 		address: Address{
// 			city:    "Paterson",
// 			country: "US",
// 		},
// 		PhoneBook: PhoneBook{
// 			home: "1-234-5678",
// 			cell: "1-987-6543",
// 		},
// 	}

// 	fmt.Println("Xan age:", xan.age)
// 	//* Accessing method
// 	xan.printPerson()
// 	xan.birthday()

// 	//* Anonymous Struct
// 	user := struct {
// 		username string
// 		email    string
// 	}{
// 		username: "xanity01",
// 		email:    "psudoemail@example.com",
// 	}
// 	fmt.Println(user)
// }
