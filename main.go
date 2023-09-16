package main

// This line indicates that this file belongs to the main package.
// In Go, a package is a way to organize and reuse code.
// The main package is a special package used for creating executable programs.

import (
	"fmt"      // "fmt": This package provides functions for formatted input and output.
	"net/http" // "net/http": This package is part of Go's standard library and is used for
	// building HTTP servers and clients.
	"github.com/LaureneT/go_rest_api/api"
)

// // ProjectJSON represents the JSON structure for projects.
// type ProjectJSON struct {
// 	Name string `json:"name"`
// }

func main() {
	// Set up the HTTP server
	server := &http.Server{
		Addr: ":8080",
	}

	// Set up the route and inject the RealReadmeGetter instance
	http.HandleFunc("/readme", func(serverResponse http.ResponseWriter, clientRequest *http.Request) {
		api.HandleReadme(serverResponse, clientRequest)
	})
	http.HandleFunc("/projects", func(serverResponse http.ResponseWriter, clientRequest *http.Request) {
		api.HandleProjects(serverResponse, clientRequest)
	})
	http.HandleFunc("/hello", api.HandleHelloWorld)

	fmt.Println("Server started on", server.Addr)

	// Start the server and handle errors
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
