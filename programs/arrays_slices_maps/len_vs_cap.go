package main

import "fmt"

func main() {
	var myArr [5]int = [5]int{123, 124, 125, 126}
	fmt.Println(len(myArr))
	// fmt.Println(myArr) // [123 124 125 126 0]
	fmt.Println(cap(myArr))
}

// in case of arrays length and capacity is the same.
