package main

import (
	"fmt"
	"time"
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

	ch := make(chan int)
	go consume(ch)

	// go func() {
	for i := 0; i < 5; i++ {
		go func() { ch <- i }()
	}
	//}()
	time.Sleep(1 * time.Second)
}

