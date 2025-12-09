package intermediate

import (
	"fmt"
	"path/filepath"
)

func IntroFilePath() {
	joinedPath := filepath.Join("home", "desktop", "text.txt")
	fmt.Println(joinedPath)

	dir, file := filepath.Split("home/desktop/text.txt")
	fmt.Println("File:", file)
	fmt.Println("Dir:", dir)

}
