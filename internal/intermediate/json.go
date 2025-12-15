package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

// * Struct with json fields
// * Provide meta data about the fields
// * As well as interact with databases
// * Tag will always be `json:"fieldname"`
// * Struct tags can be used for tags that can specify
// * column names or primary keys, etc. when
// * using database ORM
// * `db: user_id`
// * Need exported fields for marshaling
// * omitempty `json:"field_name,omitempty"`
// type person struct {
// 	Name    string `json:"name"`
// 	Age     int    `json:"age,omitempty"`
// 	Email   string `json:"email"`
// 	Address address
// }

type address struct {
	City  string
	State string
}

func IntroJson() {
	// * Marshalling JSON
	// xan := person{
	// 	Name: "xanity",
	// 	Age:  25,
	// }

	// jsonData, err := json.Marshal(xan)
	// if err != nil {
	// 	fmt.Println("Error marshaling json:", err)
	// 	return
	// }
	// fmt.Println(string(jsonData))

	// xan2 := person{
	// 	Name: "dsa",
	// 	Age:  24,
	// 	Address: address{
	// 		City:  "smodmn",
	// 		State: "fafas",
	// 	},
	// }
	// jsonData1, err1 := json.Marshal(xan2)
	// if err1 != nil {
	// 	fmt.Println("Error marshaling json:", err1)
	// 	return
	// }
	// fmt.Println(string(jsonData1))

	//* Unmarshalling JSON

	jsonData := `{"name": "xan1ty", "age": 25, "emp_id": "0009", "address": {"City": "Paterson", "State":"NJ"}}`

	var employee Employee
	err := json.Unmarshal([]byte(jsonData), &employee)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Printf("%v\n", employee)

	listCityState := []address{
		{City: "Paterson", State: "NJ"},
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
		{City: "Las Vegas", State: "NV"},
		{City: "Modesto", State: "CA"},
	}

	jsonList, err := json.Marshal(listCityState)
	if err != nil {
		log.Fatalln("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonList))

	// * Handling unknow JSON structs / data
	jsonD := `{"name": "John", "age": 30, "address": {"city": "Paterson", "State": "NJ"}}`
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonD), &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println("Unknown JSON:", string(jsonD))
}

type Employee struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	EmpID   string `json:"emp_id"`
	Address address
}
