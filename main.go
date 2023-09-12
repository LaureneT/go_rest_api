package main

// This line indicates that this file belongs to the main package.
// In Go, a package is a way to organize and reuse code.
// The main package is a special package used for creating executable programs.

import (
	"fmt"
	"net/http"
)

// "fmt": This package provides functions for formatted input and output.
// "net/http": This package is part of Go's standard library and is used for
// building HTTP servers and clients.

func handleProjects(serverResponse http.ResponseWriter, clientRequest *http.Request) {
	// If the path is not /projects, it responds with an HTTP 404 "Not Found" status by calling http.NotFound.
	if clientRequest.URL.Path != "/projects" {
		http.NotFound(serverResponse, clientRequest)
		return
	}

	fmt.Fprintln(serverResponse, "Hello, this is your /projects endpoint!")
}

func main() {
	//This line sets up an HTTP route. It tells the web server that when a request is made 
	// to the path /projects, it should call the handleProjects function to handle that request.
	http.HandleFunc("/projects", handleProjects)

	fmt.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
