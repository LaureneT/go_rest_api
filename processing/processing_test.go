package processing_test

import (
	//"regexp"
	"encoding/json"
	"testing"

	"github.com/LaureneT/go_rest_api/processing"
	"github.com/stretchr/testify/assert"
)

func TestExtractProjectsFromReadme_ValidInput(t *testing.T) {
	// Sample README content with GitHub repo URLs
	readmeContent := `
		Check out these GitHub repos:
		https://github.com/user1/repo1
		https://github.com/user2/repo2
		https://github.com/user3/repo3
	`
	expectedProjects := []processing.Project{
		{Name: "repo1", URL: "https://github.com/user1/repo1"},
		{Name: "repo2", URL: "https://github.com/user2/repo2"},
		{Name: "repo3", URL: "https://github.com/user3/repo3"},
	}

	// Call the function
	projects, err := processing.ExtractProjectsFromReadme(readmeContent)

	// Assertions
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, expectedProjects, projects, "Expected extracted projects to match")
}

func TestExtractProjectsFromReadme_NoMatches(t *testing.T) {
	// Sample README content with no GitHub repo URLs
	readmeContent := "This is a readme with no GitHub repo URLs."

	// Call the function
	projects, err := processing.ExtractProjectsFromReadme(readmeContent)

	// Assertions
	assert.NoError(t, err, "Expected no error")
	assert.Empty(t, projects, "Expected no extracted projects")
}

func TestFormatToJSON_ValidInput(t *testing.T) {
	// Sample projects to format to JSON
	projects := []processing.Project{
		{Name: "repo1", URL: "https://github.com/user1/repo1"},
		{Name: "repo2", URL: "https://github.com/user2/repo2"},
		{Name: "repo3", URL: "https://github.com/user3/repo3"},
	}

	// Call the function
	jsonStr, err := processing.FormatToJSON(projects)

	// Assertions
	assert.NoError(t, err, "Expected no error")

	// Unmarshal the JSON string to verify its structure
	var parsedJSON map[string][]map[string]string
	err = json.Unmarshal([]byte(jsonStr), &parsedJSON)
	assert.NoError(t, err, "Expected valid JSON")

	// Check the structure of the parsed JSON
	expectedJSON := map[string][]map[string]string{
		"projects": {
			{"url": "https://github.com/user1/repo1"},
			{"url": "https://github.com/user2/repo2"},
			{"url": "https://github.com/user3/repo3"},
		},
	}
	assert.Equal(t, expectedJSON, parsedJSON, "Expected JSON structure to match")
}

func TestFormatToJSON_EmptyInput(t *testing.T) {
	// Empty input, no projects to format
	projects := []processing.Project{}

	// Call the function
	jsonStr, err := processing.FormatToJSON(projects)

	// Assertions
	assert.NoError(t, err, "Expected no error")

	// Unmarshal the JSON string to verify its structure
	var parsedJSON map[string][]map[string]string
	err = json.Unmarshal([]byte(jsonStr), &parsedJSON)
	assert.NoError(t, err, "Expected valid JSON")

	// Check the structure of the parsed JSON
	expectedJSON := map[string][]map[string]string{
		"projects": {},
	}
	assert.Equal(t, expectedJSON, parsedJSON, "Expected empty JSON structure")
}
