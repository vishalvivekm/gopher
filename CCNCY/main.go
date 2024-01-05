package main

import ( 
	"fmt"
	"time"
)

func calculateSq(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i*i)

}

func main() {

	start := time.Now();
	 for i:=1; i<=10000; i++ {
		go calculateSq(i);
	}
	// fmt.Println("hello")
        elapsed1 := time.Since(start)
        time.Sleep(2 * time.Second)
	elapsed := time.Since(start)
	fmt.Println("Function took: ", elapsed, "\t elapsed1: ", elapsed1)
}
