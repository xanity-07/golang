package api

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"
// )

// // ? To create a HTTP server you need to define handlers they are functions
// // ? Functions that handle the business logic for an endpoints
// // ? Then we start the server to listen on an specific PORT

// func IntroHTTPServer() {
// 	// ? Create server address
// 	const port string = ":8080"
// 	server := http.Server{
// 		Addr:         port,
// 		WriteTimeout: 10 * time.Millisecond,
// 		ReadTimeout:  10 * time.Millisecond,
// 		IdleTimeout:  30 * time.Millisecond,
// 	}
// 	// ? Root route is "/"
// 	http.HandleFunc("/", handlerFunc)
// 	fmt.Println("Listening on port", port)

// 	// ? Start Server
// 	err := server.ListenAndServe()
// 	if err != nil {
// 		log.Fatalln("error starting server:", err)
// 	}
// }

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, listening on port:", r.URL.Port())
// }
