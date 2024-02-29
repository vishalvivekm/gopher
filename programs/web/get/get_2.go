package main

// receive a json response
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data       []struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	} `json:"data"`
	Support struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}

func main() {
	response, err := http.Get("https://reqres.in/api/users?page=2")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body) // response body is []byte
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(reflect.TypeOf(body)) // []uint8

	//	fmt.Println(string(body)) // convert to string before print
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		fmt.Println("request is successful")
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("can not unmarshal JSON")
	}
	for _, person := range result.Data {
		fmt.Println(person.FirstName)
	}
}

// convert JSON to go struct: https://mholt.github.io/json-to-go/
