package main

import (
"fmt"
"sync"
)


func main() {
/* ch := make(chan int, 5)
ch <- 1
ch <- 2
ch <- 3

close(ch)
for val := range ch { // while using for range always make sure to close the channel
 fmt.Println(val, " <- ch")
}*/

 var wg sync.WaitGroup
wg.Add(2)
ch := make(chan int, 3)
go sell(ch, &wg)
wg.Wait()

}
func sell(ch chan int, s *sync.WaitGroup) {

ch <- 1
ch <- 2
ch <- 3

close(ch)

 go buy(ch, s)
 for i := 0; i < 5999999; i++ {
 //  

 }
fmt.Println("send data to the channel")
// go buy(ch, s)
s.Done()
}
func buy(ch chan int, s *sync.WaitGroup) {
fmt.Println("Waiting for data from the channel: ")
// for range to iterate through channel
for val := range ch {
fmt.Println("Received from channel: ", val)
}

fmt.Println("Received all data from the channel")
s.Done()
}
