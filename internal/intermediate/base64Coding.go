package intermediate

import (
	"encoding/base64"
	"fmt"
)

func IntroEncoding() {

	data := []byte("Hello, Base64 Encoding")

	//* Encode to Base64
	encode := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encode)

	//* Decode from Base64
	decode, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println(string(decode))

	//* URL safe, avoid '/' and '+'
	//* This makes the url safe by removing / + and anything
	//* that would make the url not be safe
	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println(urlSafeEncoded)
}
