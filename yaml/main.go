package main

import (
	_ "bytes"
	"fmt"
	"log"
	"os"
	"time"
	"vishalvivekm/yaml/utils"
)

type Person struct {
    Name         string
    Friends        []string
    Age          int
    privateRandomField int32 // private fields not accessible to foreign packages i.e. the yaml.Marshal() function here, of gopkg.in/yaml.v3 package and is used for encoding in the main package.
    CreditScore      float32
    CurrentDate time.Time
}

func main() {
	myself := Person{
		Name: "Vivek Vishal",
		Friends: []string{"Shashwat, Sumit, Ritik, Leetcoder"},
		Age: 21,
		privateRandomField: 600,
		CreditScore: 850.34,
		CurrentDate: time.Now(),
	}
	// Serialize the `myself` struct into a YAML object
    // yaml.Marshal() (bytes, error) {}:
   
	myYAML, err := utils.SerializeToYaml(myself) //  //myYAML, err := yaml.Marshal(myself)
    if err != nil {
        log.Println(err)
		os.Exit(1)
    }
    //fmt.Println(string(myYAML))
	file, err := os.OpenFile("info.yaml", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0644)
	if err != nil {
		log.Println("error opening / creating file: ", err)
	}
	defer file.Close()

	if _, err = fmt.Fprintln(file, string(myYAML)); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("data written to info.yaml...")
}


