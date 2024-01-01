package main

import (
	"fmt"
	"strings"
)

func main() {
	//var roll int
	//if val, err := fmt.Scan(&roll); err != nil {
	//	fmt.Println("reported error: ", err)
	//	return
	//} else {
	//	fmt.Println(val, err)
	//}
	algorithms := parseAlgorithms()
	fmt.Println(strings.Join(algorithms, ";")) // TODO: format algorithms
}

// DO NOT CHANGE
func parseAlgorithms() []string {
	var algorithms []string // this is slice consisting of strings
	for {                   // a infinite for loop ( CTRL + D)
		var algo string
		if _, err := fmt.Scan(&algo); err != nil {
			break
		}
		algorithms = append(algorithms, algo)
	}
	return algorithms
}
