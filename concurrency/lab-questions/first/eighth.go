/*Select all the correct options for a nil channel.

a. Read operation blocks forever.


b. Write operation blocks forever.


c. Read operations works and returns a zero value.


d. Close operation creates a panic situation.


ans:

The correct answer is a, b, d.

Explanation:

Try declaring a channel var c chan int.

Read, write and close this channel to observe the behavior.


For an example, over here: -

package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    var ch chan int
    go sell(ch, wg)
    go buy(ch, wg)
    wg.Wait()
}

func sell(ch chan int, wg sync.WaitGroup) {
    ch <- 1
    ch <- 2
    close(ch)
    wg.Done()
}

func buy(ch chan int, wg sync.WaitGroup) {
    for i := 1; i <= 3; i++ {
        val, ok := <-ch
        fmt.Println(val, ok)
    }
    wg.Done()
}


Here, if you run the program, you'll receive the following error: -

fatal error: all goroutines are asleep - deadlock!


And this is because the Read and Write operations on the channel are blocking forever.

Whereas if you run the following program:

package main

func main() {
    var ch chan int
    close(ch)
}


This will return a panic situation error, as expected:

panic: close of nil channel
*/
