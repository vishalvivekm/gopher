package main

import "fmt"

/*passing by references
func modify(numbers ...int) {
	for i := range numbers {
			numbers[i] -= 5
	}
}
func main() {
	arr := []int{10, 20, 30} // a slice
	fmt.Println(arr)
	modify(arr...)
	fmt.Println(arr)
}
*/
// passing by value

func modify(numbers [3]int) {
	for i := range numbers {
		numbers[i] -= 5
	}
}
func main() {
	arr := [3]int{10, 20, 30}
	fmt.Println(arr)
	modify(arr) //cannot use arr (type [3]int) as type *[3]int in argument to modify, to fix: modify(&arr)
	// modify(&arr) , then signature changes to: func modify(numbers *[3]int ) {
	fmt.Println(arr)

}

/*
The ... notation after arr tells Go to "unpack" the elements of the arr slice and pass them as separate arguments to the modify function.

So, if arr is [10, 20, 30], modify(arr...) is equivalent to calling modify(10, 20, 30). It passes each element of the slice as a separate argument to the modify function*/

/*
func modify(s *string) {
	*s = strings.ToUpper(*s)
}
func main() {
	s := "hello"
	fmt.Println(s)
	modify(&s)
	fmt.Println(s)
}

*/

/*  maps, slices, and channels are reference types in Go, and their behavior is different from value types like integers or structs, which are copied when passed to functions.

func modify(s map[string]int) {
	s["A"] = 100
}
func main() {
	ascii_codes := map[string]int{}
	ascii_codes["A"] = 65
	fmt.Println(ascii_codes)
	modify(ascii_codes)
	fmt.Println(ascii_codes)
}
output:
map[A:65]
map[A:100]

*/
/*
when passing a map to a function or assigning it to another variable, we're actually passing and working with a reference to the underlying map data structure => changed made to a map within a func will affect te original map outside the func
*/
