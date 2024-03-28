package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
reverse of encoding/Marshaling/serializing - decoding/unmarshaling/deserializing

func Unmarshal(data []byte, v any) error
Json.Unmarshal - expects the Json object as slice of bytes i.e []byte and a pointer (&.) to the variable where we'll be storing the
decoded json object.
*/

func main() {
	studentJson := `{"CGPA":9.4,"name":"Vivek Vishal"}`
	// here's the  JSON object is collection of key:value pairs, a map
	// we'll need a map to store the decoded/unmarshaled JSON data

	var student map[string]interface{}

	if err := json.Unmarshal([]byte(studentJson), &student); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(student) // map[CGPA:9.4 name:Vivek Vishal]

	friendsJson := ` ["Sumit", "Diru", "Shash", "Vivek", 21]`
	var s []interface{}
	err := json.Unmarshal([]byte(friendsJson), &s)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(s) // [Sumit Diru Shash Vivek 21]
}

// JSON object -> []byte ->
