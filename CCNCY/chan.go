package main

import ( 
	"fmt"
        "time"
)

func main() {
 ch := make( chan string) // declare channel which transfers string data type

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
 val := <- ch
fmt.Println("Recieved data - ", val)

}
