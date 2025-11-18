package basics

import (
	"fmt"
	"net/http"
)

func ImportsExample() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("HTTP Respnse Status: ", res.Status)
}
