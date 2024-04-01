package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I love football"
	newStr := strings.ReplaceAll(str, "I", "We")
	fmt.Println(newStr)
}
