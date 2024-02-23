package main

import (
"fmt"
"time"
)

func main() {

ch1 := make(chan string)
ch2 := make(chan string)

go goOne(ch1)
go goTwo(ch2)

select {
  case val1 := <- ch1:
    fmt.Println(val1)
  case val2 := <- ch2:
    fmt.Println(val2)
  default:
    fmt.Println("Executed default block")
}
time.Sleep(1 * time.Second)
}

func goTwo(ch2 chan string){
ch2 <- "Channel-2"

}
func goOne(ch1 chan string){
ch1 <- "Channel-1"

}
// output:
// channel-2, but is non-deterministic 
// select statement blocks until any of the go-routines executes

// to make the select non-blocking: default case, which will execute if all the other cases are blocked

// to make the select non-blocking: default case, which will execute if all the other cases are blocked i.e default case will be executed if no send or receive opearation is ready on any of the case statements.
