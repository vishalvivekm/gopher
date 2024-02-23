### Goroutine: A function that can run concurrently with other similar functions in the same address space.
- They start with an initial stack capacity of just 2KB.
- They support concurrency models.
- They are managed by Go scheduler.
- main function simply starts all the goroutines and continues its execution without waiting for them, unless we say it to do so.
```yamls
Run in concurrency with other functions. // nope
The order of execution of goroutines is not guaranteed. // yes
Run in concurrency with other goroutines. // yes
```
- by default goroutines aren't executed in the order they're started
- when main function exits all the goroutines called in main also exit irrespective of whether they finished executing or not.


```go
package main

import (
    "fmt"
    "time"
)

func numberQuad(i int) int {
    return i * i
}

func main() {
    result := 1_000_000
    go func() {
        result = numberQuad(2)
    }()
    time.Sleep(time.Millisecond)
    // Try to comment the line with time.Sleep
    fmt.Println(result)
}
```
methods of getting values from function runnign as goroutine
// go quad := numberQuad(2) // wrong error

```go
package main

import "fmt"

func workWithChannel(i int, quad chan int) {
    quad <- i * i
}

func main() {
    quad := make(chan int)
    go workWithChannel(3, quad)
    result := <-quad
    fmt.Println(result)
}
```
 channels : through channels, goroutines communicate through each other
