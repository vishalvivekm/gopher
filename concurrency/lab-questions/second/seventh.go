package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        fmt.Println("Goroutine 1")
        wg.Done()
    }()
    go func() {
        fmt.Println("Goroutine 2")
        wg.Done()
    }()
    wg.Wait()
    fmt.Println("Done")
}


// guess the output: 
/*
Goroutine 1
Goroutine 2
Done

or, 

Goroutine 2
Goroutine 1
Done

*/
