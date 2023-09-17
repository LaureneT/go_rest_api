package main

import (
	"fmt"
	"net/http"

	"github.com/LaureneT/go_rest_api/api"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/projects", func(serverResponse http.ResponseWriter, clientRequest *http.Request) {
		api.HandleProjects(serverResponse, clientRequest)
	})

	fmt.Println("Server started on", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
