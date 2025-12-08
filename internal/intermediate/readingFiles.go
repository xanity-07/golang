package intermediate

import (
	"fmt"
	"os"
)

func IntroReadingFiles() {
	//* Open the file
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func() {
		fmt.Println("Closing open file.")
		file.Close()
	}()

	//* Read the content of the file
	data := make([]byte, 1024)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error Reading file:", err)
		return
	}
	fmt.Println(string(data))
}
