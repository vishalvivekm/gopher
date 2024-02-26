// go-routine leak: a go-routine that would never terminate, forever occupies the memory it's reserved, this kind of memory leak is called go-routine leak.

// go routines leak if they end up either blocked forever on I/O like channel commn or fall into inf loops.


package main



import (
"fmt"
"sync"
)

func main() {

var wg sync.WaitGroup
wg.Add(2)
go leak(&wg)
wg.Wait()

}
func leak( s *sync.WaitGroup) {
 ch := make(chan int)

 go func() { // wait forever
  val := <- ch
  fmt.Println("received from channel: ", val)
  s.Done()
 }()

fmt.Println("Exiting leak method")
s.Done()
}
/*

Exiting leak method
fatal error: all goroutines are asleep - deadlock!



*/
// make sure the go-routines eventually exit to avoid the memory-leaks in the programs.
