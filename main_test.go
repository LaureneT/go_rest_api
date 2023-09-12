package main

import (
	"errors"
	"fmt"
	"reflect"
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
	req := httptest.NewRequest("GET", "/projects", nil)
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
    req = httptest.NewRequest("GET", "/projects", nil)
    w = httptest.NewRecorder()
    handleReadme(w, req, mockGetter)
    if w.Code != http.StatusInternalServerError {
        t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
    }
}


func TestGetProjects(t *testing.T){
	// Mocked readmeContent
	mockReadmeContent := "github.com/user1/repo1 github.com/user2/repo2 github.com/user3/repo3"

	// Call the GetProjects function with the mocked readmeContent
	result := GetProjects(mockReadmeContent)

	// Check if the result is a string
	if reflect.TypeOf(result).Kind() != reflect.Slice || reflect.TypeOf(result).Elem().Kind() != reflect.String {
		t.Errorf("Expected a []string, got %v", reflect.TypeOf(result))
	}
}