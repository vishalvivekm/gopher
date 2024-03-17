package main

import (
    "fmt"
    "sync"
)

func main() {
    var c chan int
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        c <- 1
        wg.Done()
    }()

    go func() {
        val := <-c
        fmt.Println(val)
        wg.Done()
    }()
    wg.Wait()

}
// why this fails

/*
The correct answer would be Sending to and receiving from a "nil" channel blocks forever because the zero value of the channel is nil. Hence only declaring a channel creates a nil channel, as the default zero value of the channel is nil.
Sending to a nil channel blocks forever

Receiving from nil channel blocks forever 
*/
