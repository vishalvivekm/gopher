// channels are the communication medium for various goroutiens to exchange data
// here's default unbuffered (blocking) channel:
package main

import (
	"fmt"
	"time"
)

func main() {
	// channel created using var is empty and can't be initialised
	ch := make(chan string) // declare channel which transfers string data type

	go sell(ch)
	go buy(ch)
	time.Sleep(2 * time.Second)

}

// sends data to the channel
func sell(ch chan string) { // channels just like maps and slice are sent by reference implicitely, so no need of & *
	ch <- "Furniture"
	fmt.Println("Sent data to the channel")
}

// receive data from the channel
func buy(ch chan string) {
	fmt.Println("Waiting for data")
	val := <-ch
	fmt.Println("Recieved data - ", val)

}

// operations on channel:
/*
receiving to channel
ch <- value

receive from the channel

val := <-ch // date read from channel will be stored in val variable

close the channel using built-in function close()
ch.close()

 querying buffer of a channel
 cap(ch)

 querying length of a channel
 len(ch)



*/

// channels by default when initialised like make(ch chan string) is unbuffered and is blocking, means is sequential in the essence that if we're sending to channel, we'll be blocked until is recieved by other goroutine and vice-versa.
//a sent on an unbuffered channel is blocked until a receive happens on that channel in some other go-routine
