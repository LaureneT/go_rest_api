package main

import (
	"errors"
	"fmt"
	"testing"

	"net/http"
	"net/http/httptest"
)

func TestMyFunction(t *testing.T) {
	// Write your test logic here
	fmt.Println("Testing")
}

type MockReadmeGetter struct {
    MockedREADMEContent string
    MockedError         error
}

func (g *MockReadmeGetter) GetREADME() (string, error) {
    return g.MockedREADMEContent, g.MockedError
}

func TestHandleReadme(t *testing.T) {
    // Create a mock implementation of ReadmeGetter
    mockGetter := &MockReadmeGetter{
        MockedREADMEContent: "Mocked README Content.",
        MockedError:         nil,
    }

	// Create a mock HTTP request for testing
	req := httptest.NewRequest("GET", "/readme", nil)
	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Call the handleReadme function with the mock request and response
	handleReadme(w, req, mockGetter)

	// Check the HTTP status code (should be 200 OK for a successful response)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body 
	expectedResponseBody := "Mocked README Content.\n"
	actualResponseBody := w.Body.String()
	if expectedResponseBody != actualResponseBody {
		t.Errorf("Expected response body '%s', got '%s'", expectedResponseBody, actualResponseBody)
	}

	// Test case: Client request path is not "/projects"
    req = httptest.NewRequest("GET", "/otherpath", nil)
    w = httptest.NewRecorder()
    handleReadme(w, req, mockGetter)
    if w.Code != http.StatusNotFound {
        t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
    }

    // Test case: README Getter returns an error
    mockGetter.MockedError = errors.New("README fetch error")
    req = httptest.NewRequest("GET", "/readme", nil)
    w = httptest.NewRecorder()
    handleReadme(w, req, mockGetter)
    if w.Code != http.StatusInternalServerError {
        t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
    }
}

func TestGetProjects(t *testing.T){
	// Define some example README content
	readmeContent := `
	Here is a project: github.com/owner1/repo1
	Another project: github.com/owner2/repo2
	One more project: github.com/owner3/repo3
	`

	// Call the GetProjects function with the example README content
	projects := GetProjects(readmeContent)

	// Define the expected output in the requested format
	expectedOutput := []map[string]string{
		{"url": "https://github.com/owner1/repo1"},
		{"url": "https://github.com/owner2/repo2"},
		{"url": "https://github.com/owner3/repo3"},
	}

	// Check if the actual output format matches the expected output format
	if !compareProjects(projects, expectedOutput) {
		t.Errorf("GetProjects did not produce the expected output")
	}
}

// Helper function to compare project lists
func compareProjects(actual, expected []map[string]string) bool {
	if len(actual) != len(expected) {
		return false
	}
	for i, a := range actual {
		e := expected[i]
		if a["url"] != e["url"] {
			return false
		}
	}
	return true
}
