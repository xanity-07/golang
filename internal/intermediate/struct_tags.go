package intermediate

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // ? XML struct tags = `xml:"first"`
// type person struct {
// 	FirstName string `json:"first_name"` //? store field in database = db:"first_name"
// 	LastName  string `json:"last_name"`  //? Ommit if empty or 0 value = `json:"fieldname,omitempty"`
// 	Age       int    `json:"age"`        //? Ommit a field regardless of value = `json:"-"`
// }

// func IntroStructTags() {

// 	xan := person{
// 		FirstName: "Xanity",
// 		LastName:  "Dev",
// 		Age:       25,
// 	}
// 	jsonData, err := json.Marshal(xan)
// 	if err != nil {
// 		fmt.Println("Error marshalling JSON:", err)
// 		return
// 	}

// 	fmt.Println(string(jsonData))
// }
