
// In golang, defer statement is used to
// to delay the execution of a function until the surrounding function returns.

package main
import "fmt"

func main () {

	defer fmt.Println("Printed second!")
    fmt.Println("Printed first!")
/*
any statement that is preceded by the defer statement
isn't invoked until the end of the function in which defer was used, here the main().
*/

/* In case of multiple defer statement:
   They are stored and executed as a stack ( First in Last out)
*/
defer fmt.Printf(" Bob") // first deferred statement
defer fmt.Printf("there!") // second deferred statement
defer fmt.Printf("Hey, ") // third deferred statement

// output: Hey, there! Bob
//  each time we defer a statement, we push it onto a stack and then call it out in Last In, First Out (LIFO) order


}