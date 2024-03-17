package main

import (
"fmt"
"sync"
)

func main() {

var wg sync.WaitGroup
wg.Add(10)

/*	for i := 1; i <= 10; i++ {

	go func() {
	fmt.Println(i)
	wg.Done()
	}()

	}
*/

for i := 1; i <= 10; i++ {

go func(i int) {
fmt.Println(i)
wg.Done()
}(i)

}


wg.Wait()
fmt.Println("Done")
}




/*	for i := 1; i <= 10; i++ {

	go func() {
	fmt.Println(i)
	wg.Done()
	}()

	}

here 10 goroutines are spawned
we'd expect that this will print 1 to 10 ( in random order as go routines aren't non-deterministic)
but result is not quite the same, we might get all 11s or 7 11s and 3 8s :)

// reason: parallelism and concurrently running are two different things. Though they're ready to run,  go-routines execution might get delayed by the go runtime-scheduler and it might happen that when the go-routines started running the i's value was 11 and all of go routines 11 as i's value printing 11 10 times.

output:
11
11
11
11
11
11
11
11
11
11

// soln: pass i as arg to closure
*/
