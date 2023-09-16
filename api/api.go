package api

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	// Fetch the README content from GitHub
	readmeContent, err := network.FetchReadmeFromGitHub()
	if err != nil {
		http.Error(serverResponse, "Error fetching README", http.StatusInternalServerError)
		fmt.Println("Error fetching README:", err)
		return
	}

	// Extract project URLs from the README content
	projects := processing.GetProjects(readmeContent)

	// Create a map with a "projects" key
	responseMap := map[string][]map[string]string{
		"projects": projects,
	}

	// Marshal the response map into JSON format
	responseJSON, err := json.Marshal(responseMap)
	if err != nil {
		http.Error(serverResponse, "Error encoding JSON", http.StatusInternalServerError)
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Set the response content type to JSON
	serverResponse.Header().Set("Content-Type", "application/json")

	// Send the JSON response as the HTTP response
	serverResponse.Write(responseJSON)
}
