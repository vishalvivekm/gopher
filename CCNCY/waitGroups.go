package main

import ( 
	"fmt"
        "sync"
	"time"
)

func calculateSq(i int, wg *sync.WaitGroup) {
	fmt.Println(i*i)
        wg.Done() // reduces internal counter by 1
}

func main() {

          var wg sync.WaitGroup // waitGroup Variable, internal Count = 0;
	start := time.Now();
wg.Add(10); // set the interal counter to 10 ( we've 10 go routines to wait for )
	 for i:=1; i<=10; i++ {
		go calculateSq(i, &wg); // pass waitGroup variable as reference
	}
	// fmt.Println("hello")
        // time.Sleep(2 * time.Second)
         wg.Wait() // wait until the internal counter reduces to 0
	elapsed := time.Since(start)
	fmt.Println("Function took: ", elapsed) 
}
