package intermediate

import (
	"log"
	"os"
)

func IntroLogging() {

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}

	defer file.Close()

	infoLogger := log.New(file, "INFO: ", log.Ltime|log.Ldate|log.Lshortfile)
	warnLogger := log.New(file, "WARN: ", log.Ltime|log.Ldate|log.Lshortfile)
	errorLogger := log.New(file, "ERROR: ", log.Ltime|log.Ldate|log.Lshortfile)
	debugLogger := log.New(file, "DEBUG: ", log.Ltime|log.Ldate|log.Lshortfile)

	infoLogger.Println("This is a info message.")
	warnLogger.Println("This is a warm message.")
	errorLogger.Println("This is a error message.")
	debugLogger.Println("This is a debug message.")
}
