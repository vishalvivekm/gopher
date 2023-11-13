package main

import "fmt"

func main() {
	var x, y int = 100,9
	x /= y
	fmt.Println(x)
	x %= y
	fmt.Println(x)
}
