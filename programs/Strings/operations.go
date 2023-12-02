package main

import (
	"fmt"
	"strings"
)

func main() {
	// func ToLower(str string) string
	firstName := "VIVEK"
	secondName := "vishal"
	fmt.Println(strings.ToLower(firstName))
	fmt.Println(strings.ToUpper(secondName))

	// func Trim(str, cutset string) string - to trim from both side
	password := "abc48password8ba4c"
	fmt.Println(strings.Trim(password, "abc48"))

	// func TrimLeft(s, cutset string) string
	//   func TrimRight(s, cutset string) string

	// strings.TrimSpace: removes space characters, such as \n\r\t from both ends of the strings:
	//
	outerSpace := "    \tOuter Space...  \r\n"
	fmt.Println(strings.TrimSpace(outerSpace)) // Outer Space...

	// separating and concatenating
	cellNumber := "8790-432-43-98"
	fmt.Println(strings.Split(cellNumber, "-"))
	fullName := []string{
		"Vivek",
		"Vishal",
		"Mishra",
	}
	fmt.Println(strings.Join(fullName, "-"))

	banned := "My vame is Vivek Vishal vame"
	fmt.Println(strings.Replace(banned, "vame", "name", 1)) // if -2 ( -ve number) replace all occurence
	fmt.Println(strings.ReplaceAll(banned, "vame", "name")) // all occurrences

}

// the standard lib in go has many handy packages such as strings package

//ToLower, ToUpper, Trim, TrimRight, TrimLeft, and TrimSpace;
//Splitting a string and joining slices of strings into a new string; strings.Split()
//
//Replacing a substring from a string with the use of the Replace and ReplaceAll functions.
