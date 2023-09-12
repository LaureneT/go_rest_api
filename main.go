package main

// This line indicates that this file belongs to the main package.
// In Go, a package is a way to organize and reuse code.
// The main package is a special package used for creating executable programs.

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// "fmt": This package provides functions for formatted input and output.
// "net/http": This package is part of Go's standard library and is used for
// building HTTP servers and clients.
// "github.com/google/go-github/github" package to interact with the GitHub API.

func getREADME() (string, error) {
	// Create a context
    ctx := context.Background()
	
	// Create a GitHub client with your personal access token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	owner := "avelino"
	repo := "awesome-go"

	// Fetch the README file from the repository
	readme, _, err := client.Repositories.GetReadme(ctx, owner, repo, nil)
	if err != nil {
		return "", err
	}

	// Decode the README content
	readmeContent, err := readme.GetContent()
	if err != nil {
		return "", err
	}

	return readmeContent, nil
}

func handleHelloWorld(serverResponse http.ResponseWriter, clientRequest *http.Request) {
	// If the path is not /projects, it responds with an HTTP 404 "Not Found" status by calling http.NotFound.
	if clientRequest.URL.Path != "/projects" {
		http.NotFound(serverResponse, clientRequest)
	}
	fmt.Fprintln(serverResponse, "Hello, this is your /projects endpoint!")
}

func handleReadme(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/projects" {
		http.NotFound(w, r)
		return
	}

	// Fetch the README content from GitHub
	readmeContent, err := getREADME()
	if err != nil {
		http.Error(w, "Error fetching README", http.StatusInternalServerError)
		fmt.Println("Error fetching README:", err)
		return
	}

	// Set the response content type to plain text
	w.Header().Set("Content-Type", "text/plain")

	// Send the README content as the HTTP response
	fmt.Fprintln(w, readmeContent)
}

func main() {
	//This line sets up an HTTP route. It tells the web server that when a request is made 
	// to the path /projects, it should call the handleProjects function to handle that request.
	//http.HandleFunc("/projects", handleHelloWorld)
	http.HandleFunc("/projects", handleReadme)

	fmt.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
