package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup
    // Line 1
    go func() {
        time.Sleep(time.Millisecond)
        fmt.Println("Inside Goroutine")
        // Line 2
    }()
    // Line 3
}

// wg.Add(1)
// wg.Done()
// wg.Wait()
