package intermediate

import (
	"fmt"
	"os"
)

func IntroWritingFiles() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	defer file.Close()

	data := []byte("Writing some text in here.\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
