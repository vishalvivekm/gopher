package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file ,err := os.Open("./friends_1.csv") // ; as delimiter and # as comment
	if err != nil {
		log.Fatalln("can not open csv:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.Comment = '#'
	
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("can not read CSV file: ", err)
	}
	for _, row := range rows {
		for _, val := range row {
			fmt.Printf("%v ",val)
		}
		fmt.Println()
	}
}