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
	done := make(chan bool) // unbuffered: meaning blocking

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true // sends to unbuffered channel done
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

	<-done // will be blocking until someone sends to it
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

/*
chad gpt explains the code:
In this code, the done channel is used to signal when all jobs have been received and processed. Let me explain how the code works:

    It creates two channels: jobs and done.
    A goroutine is started with an anonymous function, which continuously listens on the jobs channel for incoming jobs. If there are more jobs to receive (more is true), it prints a message indicating the received job. If there are no more jobs to receive (more is false), it prints a message indicating that all jobs have been received and sends a signal on the done channel.
    In the main function, a loop sends three jobs to the jobs channel and prints a message for each sent job.
    After sending all jobs, the jobs channel is closed, indicating that no more jobs will be sent.
    It prints a message indicating that all jobs have been sent.
    Finally, it waits for a signal on the done channel using <-done. This is a blocking operation, meaning the program will wait at this line until it receives a value from the done channel. Once it receives a value (in this case, true), it proceeds to the end of the main function and terminates the program.

So, <-done is used to synchronize the main goroutine with the goroutine processing the jobs, ensuring that the main goroutine waits until all jobs have been processed before exiting.


*/
