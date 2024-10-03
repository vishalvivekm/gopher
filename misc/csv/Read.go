// Read(): Read one row at a time into a slice of strings

package main

import (
	"encoding/csv"
	"fmt"
	"io"
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

	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("can not read CSV file: ",err)
		}		
		for _, val := range row {
			fmt.Printf("%v ",val)
		}
		fmt.Println()

	}

}