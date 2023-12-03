package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func finalWords(obiwanWords, anakinWords string) string {

	return strings.Join([]string{strings.ToLower(obiwanWords), strings.ToUpper(anakinWords)}, " ")
	// return strings.Join([]string{strings.ToLower(obiwanWords), strings.ToUpper(anakinWords)}, " ")
	// return strings.ToLower(obiwanWords)+ " " + strings.ToUpper(anakinWords)

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	obiwanWords := scanner.Text()
	scanner.Scan()
	anakinWords := scanner.Text()

	fmt.Println(finalWords(obiwanWords, anakinWords))
}
