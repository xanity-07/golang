package intermediate

import (
	"encoding/xml"
	"fmt"
	"log"
)

type person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	City    string   `xml:"city"`
	Email   string   `xml:"email"`
}

func IntroXML() {
	xan := person{
		Name:  "Xanity",
		Age:   25,
		City:  "Paterson",
		Email: "xan@example.com",
	}
	//* Marshalling XML
	// xmlData, err := xml.Marshal(xan)
	// if err != nil {
	// 	log.Fatalln("Error marshalling XML:", err)
	// 	return
	// }
	// fmt.Println(string(xmlData))

	//* Formating with indent
	xmlData, err := xml.MarshalIndent(xan, "", "  ")
	if err != nil {
		log.Fatalln("Error marshalling XML:", err)
	}
	fmt.Println(string(xmlData))

	//* Unmarshalling XML
	// xmlRawData := `
	// <person>
	// 	<name>John</name>
	// 	<age>25</age>
	// 	<empID>0009</empID>
	// </person>
	// `
	// var data string
	// err = xml.Unmarshal([]byte(xmlRawData), &data)
	// if err != nil {
	// 	log.Fatalln("Error unmarshalling XMLData:", err)
	// }
	// fmt.Println(xmlRawData)

	book := book{
		ISBN:   "584-75-437-587-34-32",
		Tittle: "Go Bootcamp",
		Author: "Xan",
	}

	xmlBookData, err := xml.MarshalIndent(book, "", "  ")
	if err != nil {
		log.Fatalln("Error marshalling XML:", err)
	}
	fmt.Println(string(xmlBookData))
	fmt.Println(book.Author)
}

type book struct {
	XMLName xml.Name `xml:"book"`
	ISBN    string   `xml:"isbn,attr"`
	Tittle  string   `xml:"title,attr"`
	Author  string   `xml:"author,attr"`
}
