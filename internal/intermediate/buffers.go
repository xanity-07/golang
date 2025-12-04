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

	writer := bufio.NewWriter(os.Stdout)

	//* Writinf byte slice
	data := []byte("Hello, bufio package! \n")

	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error writing to data", err)
		return
	}
	fmt.Printf("Wrote %d bytes.\n", n)

	//* Have to flush the buffer.
	//? It does not automatically flush the writer.
	//* Flush the buffer to ensure all data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
	}

	//* Writing string
	str := "This is some random string\n"
	s, err := writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes.\n", s)

	//* Flush buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer on line 57:", err)
		return
	}
}
