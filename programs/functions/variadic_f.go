package main

import "fmt"

// variable number of arguments that are passed into variadic parameters of say X datatype are stored in a slice of X datatype

func printStrings(s string, names ...string) { // variadic parameter of string datatype, has to be put towards the end of the fsig
	fmt.Println(s)

	for _, value := range names { // blank identifier - anonymous placeholder, used to ignore values returned from a function
		fmt.Printf("%s, ", value)
	}
}

func main() {
	printStrings("Hey there", "Joe", "MOnica", "Gunjan")
}
