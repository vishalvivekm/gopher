package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		if time.Since(start) > time.Second*2 {
			fmt.Println("you're late buddy")
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

/*
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	start := time.Now()
	defer func() {
		if time.Since(start) > time.Second*2 {
			fmt.Println(10000)
		}
	}()

	result := [5]int{}
	for i := 0; i < 5; i++ {
		// runing the code in the loop concurrently
		go func(i int, s *sync.WaitGroup) {
			value := cube(i)
			result[i] = value
			s.Done()

		}(i, &wg)
	}
	time.Sleep(time.Second)
	wg.Wait()

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


*/
