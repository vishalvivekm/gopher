package main

import "fmt"

func doSomething(int, int) int {
	return 2
}

func main() {
	fmt.Println(doSomething(1, 2))
}
// ouput: 2
