package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		if time.Since(start) > time.Second*2 {
			fmt.Println(10000)
		}
	}()

	result := [5]int{}
	for i := 0; i < 5; i++ {
		// runing the code in the loop concurrently
		go func(i int) {
			value := cube(i)
			result[i] = value
		}(i)
	}
	time.Sleep(time.Second)

	printResult(result)
}
func cube(i int) (cubed int) {
	cubed = i * i * i
	return
}
func printResult(arr [5]int) {
	for _, ele := range arr {
		fmt.Println(ele)
	}
}
