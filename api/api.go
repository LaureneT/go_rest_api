package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/LaureneT/go_rest_api/network"
	"github.com/LaureneT/go_rest_api/processing"
)

func HandleHelloWorld(serverResponse http.ResponseWriter, clientRequest *http.Request) {
	if clientRequest.URL.Path != "/hello" {
		http.NotFound(serverResponse, clientRequest)
	}
	fmt.Fprintln(serverResponse, "Hello world!")
}

func HandleReadme(serverResponse http.ResponseWriter, clientRequest *http.Request) {
	if clientRequest.URL.Path != "/readme" { // useful ?
		http.NotFound(serverResponse, clientRequest)
		return
	}

	// Fetch the README content from GitHub
	readmeContent, err := network.FetchReadmeFromGitHub()
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

func HandleProjects(serverResponse http.ResponseWriter, clientRequest *http.Request) {
	if clientRequest.URL.Path != "/projects" { // Necessary ?
		http.NotFound(serverResponse, clientRequest)
		return
	}

	queryParams := clientRequest.URL.Query()
	projectName := queryParams.Get("name")
	// Remove curly braces from projectName if present
	projectName = strings.Trim(projectName, "{}")

	// Fetch the README content from GitHub
	readmeContent, err := network.FetchReadmeFromGitHub()
	if err != nil {
		http.Error(serverResponse, "Error fetching README", http.StatusInternalServerError)
		fmt.Println("Error fetching README:", err)
		return
	}

	// Extract projects from the README content
	projects, err := processing.ExtractProjectsFromReadme(readmeContent)
	if err != nil {
		http.Error(serverResponse, "Failed to extract project names", http.StatusInternalServerError)
		return
	}

	// filter projects by name if a project_name is input
	if projectName != "" {
		projects = processing.FilterProjectsByName(projects, projectName)
	}

	// format the data to JSON
	responseJSON, err := processing.FormatToJSON(projects)
	if err != nil {
		http.Error(serverResponse, "Failed to format data to JSON", http.StatusInternalServerError)
		return
	}

	// Send the JSON response
	serverResponse.Header().Set("Content-Type", "application/json")
	serverResponse.WriteHeader(http.StatusOK)
	serverResponse.Write([]byte(responseJSON))
}