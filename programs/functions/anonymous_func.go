<<<<<<< Updated upstream
// anonymous or function literal - function that do not have a defined name

// In Go, there are no nested named functions.

package main 
import "fmt"

func main() {

}
=======
package main

import "fmt"

func main() {

	x := func(a int, b int) int {
		return a * b
	}
	fmt.Println("---------------------------------------x---------------------------------------")
	fmt.Printf("Type of x is : %T\n", x) // Type of x is : func(int, int) int
	fmt.Println("Value of x(10,20): ", x(10, 20))
	fmt.Println("---------------------------------------X---------------------------------------")
	X := func(a int, b int) int {
		return a * b
	}(10, 24)
	fmt.Printf("Type of X : %T\n", X)
	fmt.Println("X: ", X)





/* correct ways to define an ano func
	    var (
                my_func = func(s string) { fmt.Println("Hey there,", s) }
        )
        my_func("Joe")
var (my_func1 = func(s string) {fmt.Println("Hey there,", s)})
}
*/
/*incorrect ways : 
func main() {
        var (
                my_func = func(s string) { fmt.Println("Hey there,", s) }
        )
        my_func("Joe")
}

 my_func func := (func(s string) {fmt.Println("Hey there,", s)})
*/
}
>>>>>>> Stashed changes
