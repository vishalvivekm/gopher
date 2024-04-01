package main

import (
	"fmt"
	"sort"
)

func main() {
	vars := []int{34, 1, 67, 0, 102}
	sort.Ints(vars)
	fmt.Println(vars)

	str := []string{"Learning", "golang", "Is", "fun"}
	sort.Strings(str)
	fmt.Println(str)
}
