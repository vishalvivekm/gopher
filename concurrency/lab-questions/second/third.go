package main

import "fmt"

func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    close(ch)
    for n := range ch {
        fmt.Println(n)
    }
}

/*1
2



Explanation:
1
2
The code creates a buffered channel ch with a buffer size of 2 and sends the values 1 and 2 to the channel. It then closes the channel. The main function then range over the channel and prints the received values. Since the channel has a buffer size of 2 and the values 1 and 2 were sent to the channel before it was closed, the main function will receive and print both values.*/
