package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	//for _, ele := range text {
	//	fmt.Println(string(ele))
	//}
	//  strings.Split(str, "sep") []string
	for _, ele := range strings.Split(text, " ") {
		fmt.Println(ele)
	}

}
