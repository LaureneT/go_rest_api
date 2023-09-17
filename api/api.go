package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/LaureneT/go_rest_api/network"
	"github.com/LaureneT/go_rest_api/processing"
	"github.com/spf13/viper"
)

func HandleProjects(serverResponse http.ResponseWriter, clientRequest *http.Request) {
	queryParams := clientRequest.URL.Query()
	projectName := queryParams.Get("name")
	// As requested, presence of curly braces do not prevent the searching feature.
	// Should be evaluated with client and remove if necessary as it can be a source of errors.
	projectName = strings.Trim(projectName, "{}")

	// Load the configuration file
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error reading config file: %v", err)
		return
	}

	owner := viper.GetString("github_repo_owner")
	repo := viper.GetString("github_repo_name")
	readmeContent, err := network.FetchReadmeFromGitHub(owner, repo)
	if err != nil {
		http.Error(serverResponse, "Error fetching README", http.StatusInternalServerError)
		fmt.Println("Error fetching README:", err)
		return
	}

	projects := processing.ExtractProjectsFromReadme(readmeContent)

	if projectName != "" { // "" filter is always true
		projects = processing.FilterProjectsByName(projects, projectName)
	}

	responseJSON, err := processing.JSONifyProjects(projects)
	if err != nil {
		http.Error(serverResponse, "Failed to format data to JSON", http.StatusInternalServerError)
		return
	}

	// Send the JSON response
	serverResponse.Header().Set("Content-Type", "application/json")
	serverResponse.WriteHeader(http.StatusOK)
	serverResponse.Write([]byte(responseJSON))
}
