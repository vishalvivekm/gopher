package main

// receive a json response
import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(reflect.TypeOf(body)) // []uint8

	fmt.Println(string(body))
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		fmt.Println("request is successful")
	}
}
