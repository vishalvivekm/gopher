package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func main() {
	// Replace this URL with the webpage you want to check
	url := "https://example.com"

	// Send an HTTP GET request to the URL
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer response.Body.Close()
	fmt.Println(reflect.TypeOf(1))
	// Check the HTTP response status code
	if response.StatusCode == http.StatusOK {
		fmt.Printf("Webpage exists (HTTP Status 200 OK)\n")
	} else if response.StatusCode == http.StatusNotFound {
		fmt.Printf("Webpage does not exist (HTTP Status 404 Not Found)\n")
	} else {
		fmt.Printf("Webpage status code: %d\n", response.StatusCode)
	}
}
