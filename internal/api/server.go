package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Person Struct
type Person struct {
	Name string
	Age  int
}

// Sample Data
var personData = map[string]Person{
	"1": {Name: "John Doe", Age: 33},
	"2": {Name: "Jane Doe", Age: 24},
	"3": {Name: "Jack Doe", Age: 28},
}

// Handler Fcuntion for Endpoint
func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from url query params
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}

	// Check if person exists
	person, exists := personData[id]
	if !exists {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	// Set response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the person data to JSON and write to the response
	if err := json.NewEncoder(w).Encode(person); err != nil {
		http.Error(w, "Failed to enconde response", http.StatusInternalServerError)
		return
	}
}

func Start() {
	// Define port
	const port = ":8080"

	// Server struct
	server := http.Server{
		Addr:         port,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// Print the confirmation message
	fmt.Printf("Server started on port %s\n", port)

	// Set up the endpoint and handler
	http.HandleFunc("/person", getPersonHandler)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("error starting server:", err)
	}
}
