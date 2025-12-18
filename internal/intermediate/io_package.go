package intermediate

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"strings"
)

func readerFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error reading from reader:", err)
	}
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writting to writer")
	}
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error closing ")
	}
}

func bufferExample() {
	var buf bytes.Buffer //? Creates memory on the stack
	buf.WriteString("Hello buffer!")
	fmt.Println(buf.String())
}

func multiReader() {
	r1 := strings.NewReader("Hello ")
	r2 := strings.NewReader("world")
	mr := io.MultiReader(r1, r2)

	//? buf is a pointer to bytes.Buffer not a type (allocates memory in heap)
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln("Error reading from multi-reader:", err)
	}
	fmt.Println(buf.String())
}

func PipeExample() {
	pr, pw := io.Pipe()
	go func() {
		_, err := pw.Write([]byte("Hello Pipe!"))
		if err != nil {
			log.Fatalln("Error writting to pipe:", err)
		}
		pw.Close()
	}()
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(pr)
	if err != nil {
		log.Fatalln("Error reading:", err)
	}
	fmt.Println(buf.String())
}

func writeToFile(filepath string, data string) {
	//* Create file
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer closeResource(file)

	_, err = file.Write([]byte(data))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	//* Write to the file with type conversion
	// writer := io.Writer(file)
	// _, err = writer.Write([]byte(data))
	// if err != nil {
	// 	fmt.Println("Error writting file:", err)
	// 	return
	// }
}

type myResource struct {
	Name string
}

// io.Close method
func (m myResource) Close() error {
	fmt.Println(m.Name)
	return nil
}

func IntroIOPackage() {
	fmt.Println("=== Read From Reader ===")
	readerFromReader(strings.NewReader("reading from reader"))

	fmt.Println("=== Write From Writer")
	var writer bytes.Buffer
	writeToWriter(&writer, "Something to write")
	fmt.Println(writer.String())

	fmt.Println("=== Buffer Example ===")
	bufferExample()

	fmt.Println("=== Multi-Reader ===")
	multiReader()

	fmt.Println("=== Pipe Example ===")
	PipeExample()

	filepath := "io.txt"
	writeToFile(filepath, "added some text here !\n")

	resource := myResource{
		Name: "something",
	}
	closeResource(resource)

}
