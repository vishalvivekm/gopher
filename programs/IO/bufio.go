// return number of words from a input line
//assumption:  the words in line are seperated by a single comma

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//func numberOfWords(line string) int {
//	return len(line) // number of words
//} // returns byte-length

func numberOfWords(line string) int {
	return len(strings.Fields(line)) // number of words, no assumption

	/*slice := strings.Split(strings.TrimSpace(line), " ") //
	  return len(slice)
	*/

	//or, return len(strings.Split(line, " "))

	//return len(strings.Split(line, " "))
}

//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Scan()
//	fmt.Println(numberOfWords(scanner.Text()))
//	fmt.Println(reflect.TypeOf(scanner)) // *bufio.Scanner
//}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(numberOfWords(scanner.Text()))
	fmt.Print("Hi, everyone")
	var roll int
	if val, err := fmt.Scan(&roll); err != nil {
		fmt.Println("reported error: ", err)
		return
	} else {
		fmt.Println(val, err)
	}
}
