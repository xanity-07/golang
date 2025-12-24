package api

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// var BASE_URL string = "https://jsonplaceholder.typicode.com/posts/1"

// func IntroHTTPClient() {
// 	//* Make http client struct using http package
// 	client := http.Client{}

// 	//* GET data from api
// 	res, err := client.Get(BASE_URL)
// 	if err != nil {
// 		fmt.Println("Error using GET Request:", err)
// 		return
// 	}
// 	//* Defer and close
// 	defer res.Body.Close()

// 	//* Read from body using IO package
// 	data, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println("Error Reading from res.Body", err)
// 		return
// 	}
// 	//* Print to console
// 	fmt.Println(string(data))
// }
