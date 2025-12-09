package intermediate

import (
	"bufio"
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
	// data := make([]byte, 1024)
	// _, err = file.Read(data)
	// if err != nil {
	// 	fmt.Println("Error Reading file:", err)
	// 	return
	// }
	// fmt.Println(string(data))

	//* Another way of reading using bufio
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error scanning file.")
	}
}
