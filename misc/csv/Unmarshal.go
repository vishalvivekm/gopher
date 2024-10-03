// unmarshalling friends csv to friends struct
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Friends struct {
	Id int
	FirstName, LastName, Email string
}

func main(){
	file ,err := os.Open("./friends.csv")
	if err != nil {
		log.Fatalln("can not open csv:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Read() // skip header 
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("can not read CSV file: ", err)
	}

	friends := []Friends{}
	for _, row := range rows {
		id, _ := strconv.ParseInt(row[0], 0, 0)
		friend := Friends{
			Id: int(id),
			FirstName: row[1],
			LastName: row[2],
			Email: row[3],
		}
		friends = append(friends, friend)
	}
	for _, friend := range friends {
	fmt.Printf("%+v\n", friend)
	}

}