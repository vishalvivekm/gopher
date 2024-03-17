package main
import (
    "sync"
)
func main() {
    ch := make(chan int)
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        ch <- 1
		close(ch)
        wg.Done()
    }()
    
    
    <-ch
	wg.Wait()
}

/*
predict behaviour of this program:
package main
import (
    "sync"
)
func main() {
    ch := make(chan int)
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        ch <- 1
        wg.Done()
    }()
    wg.Wait()
    close(ch)
    <-ch
}

*/

//hint: 
/*The correct answer would be The program ends in a deadlock, causing a fatal error.

Because the go-routine gets blocked on the following line: -

ch <- 1

Causing the go-routine to never exit. Hence, causing a deadlock.*/
