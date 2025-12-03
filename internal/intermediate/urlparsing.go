package intermediate

import (
	"fmt"
	"net/url"
)

func IntroUrl() {
	// rawUrl := "https://example.com:8080/path?query=param#fragment"

	// parsedUrl, err := url.Parse(rawUrl)
	// if err != nil {
	// 	fmt.Println("Error parsing the URL:", err)
	// 	return
	// }
	// fmt.Println("Scheme:", parsedUrl.Scheme)
	// fmt.Println("Host:", parsedUrl.Host)
	// fmt.Println("Port:", parsedUrl.Port())
	// fmt.Println("Path:", parsedUrl.Path)
	// fmt.Println("Query:", parsedUrl.RawQuery)
	// fmt.Println("Fragment:", parsedUrl.Fragment)

	// rawOtherUrl := "https://example.com/path?name=john&age=30"
	// parsedOtherUrl, err := url.Parse(rawOtherUrl)
	// if err != nil {
	// 	fmt.Println("Error Parsing Url:", err)
	// 	return
	// }
	// queryParams := parsedOtherUrl.Query()
	// fmt.Println(queryParams)
	// fmt.Println(queryParams.Get("name"))
	// fmt.Println(queryParams.Get("age"))

	//* Building a URL
	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := baseUrl.Query()
	fmt.Println(query)
	query.Set("name", "John")
	fmt.Println(query)
	baseUrl.RawQuery = query.Encode()

	fmt.Println("Built URL:", baseUrl.String())

	values := url.Values{}
	//* Add key value pairs to the values object
	values.Add("name", "jane")
	values.Add("age", "30")
	values.Add("city", "London")
	values.Add("country", "UK")

	//* Encode
	encodedQuery := values.Encode()
	fmt.Println("Encoded Query:", encodedQuery)

	//* Build a URL
	newBaseUrl := "https://example.com/search"
	fullNewBaseUrl := newBaseUrl + "?" + encodedQuery
	fmt.Println(fullNewBaseUrl)
}
