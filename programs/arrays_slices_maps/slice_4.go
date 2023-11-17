package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 90, 70, 60}
	slice := append(arr[:2], arr[3:])
	fmt.Println(slice)
}

// investigate the error
// remember append func signature
// func append(s []T, vs ...T) []T // returns a slice and takes as arguments a slice and the elements
// func append(slice, element1, element2, ...)

/* cannot use arr[3:] (value of type []int) as int value in argument to append
Hint: wanna append a slice to anther slice use this ``` destinationSlice = append(destinationSlice, sourceSlice...)
*/
