package main

import (
	"fmt"
	"unicode/utf8"
)

func changeString(str string) string {
	str = "hi there"
	return str
}

func main() {
	// a string - a sequence of variable-width chars, each represented by >= 1 bytes using UTF-8 encoding.
	// double quotes
	var firstString string
	firstString = "ABCD\n\tEF"
	fmt.Println(firstString)
	secondString := `ABCD\n\t` // retains special chars, value: ABCD\n\t
	fmt.Println(secondString)

	// In golang, string contents are immutable, on concatenating 2, it creates a new one in memory.

	greeting := "Hi"               // the variable contains the "Hi" string
	greeting = greeting + ", Sam!" // the variable contains a new "Hi, Sam!" string; the "Hi" string will soon be removed from the memeory

	fmt.Println(greeting)
	str := "vivek"
	fmt.Println(changeString(str)) // hi there
	fmt.Println(str)               // vivek meaning strings are passed as value in functions in golang

	// The actual strings like "Hi" and "vivek" are immutable, we can change the value of a string variable. It still means that the string "Hi" exists in
	// the memory somewhere and Golang will not change the contents of that memory location
	//
	// to find out a string's bytelength: len(stringName)
	asciiString := "ABCDE"
	utf8String := "БГДЖИ"

	fmt.Println(len(asciiString)) // 5
	fmt.Println(len(utf8String))  // 10 - byte length

	// Rune type is used in go to represent unicode char. ( rune is alias for int32, so it can be evident when an int value is representing a codepoint)
	// strings are not only sequence of bytes but also sequence of runes
	// as sequence of bytes: when transferring data,
	// as sequence of runes: when checking each individual char of the string

	// to find length of string in characters: func RuneCountInString from unicode/utf8 package:
	fmt.Println(utf8.RuneCountInString(asciiString)) // 5
	fmt.Println(utf8.RuneCountInString(utf8String))  // 5

	emoji := "🙋🌍❗"
	fmt.Println(len(emoji), " ", utf8.RuneCountInString(emoji)) // 11 3

	fmt.Println(len([]rune("hi there")))
}

// string literals in Go: "" or `` ( raw string literal)
// zero values: empty string
// convert strings to []rune for security

// rune type:
/*
as an alias for int32, so programs can be clear when an integer value represents a code poitn
to process individual chars in a string when character-by-character analysis is req
representing unicode code point mapping
*/
