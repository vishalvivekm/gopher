package main

import (
"fmt"
)


func main() {
ch := make( chan int, 10)
 ch <- 10
 ch <- 11
 val, ok := <-ch
 fmt.Println(val, ok)
 close(ch)

val, ok = <-ch
fmt.Println(val, ok)

val, ok = <-ch
fmt.Println(val, ok)
}
package main

import (
"fmt"
"sync"
)


func main() {
 var wg sync.WaitGroup
 wg.Add(2)
 ch := make(chan int, 3)  
 go sell(ch, &wg)
 wg.Wait()

}
func sell(ch chan int, s *sync.WaitGroup) {
// sending nothing to the channel, channel is empty 

// ch <- 1 // uncomment me and run the program to avoid deadlock

go buy(ch, s) // as channel is empty, buy method is going to be blocked.
fmt.Println("Sent all the data to the channel")
s.Done()
}

func buy(ch chan int, s *sync.WaitGroup ) {

fmt.Println("Waiting for data")
fmt.Println("received from the channel: ", <-ch) // expecting to receive from empty channel: blocking. Errors
s.Done()

}
// while sending if the buffer limti is exceeded or while receiving from a empty channel - block, deadlock

