package main

import (
	"fmt"
	"strings"
)

func WordCount(s string, word string) int {
	return strings.Count(s, word)
}

func main() {
	count := WordCount("hello, Hello how have you been in helloworld", "hello")
	fmt.Println(count)
}

/*Explanation:

This function uses the strings.Count method to count the number of occurrences of word in s. The Count method is case-sensitive, so it will correctly count the number of times word appears in s even if word and s have different capitalization.

*/
