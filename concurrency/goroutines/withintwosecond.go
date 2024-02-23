package main

import (
	"fmt"
	"time"
)

func cleanUp() {
	fmt.Println("done")
}
func doCleanUps() {
	for range [5]struct{}{} {
		go cleanUp()
	}
}

func main() {
	start := time.Now()
	defer func() {
		if time.Since(start) > time.Second*2 {
			fmt.Println("you're late!")
		}
	}()

	doCleanUps()
	time.Sleep(time.Second)
}

// do cleanup within 2 second
