# My RESTful Go API

Welcome to the My RESTful Go API! This API is built in Go and provides functionality for managing GitHub projects listed in the README of GitHub repositories.

## Table of Contents
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Configuration](#configuration)
  - [Adding Your GitHub Access Token](#adding-your-github-access-token)
- [Usage](#usage)
  - [Endpoints](#endpoints)
- [Tests](#tests)
- [Build](#build)
- [Run](#run)

## Getting Started

### Prerequisites
- Go 1.15+

### Installation
1. Clone the repository: `git clone https://github.com/LaureneT/go_rest_api.git`
2. Navigate to the project directory: `cd go_rest_api`
3. Install dependencies: `go mod tidy`
4. Configure the environment
5. Run the API: `go run main.go`

## Configuration

This project uses a `config.json` file to manage configuration settings, including the GitHub access token. To get started, follow these steps:

1. Locate the `config.json.example` file in the project's root directory.
2. Make a copy of `config.json.example` and rename it to `config.json`.
3. Open `config.json` using a text editor of your choice.

### Adding Your GitHub Access Token

In `config.json`, you will find a placeholder for the GitHub access token:

```json
{
  "github_access_token": "YOUR_GITHUB_ACCESS_TOKEN"
}
```

## Usage

### Endpoints
- **GET /projects**: Retrieve a list of all GitHub projects listed in the readme of a GitHub repo.
- **GET /projects?name={project_name}**: Retrieve a list of GitHub projects with name that contains project_name.


# Tests

You can run tests using the following command:
go test

# Build
To build the project, execute:
go build

# run
To run the API, use the following command:

```sh
./go_rest_api.exe
```

You can access the API at http://localhost:8080/projects.