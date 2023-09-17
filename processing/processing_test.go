package processing_test

import (
	//"regexp"
	"encoding/json"
	"reflect"
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

func TestFilterProjectsByName(t *testing.T) {
	// Sample projects for testing
	projects := []processing.Project{
		{Name: "Project1", URL: "URL1"},
		{Name: "Project2", URL: "URL2"},
		{Name: "AnotherProject", URL: "URL3"},
	}

	t.Run("Filter projects by name (case-sensitive)", func(t *testing.T) {
		// Test filtering with a name that exists in project names
		filtered := processing.FilterProjectsByName(projects, "project1")
		expected := []processing.Project{{Name: "Project1", URL: "URL1"}}
		if !reflect.DeepEqual(expected, filtered) {
			t.Errorf("Expected %v, got %v.", expected, filtered)
		}

		// Test filtering with a name that doesn't exist in project names
		filtered = processing.FilterProjectsByName(projects, "nonexistent")
		if len(filtered) != 0 {
			t.Errorf("Expected 0 projects, got %d", len(filtered))
		}
	})

	t.Run("Filter projects by name (empty name)", func(t *testing.T) {
		// Test filtering with an empty name
		filtered := processing.FilterProjectsByName(projects, "")
		if len(filtered) != len(projects) {
			t.Errorf("Expected %d projects, got %d", len(projects), len(filtered))
		}
		for i, p := range projects {
			if p != filtered[i] {
				t.Errorf("Expected '%s', got '%s'", p.Name, filtered[i].Name)
			}
		}
	})
}
