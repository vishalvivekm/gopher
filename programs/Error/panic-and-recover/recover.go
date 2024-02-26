/*
A panic is an emergency end of the program. An unexpected end can lead to corrupted files or databases. Naturally, it's not great when wer program starts panicking as part of the data processing server. we can't prevent the panic, but we can reduce the possible damage (close the file, break the connection to the database, switch to an alternative data handler).

At the end of the previous section, we saw that panic doesn't cancel the defer statement. To process the panic, we need to get the signal about it inside the defer scope. Let's return to the first example of the panic section. The code below catches the zero divide condition and prints a custom message about it.
*/
package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	defer func() {
		onPanic := recover() // catch the panic
		if onPanic != nil {
			fmt.Printf("%d and %d are unacceptable for integer division\n", num1, num2)
		}
	}()

	fmt.Println(num1 / num2)
}

// Input:
// 1 0
// Output:
// 1 and 0 are unacceptable for integer division

/*Here the defer statement goes after the fmt.Scan(&num1, &num2) line. This is necessary to pass the scanned values of num1 and num2 to the fmt.Printf() statement of the deferred anonymous function.

The program still panics, but now we control the process! Now we get a meaningful error message. Also, we can log this error message to a file, for example. Or we can handle the error and try to execute the program again. The method of ending the program after a serious error during runtime is called a Graceful exit.

we can benefit from a graceful exit after a panic in Go by using the recover() function. The recover() function gets the signal of the panic and returns the information about error that occurred during runtime.
*/
