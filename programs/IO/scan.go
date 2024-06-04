package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func trimmer(word string) string {
	return strings.TrimSpace(word)
}

func main() {
	dictionary := make([]string, 0)
	var word string
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word = scanner.Text()
		word = trimmer(word)
		dictionary = append(dictionary, word)
	}

	for _, word = range dictionary {
		fmt.Println(word)
	}
}
