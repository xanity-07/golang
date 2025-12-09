package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IntroLineFilter() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	keyword := "to"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			fmt.Println("Filtered line:", line)
			return
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error scanning:", err)
		return
	}
}
