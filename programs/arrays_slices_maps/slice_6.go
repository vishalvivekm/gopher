package main

import "fmt"

func main() {
	arr := []int{10, 20, 90, 70, 60}
	slice := make([]int, 10)
	copy(slice, arr)
	slice[1] = 1000
	fmt.Println(arr)
	fmt.Println(slice)
}
