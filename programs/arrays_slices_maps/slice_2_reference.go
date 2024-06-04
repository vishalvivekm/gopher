package main

import "fmt"

func main() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	slice := arr[:4]
	fmt.Println(arr)
	fmt.Println(slice)
	slice[1] = "x"
	fmt.Println(arr)
	fmt.Println(slice)

}

//output:
// essence is that : unlike arrays, slice is a reference to a memory address, so changing an element
// in the slice also changes the element in the underlying array
