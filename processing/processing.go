package processing

import (
	"encoding/json"
	"regexp"
	"strings"
)

type Project struct {
	Name string
	URL  string
}

func ExtractProjectsFromReadme(readmeContent string) ([]Project) {
	// Define a regular expression pattern to match GitHub repo URLs
	re := regexp.MustCompile(`https?://github\.com/([\w\-]+)/([\w\-]+)/?`)

	matches := re.FindAllStringSubmatch(readmeContent, -1)

	var projects []Project
	for _, match := range matches {
		if len(match) == 3 {
			owner := match[1]
			repo := match[2]
			// Construct the GitHub project URL
			projectURL := "https://github.com/" + owner + "/" + repo
			project := Project{Name: repo, URL: projectURL}
			projects = append(projects, project)
		}
	}
	return projects
}

func JSONifyProjects(projects []Project) (string, error) {
	if len(projects) == 0 {
		// If projects is empty, create an empty JSON structure
		emptyJSON := map[string][]map[string]string{
			"projects": {},
		}

		responseJSON, err := json.Marshal(emptyJSON)
		if err != nil {
			return "", err
		}
		return string(responseJSON), nil
	}

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

func FilterProjectsByName(projects []Project, name string) []Project {
	var filteredProjects []Project
	for _, project := range projects {
		if strings.Contains(strings.ToLower(project.Name), strings.ToLower(name)) {
			filteredProjects = append(filteredProjects, project)
		}
	}
	return filteredProjects
}
