package processing_test

import (
	"encoding/json"
	"testing"

	"github.com/LaureneT/go_rest_api/processing"
	"github.com/stretchr/testify/assert"
)

func TestExtractProjectsFromReadme_ValidInput(t *testing.T) {
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

	projects := processing.ExtractProjectsFromReadme(readmeContent)

	assert.Equal(t, expectedProjects, projects, "Expected extracted projects to match")
}

func TestExtractProjectsFromReadme_NoMatches(t *testing.T) {
	readmeContent := "This is a readme with no GitHub repo URLs."

	projects := processing.ExtractProjectsFromReadme(readmeContent)

	assert.Empty(t, projects, "Expected no extracted projects")
}

func TestFormatToJSON_ValidInput(t *testing.T) {
	projects := []processing.Project{
		{Name: "repo1", URL: "https://github.com/user1/repo1"},
		{Name: "repo2", URL: "https://github.com/user2/repo2"},
		{Name: "repo3", URL: "https://github.com/user3/repo3"},
	}

	jsonStr, err := processing.JSONifyProjects(projects)

	assert.NoError(t, err, "Expected no error")

	// Unmarshal the JSON string to verify its structure
	var parsedJSON map[string][]map[string]string
	err = json.Unmarshal([]byte(jsonStr), &parsedJSON)
	assert.NoError(t, err, "Expected valid JSON")

	// Check the structure of the parsed JSON
	expectedObject := map[string][]map[string]string{
		"projects": {
			{"url": "https://github.com/user1/repo1"},
			{"url": "https://github.com/user2/repo2"},
			{"url": "https://github.com/user3/repo3"},
		},
	}
	assert.Equal(t, expectedObject, parsedJSON, "Expected JSON structure to match")
}

func TestFormatToJSON_EmptyInput(t *testing.T) {
	projects := []processing.Project{}

	jsonStr, err := processing.JSONifyProjects(projects)

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

func TestFilterProjectsByName_CaseSensitive(t *testing.T) {
	projects := []processing.Project{
		{Name: "Project1", URL: "URL1"},
		{Name: "Project2", URL: "URL2"},
		{Name: "AnotherProject", URL: "URL3"},
	}

	// Test filtering with a name that exists in project names
	filtered := processing.FilterProjectsByName(projects, "project1")
	expected := []processing.Project{{Name: "Project1", URL: "URL1"}}
	assert.Equal(t, expected, filtered, "Filtered projects do not match expected")

	// Test filtering with a name that doesn't exist in project names
	filtered = processing.FilterProjectsByName(projects, "nonexistent")
	assert.Empty(t, filtered, "Filtered projects should be empty")
}

func TestFilterProjectsByName_EmptyName(t *testing.T) {
	projects := []processing.Project{
		{Name: "Project1", URL: "URL1"},
		{Name: "Project2", URL: "URL2"},
		{Name: "AnotherProject", URL: "URL3"},
	}

	// Test filtering with an empty name
	filtered := processing.FilterProjectsByName(projects, "")
	assert.Len(t, filtered, len(projects), "Filtered projects length does not match expected")
	assert.ElementsMatch(t, projects, filtered, "Filtered projects do not match expected")
}
