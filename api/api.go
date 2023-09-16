package api

import (
	"fmt"
	"net/http"
)

func HandleHelloWorld(serverResponse http.ResponseWriter, clientRequest *http.Request) {
	if clientRequest.URL.Path != "/hello" {
		http.NotFound(serverResponse, clientRequest)
	}
	fmt.Fprintln(serverResponse, "Hello world!")
}
