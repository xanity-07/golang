package app

import (
	"log"
	"os"
)

type Applicatio struct {
	Logger *log.Logger
}

func NewApplication() (*Applicatio, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Applicatio{
		Logger: logger,
	}

	return app, nil
}
