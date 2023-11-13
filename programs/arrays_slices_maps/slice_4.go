package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 90, 70, 60}
	slice := append(arr[:2], arr[3:])
	fmt.Println(slice)
}

// investigate the error
