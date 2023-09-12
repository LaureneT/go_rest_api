package main

// This line indicates that this file belongs to the main package.
// In Go, a package is a way to organize and reuse code.
// The main package is a special package used for creating executable programs.

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"github.com/spf13/viper"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// "fmt": This package provides functions for formatted input and output.
// "net/http": This package is part of Go's standard library and is used for
// building HTTP servers and clients.
// "github.com/google/go-github/github" package to interact with the GitHub API.

type ReadmeGetter interface {
	GetREADME() (string, error)
}

type RealReadmeGetter struct{}

func (g *RealReadmeGetter) GetREADME() (string, error) {
	// Create a context
	ctx := context.Background()

	// Load the configuration file
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		return "", fmt.Errorf("error reading config file: %v", err)
	}

	// Retrieve the GitHub access token from the configuration file
	githubAccessToken := viper.GetString("github_access_token")
	if githubAccessToken == "" {
		return "", fmt.Errorf("GitHub access token is missing or empty in the configuration file")
	}

	// Create a GitHub client with your personal access token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	owner := "avelino"
	repo := "awesome-go"

	// Fetch the README file from the repository
	readme, _, err := client.Repositories.GetReadme(ctx, owner, repo, nil)
	if err != nil {
		return "", fmt.Errorf("error fetching README: %v", err)
	}

	// Decode the README content
	readmeContent, err := readme.GetContent()
	if err != nil {
		return "", fmt.Errorf("error decoding README content: %v", err)
	}

	return readmeContent, nil
}

// GetProjects extracts GitHub repo names from README content.
func GetProjects(readmeContent string) []string {
	// Define a regular expression pattern to match GitHub repo names
	re := regexp.MustCompile(`github\.com/([\w\-]+)/([\w\-]+)`)

	// Find all matches in the README content
	matches := re.FindAllStringSubmatch(readmeContent, -1)

	// Extract the matched repo names
	var projects []string
	for _, match := range matches {
		if len(match) == 3 {
			// The first element is the full match, the second and third elements are the owner and repo names
			owner := match[1]
			repo := match[2]
			projectName := owner + "/" + repo
			projects = append(projects, projectName)
		}
	}

	return projects
}

func handleHelloWorld(serverResponse http.ResponseWriter, clientRequest *http.Request) {
	// If the path is not /projects, it responds with an HTTP 404 "Not Found" status by calling http.NotFound.
	if clientRequest.URL.Path != "/hello" {
		http.NotFound(serverResponse, clientRequest)
	}
	fmt.Fprintln(serverResponse, "Hello world!")
}

func handleReadme(serverResponse http.ResponseWriter, clientRequest *http.Request, readmeGetter ReadmeGetter) {
	if clientRequest.URL.Path != "/readme" {
		http.NotFound(serverResponse, clientRequest)
		return
	}

	// Fetch the README content from GitHub
	readmeContent, err := readmeGetter.GetREADME()
	if err != nil {
		http.Error(serverResponse, "Error fetching README", http.StatusInternalServerError)
		fmt.Println("Error fetching README:", err)
		return
	}

	// Set the response content type to plain text
	serverResponse.Header().Set("Content-Type", "text/plain")

	// Send the README content as the HTTP response
	fmt.Fprintln(serverResponse, readmeContent)
}

func handleProjects(serverResponse http.ResponseWriter, clientRequest *http.Request, readmeGetter ReadmeGetter) {
	if clientRequest.URL.Path != "/projects" {
		http.NotFound(serverResponse, clientRequest)
		return
	}

	// Fetch the README content from GitHub
	readmeContent, err := readmeGetter.GetREADME()
	if err != nil {
		http.Error(serverResponse, "Error fetching README", http.StatusInternalServerError)
		fmt.Println("Error fetching README:", err)
		return
	}

	// Extract project names from the README content
	projects := GetProjects(readmeContent)

	// Set the response content type to plain text
	serverResponse.Header().Set("Content-Type", "text/plain")

	// Send the list of project names as the HTTP response
	for _, projectName := range projects {
		fmt.Fprintln(serverResponse, projectName)
	}
}

func main() {
	// Set up the HTTP server
	server := &http.Server{
		Addr: ":8080",
	}

	// Create an instance of RealReadmeGetter
	realReadmeGetter := &RealReadmeGetter{}

	// Set up the route and inject the RealReadmeGetter instance
	http.HandleFunc("/readme", func(w http.ResponseWriter, r *http.Request) {
		handleReadme(w, r, realReadmeGetter)
	})
	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		handleProjects(w, r, realReadmeGetter)
	})
	http.HandleFunc("/hello", handleHelloWorld)

	fmt.Println("Server started on :8080")

	// Start the server and handle errors
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
