package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	file ,err := os.Create("new_friends.csv")
	if err != nil {
		log.Fatalln("can not create new_friends csv:", err)
	}
	defer file.Close()

	data := [][]string{
		{"1", "John", "Doe", "john@google.in"},
		{"2", "Jane", "Doe", "jane@google.in"},
		{"3", "Michael", "Smith", "michael@google.in"},
	}

	writer := csv.NewWriter(file)
	//err = writer.WriteAll(data)
	// if err != nil {
	// 	log.Fatalln("can not write to csv file: ", err)
	// }
	for _, row := range data {
		err := writer.Write(row)
		if err != nil {
	    	log.Fatalln("can not write to csv file: ", err)
		}
	}
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln("can not read file: ",err)
	}
	log.Println(string(content))
	writer.Flush()

	time.Sleep(1 * time.Second)
	conten, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln("can not read file: ",err)
	}
	log.Println(string(conten))

	log.Println("INFO: data successfully written to csv ",file.Name())
}