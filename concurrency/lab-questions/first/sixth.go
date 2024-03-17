package main

import _"fmt"

func main() {
ch := make(chan int)
close(ch)

ch <- 3

}

// panic: send on closed channel

/*
what happend when you write to unbuffered, closed channel
The correct answer is Panic Situation.

Explanation:

Try creating an unbuffered channel, and then closing it. Observe the behavior when you try to write on the closed channel.

For an example, over here: -

package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int)
    go sell(ch)
    go buy(ch)
    time.Sleep(time.Second * 2) // so main goroutine doesn't exit
}

func sell(ch chan int) {
    ch <- 1
    ch <- 2
    close(ch)
    ch <- 3
}

func buy(ch chan int) {
    for i := 1; i <= 3; i++ {
        val, ok := <-ch
        fmt.Println(val, ok)
    }

}


This code would give the output on line ch <- 3. Run the following go run command to check: -

> go run main.go

1 true
2 true
0 false
panic: send on closed channel

...


And this is because writing to a closed channel creates a panic situation.

*/
