// buffered doesn't block another go routines until we cross the buffer limit
package main

import (
"fmt"
"sync"
)


func main() {
 var wg sync.WaitGroup
 wg.Add(2)
 ch := make(chan int, 3) // buffered channel with cap 3 
 go sell(ch, &wg)
 wg.Wait()

}
func sell(ch chan int, s *sync.WaitGroup) {
ch <- 10
ch <- 11
ch <- 12
//close(ch) : don't forget to use close() while using for range in receiver goroutine

//blocking line:  // ch <- 13 fatal error: all goroutines are asleep - deadlock!, if go buy() is invoked above this , no deadlock
go buy(ch, s)
fmt.Println("Sent all the data to the channel")
s.Done()
}

func buy(ch chan int, s *sync.WaitGroup ) {

fmt.Println("Waiting for data")
fmt.Println("received from the channel: ", <-ch)
//for cal := range ch {
 //fmt.Println(cal)
//}
s.Done()

}
// while sending if the buffer limit is exceeded or while receiving from a empty channel - block, deadlock

