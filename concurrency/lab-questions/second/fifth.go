package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	jobs := make(chan int)
	results := make(chan int)

	go func() {
		for j := range jobs {
			fmt.Println("received job", j)
			result := j * 2
			results <- result
		}
		// your code goes here
		close(results)
		wg.Done()

	}()

	go func() {
		for j := 1; j <= 3; j++ {
			jobs <- j
			fmt.Println("sent job", j)
		}
		// your code goes here
		close(jobs)
		wg.Done()
	}()

	go func() {
		for r := range results {
			fmt.Println("received result", r)
		}
		// your code goes here
		wg.Done()
	}()

	wg.Wait()
}

/*i
problem:
In this question, we are working on a project that involves processing a large amount of data in parallel using goroutines. We want to use channels to communicate between the goroutines and ensure that all goroutines have finished processing before the program exits.

Complete the code, so that the program runs successfully without any errors.


A Go file is located at /root/code/code05 directory.

package main

import (
    "fmt"
    "sync"
)

func main() {

    var wg sync.WaitGroup
    wg.Add(3)
    jobs := make(chan int)
    results := make(chan int)

    go func() {
        for j := range jobs {
            fmt.Println("received job", j)
            result := j * 2
            results <- result
        }
        // your code goes here
    }()

    go func() {
        for j := 1; j <= 3; j++ {
            jobs <- j
            fmt.Println("sent job", j)
        }
        // your code goes here
    }()

    go func() {
        for r := range results {
            fmt.Println("received result", r)
        }
        // your code goes here
    }()

    wg.Wait()
}


Expected Output:

sent job 1
received job 1
received job 2
sent job 2
received result 2
received result 4
received job 3
received result 6
sent job 3

*/
/*
package main

import (
    "fmt"
    "sync"
)

func main() {

    var wg sync.WaitGroup
    wg.Add(3)
    jobs := make(chan int)
    results := make(chan int)

    go func() {
        for j := range jobs {
            fmt.Println("received job", j)
            result := j * 2
            results <- result
        }
        close(results)
        wg.Done()
    }()

    go func() {
        for j := 1; j <= 3; j++ {
            jobs <- j
            fmt.Println("sent job", j)
        }
        close(jobs)
        wg.Done()
    }()

    go func() {
        for r := range results {
            fmt.Println("received result", r)
        }
        wg.Done()
    }()

    wg.Wait()
}



You can use the Done() method of sync.WaitGroup to signal the end of Goroutine execution.
And close() method to close the channel, to allow the exit of the range loop.


*/
