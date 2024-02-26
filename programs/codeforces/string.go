package main

import (
	"fmt"
)

func constructString(trace []int) string {

}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		trace := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&trace[j])
		}

		result := constructString(trace)
		fmt.Println(result)
	}
}
