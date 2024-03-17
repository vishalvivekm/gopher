// usually the interactive program, we want them to give a response within certain amount of time, like getting a response from an API, if they're not, we should throw a time out after a particular time. 1 second in this examples.
// we can achieve this using time package's After method in select statement
package main

import (
"fmt"
_"sync"
"time"
)

func main() {
ch := make(chan int)

go sendValue(ch)

select {
 case val := <- ch:
   fmt.Println(val)
 case <- time.After(1 * time.Second):
   fmt.Println("select time out")
}

} 
func sendValue(ch chan int) {
// time.Sleep(1 * time.Second) // we delay the send, and select statement picks the time.After case
ch <- 124 


}
// select statement waits for at least 1 second and then it throws time out
