package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/LaureneT/go_rest_api/network"
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
	projects := GetProjects(readmeContent)

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

// GetProjects extracts GitHub repo URLs from README content.
func GetProjects(readmeContent string) []map[string]string {
	// Define a regular expression pattern to match GitHub repo URLs
	re := regexp.MustCompile(`github\.com/([\w\-]+)/([\w\-]+)`)

	// Find all matches in the README content
	matches := re.FindAllStringSubmatch(readmeContent, -1)

	// Extract the matched repo URLs
	var projects []map[string]string
	for _, match := range matches {
		if len(match) == 3 {
			// The first element is the full match, the second and third elements are the owner and repo names
			owner := match[1]
			repo := match[2]
			// Construct the GitHub project URL
			projectURL := "https://github.com/" + owner + "/" + repo
			projectMap := map[string]string{"url": projectURL}
			projects = append(projects, projectMap)
		}
	}

	return projects
}
