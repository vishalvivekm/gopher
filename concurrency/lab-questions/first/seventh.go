// What happens if you read from an unbuffered, closed channel? 
// returns zero value
package main

import (
	"fmt"
	"time"
)

func main() {

ch := make(chan int)
go func() {
ch <- 3
close(ch)
}()
go func() {
val, ok := <- ch
fmt.Println(val, ok)
}()
time.Sleep(time.Second * 2)
}


/*

he correct answer is Returns zero value.

Explanation:

Use the second Boolean return parameter to confirm if the channel is closed.

For an example, over here:

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
}

func buy(ch chan int) {
    for i := 1; i <= 3; i++ {
        val, ok := <-ch
        fmt.Println(val, ok)
    }

}


The sell goroutine writes the values 1 and 2 to the channel and then closes it.
The buy goroutine reads the values from the channel in a loop of 3 iterations and prints the value and a boolean value indicating if the read value is valid or not.
In the third iteration, the value of val would be 0 and ok would be false.

If you run the program, you would get this output:-

> go run main.go

1 true
2 true
0 false


Hence, if you read from an unbuffered, closed channel, it returns the zero value of the channel.
Here, the channel was of type int, so we got 0 as the output.

*/
