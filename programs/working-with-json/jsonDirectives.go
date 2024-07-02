package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

//type Student struct { // this column has the struct tags
//	ID          int       `json:"id"`
//	IsTopper    bool      `json:"isTopper"`
//	DateOfBirth time.Time `json:"dateOfBirth"`
//	Email       string    `json:"email"`
//	key         int       `json:"key"`
//}

type Student struct { // this column has the struct tags
	ID          int       `json:"id,string"`
	IsTopper    bool      `json:"isTopper,omitempty"`
	DateOfBirth time.Time `json:"dateOfBirth,omitempty"`
	Email       string    `json:"email"`
	key         int       `json:"key"`
	NickName    string    `json:"-"`
}

func main() {
	//student := Student{ID: 2140, IsTopper: true, DateOfBirth: time.Date(2009, time.April,
	//	24, 5, 72, 01, 0, time.UTC), Email: "sus@among.us", key: 23}

	student := Student{ID: 2103, Email: "sus@among.us", NickName: "vivi"}
	studentJSON, err := json.Marshal(student)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(studentJSON))
}

// Output:
// {"id":2140,"isTopper":true,"dateOfBirth":"2009-04-24T06:12:01Z","email":"sus@among.us"}
// no key as it's private field in studnet struct

// output: {"id":2103,"dateOfBirth":"0001-01-01T00:00:00Z","email":"sus@among.us"}
// why does DateOfBirth, which is not initialised, appear, even when we passed omitempty directive to it ?
// concept: omitempty doesn't consider struct type to be empty and time.Time is actually a struct

// reference: https://www.alexedwards.net/blog/json-surprises-and-gotchas#7
// json to go struct converter: https://mholt.github.io/json-to-go/

//json directives:
/*
`json:"omitempty"` - omits a struct field when it contains a zero/default value of its type
`json:"-"` - this directive makes that the field is never encoded into JSON
`json:"string"` - forces the data in an individual field to be encoded as a string in the resulting JSON


Private fields ( starting with a lowercase letter ) can never be exported or made accessible to outside packages like json...
doesn't exist: this directive allows private struct fields to be encoded into JSON
*/
