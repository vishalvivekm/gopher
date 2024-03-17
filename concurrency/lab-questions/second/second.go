/*
Which of the following is NOT a correct statement about the Golang runtime scheduler?
check the image: 
ans; 
*/

/*
it's preemptive scheduler

    The Golang runtime scheduler is a preemptive scheduler, meaning it can interrupt the execution of a goroutine and allow another goroutine to run.

    The scheduler uses a priority-based algorithm to schedule goroutines, with higher priority goroutines being given more CPU time.

    The scheduler can move goroutines between OS threads as needed to ensure optimal utilization of CPU resources.

    The scheduler allows goroutines to voluntarily yield their time slice by calling the runtime.Gosched() function.

    The scheduler does not use a round-robin algorithm to schedule goroutines. Instead, it uses a priority-based algorithm as mentioned in the 2nd option above.

*/
