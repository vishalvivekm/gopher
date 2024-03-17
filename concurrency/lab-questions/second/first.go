package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	letters := []string{"a", "b", "c", "d", "e"}
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers(nums, &wg)
	go printLetters(letters, &wg)

	wg.Wait()
}

func printNumbers(nums []int, s *sync.WaitGroup) {
	for _, val := range nums {
		fmt.Println(val)
	}
	s.Done()
}

func printLetters(letters []string, s *sync.WaitGroup) {
	for _, val := range letters {
		fmt.Println(val)
	}
	s.Done()
}

/*

package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go printNumbers(&wg)
    go printLetters(&wg)

    // Wait for both goroutines to finish
    wg.Wait()
}

func printNumbers(wg *sync.WaitGroup) {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
    wg.Done()
}

func printLetters(wg *sync.WaitGroup) {
    for i := 97; i <= 101; i++ {
        fmt.Println(string(i))
    }
    wg.Done()
}

*/
