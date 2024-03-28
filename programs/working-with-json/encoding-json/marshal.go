package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func main() {
	// enocode simple maps to json
	student := map[string]interface{}{"name": "Vivek Vishal", "CGPA": 9.40}

	// json.Marshal() encodes/marshals the 'student' map:
	studentJson, err := json.Marshal(student) // returns []bytes which contains serialized data and an error, if any.
	studentJsonWithIndentn, err := json.MarshalIndent(student, "", " ")
	checkError(err)
	fmt.Println(studentJson)                    // [123 10 32 32 34 67 71 80 65 34 58 32 57 46 52 44 10 32 32 34 110 97 109 101 34 58 32 34 86 105 118 101 107 32 86 105 115 104 97 108 34 10 125]
	fmt.Println(string(studentJson))            // {"CGPA":9.4,"name":"Vivek Vishal"}
	fmt.Println(string(studentJsonWithIndentn)) // output:
	/*
		{
		 "CGPA": 9.4,
		 "name": "Vivek Vishal"
		}
	*/

	// serialize a simple slice to json
	friends := []string{"Sumit", "Diru", "Shash", "Vivek"}

	friendsJson, err := json.Marshal(friends)
	checkError(err)
	fmt.Println(string(friendsJson))
	// Output:
	// ["Sumit", "Diru", "Shash", "Vivek"]

	// encoding/serializing a nested json objects

	friendss := map[string]interface{}{
		"friends": []interface{}{
			map[string]interface{}{
				"rollNo":  2140105,
				"name":    "Shashwat",
				"surname": "Verma",
			},
			map[string]interface{}{
				"rollNo":  2140148,
				"surname": "Singh",
				"name":    "Dhirendra",
			},
			[]string{"Sumit", "Lakshya", "Rupesh"},
		},
	}
	// fmt.Println(friendss) // map[friends:[map[name:Shashwat rollNo:2140105 surname:Verma] map[name:Dhirendra rollNo:2140148 surname:Singh] [Sumit Lakshya Rupesh]]]
	friendssJson, err := json.MarshalIndent(friendss, "", " ")
	checkError(err)
	fmt.Println(string(friendssJson))
	/*
	   	{
	   		"friends": [
	   	{
	   	"name": "Shashwat", // map entries were sorted alphabetically while serializing
	   	"rollNo": 2140105,
	   	"surname": "Verma"
	   	},
	   	{
	   	"name": "Dhirendra",
	   	"rollNo": 2140148,
	   	"surname": "Singh"
	   	},
	   	[
	   	"Sumit",
	   	"Lakshya",
	   	"Rupesh"
	   	]
	   ]
	   }*/
	// Tip: when we encode/serialize a map to JSON, the entries are sorted based on map keys, that's why name is the first key of serialized output above
}

// map || slice -> []byte -> JSON Object
