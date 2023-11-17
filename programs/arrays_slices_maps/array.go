package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Array example")
	var arr [4]int = [4]int{1, 2, 43, 5} // Array declaration and initialization
	fmt.Println(arr)
	// another way: 1
	//friends := [...] string{"Tom", "Bob", "John"}
	//fmt.Println(friends)
	// another way: 2
	names := [5]string{"Vivek", "Vishal", "Gaurav", "Lakshay", "Shashwat"}
	i := 0
	for i < len(names) {
		_, err := strconv.Atoi(names[i])
		if err != nil {
			fmt.Printf("Type of %v is:  %T", err, err)
			fmt.Println()
			fmt.Println(err)
		}
		fmt.Println()
		i++
	}
	for index, _ := range names {
		names[index] = strconv.Itoa(index)
	}
	fmt.Println(names)

}

// len(array)
// cap(array)
