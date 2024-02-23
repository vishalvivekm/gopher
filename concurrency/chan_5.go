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
// ch <- 12 // panic: send on closed channel

close(ch) // panic: close of closed channel 
}

// buffered doesn't block another go routines until we cross the buffer limit

