package main

import (
	"fmt"
	"time"
)

func main() {
	//ch := make(chan int)
	//
	//fmt.Println("sending data to the channel...")
	//ch <- 2
	//now := time.Now()
	//time.Sleep(3 * time.Second)
	//val := <-ch
	//
	//fmt.Printf("Received data from channel: %v in %v", val, time.Since(now))
	//sending data to the channel...
	//fatal error: all goroutines are asleep - deadlock!

	ch := make(chan int)

	go func() {
		fmt.Println("sending data to the channel...")
		ch <- 2
	}()
	go func() {
		time.Sleep(1 * time.Second)
		val := <-ch

		fmt.Printf("Received data from channel: %v\n", val)

	}()
	time.Sleep(2 * time.Second)
}
