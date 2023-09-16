package processing

import "regexp"

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