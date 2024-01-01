package main

import "fmt"

func returnCube(num int) int {
	return num * num * num
}

func main() {
	fmt.Println(returnCube(5))
	fmt.Println(returnSquare(5))
}

/*
works just fine

	func returnCube(num int) int {
		return num*num*num;
	}

func main() {
fmt.Println(returnCube(5))
}
*/
func returnSquare(num int) int {
	return num * num
}
