package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	words := strings.Split(text, " ")
	wordsMap := make(map[string]int)
	for _, ele := range words {
		value, found := wordsMap[ele]
		if found {
			wordsMap[ele] = value + 1
		} else {
		wordsMap[ele] = 1
		}
	}
	return wordsMap
}

func main() {
	text := "The quick brown fox jumps over the lazy dog the"
	fmt.Println(wordFrequency(text))
	//wordFrequency(text)

}



/*

package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string)  {
	words := strings.Split(text, " ")
	for _, ele := range words {
		fmt.Print(ele, " -> ")
	}
}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	//fmt.Println(wordFrequency(text))
	wordFrequency(text)

}

*/
