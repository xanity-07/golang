package practice

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func NewReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		fmt.Println("Error reading from reader:", err)
		return
	}
	fmt.Println(string(buf[:n]))
}

func NewWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		fmt.Println("Error writing to writer:", err)
		return
	}

}

func NewMultiReader() {
	r1 := strings.NewReader("this is first reader ")
	r2 := strings.NewReader("This is the seccond reader")
	mr := io.MultiReader(r1, r2)

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(mr)
	if err != nil {
		fmt.Println("Error reading multi reader:", err)
		return
	}
	fmt.Println(buf.String())
}
