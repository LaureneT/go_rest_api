package processing

import (
	"encoding/json"
	"regexp"
)

// ExtractProjectsFromReadme extracts GitHub repo URLs from README content.
func ExtractProjectsFromReadme(readmeContent string) ([]Project, error) {
	// Define a regular expression pattern to match GitHub repo URLs
	re := regexp.MustCompile(`github\.com/([\w\-]+)/([\w\-]+)`)

	// Find all matches in the README content
	matches := re.FindAllStringSubmatch(readmeContent, -1)

	// Extract the matched repo URLs
	var projects []Project
	for _, match := range matches {
		if len(match) == 3 {
			// The first element is the full match, the second and third elements are the owner and repo names
			owner := match[1]
			repo := match[2]
			// Construct the GitHub project URL
			projectURL := "https://github.com/" + owner + "/" + repo
			project := Project{Name: repo, URL: projectURL}
			projects = append(projects, project)
		}
	}

	return projects, nil
}

type Project struct {
	Name string
	URL  string
}

func FormatToJSON(projects []Project) (string, error) {
	// Create a slice to hold the URLs
	var projectURLs []map[string]string

	// Extract URLs from the Project structures
	for _, project := range projects {
		projectURL := map[string]string{
			"url": project.URL,
		}
		projectURLs = append(projectURLs, projectURL)
	}

	// Create a map with a "projects" key and the project URLs
	responseMap := map[string][]map[string]string{
		"projects": projectURLs,
	}

	// Marshal the response map into JSON format
	responseJSON, err := json.Marshal(responseMap)
	if err != nil {
		return "", err
	}

	return string(responseJSON), nil
}
