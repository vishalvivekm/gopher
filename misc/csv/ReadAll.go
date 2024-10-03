// ReadAll(): read all rows ( first header row) at once into a 2D slice of strings
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file ,err := os.Open("./friends.csv")
	if err != nil {
		log.Fatalln("can not open csv:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
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