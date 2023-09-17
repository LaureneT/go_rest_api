package network

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func FetchReadmeFromGitHub(repoOwner string, repoName string) (string, error) {
	ctx := context.Background()

	// Load the configuration file
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		return "", fmt.Errorf("error reading config file: %v", err)
	}

	// Retrieve the GitHub access token 
	githubAccessToken := viper.GetString("github_access_token")
	if githubAccessToken == "" {
		return "", fmt.Errorf("GitHub access token is missing or empty in the configuration file")
	}

	// Create a GitHub client with your personal access token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Fetch the README file from the repository
	readme, _, err := client.Repositories.GetReadme(ctx, repoOwner, repoName, nil)
	if err != nil {
		return "", fmt.Errorf("error fetching README: %v", err)
	}

	// Decode the README content
	readmeContent, err := readme.GetContent()
	if err != nil {
		return "", fmt.Errorf("error decoding README content: %v", err)
	}

	return readmeContent, nil
}
