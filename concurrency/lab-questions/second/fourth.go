/*
statement: 
Complete the code here, for the program to work successfully and give the following response with no error:

sent job 1
sent job 2
sent job 3
sent all jobs
received job 1
received job 2
received job 3
received all jobs


A Go file is located at /root/code/code04 directory.

package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }

    // your code goes here
    fmt.Println("sent all jobs")

    <-done
}





The main idea of this question is to test users ability to start and close jobs correctly.


As long as the job is closed successfully, the validation should pass for any given number of sent and received jobs.


*/


package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	// your code goes here
	close(jobs)

	fmt.Println("sent all jobs")

	<-done
}

/*
Close the channel after sending all the jobs.

package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }

    // close the channel
    close(jobs)
    fmt.Println("sent all jobs")

    <-done
}


This will allow the goroutine that is receiving values from the jobs channel to exit the loop when it receives a value with the more flag set to false.
*/
