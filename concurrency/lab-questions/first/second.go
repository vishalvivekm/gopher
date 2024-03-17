// what would the output?
package main

import (
	"fmt"
	"sync"
)

func consume(ch chan int) {
    for {
        select {
        case num := <-ch:
            fmt.Printf("%d ", num)
            break
        }
    }
	
}

func main() {
	// var wg *sync.WaitGroup
	// wg.Add(1)
    ch := make(chan int) // unbuffered channel accepting int
    go consume(ch)

    for i := 0; i < 5; i++ {
        ch <- i
    }

}
// output will either be 0 1 2 3 4 or 0 1 2 3

// KK expaination: The answer is Nondeterministic since we did not use wait groups, and the main function can exit at any point in time, without waiting for the goroutine to finish.
// Had we used the goroutine, 0 1 2 3 4 would be the right answer.

// there will be race condition between main goroutine and consume one for the 4 to be printed out, before that due to blocking nature for-select will print the integers in the order recieved from the main goroutine
