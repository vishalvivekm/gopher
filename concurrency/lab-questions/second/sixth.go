/*
Which of the following statements is true about the behavior of a channel in Go when it is closed and has no more values to receive?
1. A receive operation on a closed channel will always return an error
2. A receive operation on a closed channel will always return the zero value of the channel's type
3. A receive operation on a closed channel will always block definitely
4. A receive operation on a closed channel will always return the default value of the channel's type

*/
// ans: 
/*

The correct answer is A receive operation on a closed channel will always return the zero value of the channel's type.

Explanation: -


In Go, a channel can be closed using the close function.

Once a channel is closed, it cannot be used to send any more values, but it is still possible to receive values from the channel.

If a channel has no more values to receive and it is closed, a receive operation on the channel will return the zero value of the channel's type without blocking.

If the channel is not closed and there are no values to receive, a receive operation will block until a value is sent on the channel. 
*/
