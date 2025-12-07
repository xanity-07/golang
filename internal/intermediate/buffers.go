package intermediate

import (
	"bufio"
	"fmt"

	// "fmt"
	"os"
	// "strings"
)

func IntroBufio() {
	// reader := bufio.NewReader(strings.NewReader("Hello, bufio packageeee!\n"))

	// data := make([]byte, 20)
	// n, err := reader.Read(data)
	// if err != nil {
	// 	fmt.Println("Error reading data", err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	// line, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading string", err)
	// 	return
	// }
	// fmt.Println(line)

	//* Writing byte slice
	writer := bufio.NewWriter(os.Stdout)
	data := []byte("Hello writing to console.!\n")

	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	//* Flush the buffers to ensure all the data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing.")
		return
	}

	//* Writing string
	str := "This is a string.\n"
	s, err := writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string.")
		return
	}
	fmt.Printf("Wrote %d bytes\n", s)

	//* Flush the buffers
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing")
		return
	}
}
