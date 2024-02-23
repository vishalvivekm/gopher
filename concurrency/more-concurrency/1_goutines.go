package main

import (
	"fmt"
	"time"
)

func sayHello(num int) {
	fmt.Println("Say Hello!: ", num)
}

func main() {
	// startTime := time.Now()
	for i := 0; i < 3; i++ {
		go sayHello(i)
	}
	// elapsed := time.Since(startTime)
	time.Sleep(time.Millisecond) // try commenting this
	fmt.Println("main finished")
	// fmt.Println(elapsed)
}
