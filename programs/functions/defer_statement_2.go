package main
import "fmt"

// 	// defer statement with multiple functions

func main() {
defer fmt.Println("Printed after the main() function is completed.") // 4
sayingHi()
fmt.Println("Printed after calling the sayingHi() function.") // 3

}
func sayingHi() {
	defer fmt.Println("Printed after Hello, there! Bob") // 2

    fmt.Println("Hello, there! Bob") // 1
}

// output: 
/*
Hello, there! Bob
Printed after Hello, there! Bob
Printed after calling the sayingHi() function.
Printed after the main() function is completed.
*/

/* When a function contains multiple deferred statements, 
the program executes each deferred statement within that function once the function, 
which contains the defer statements, is completed. */

// func scopedDefer() {
// n := 0
// defer func() { fmt.Println("n =", n, "- first deferred print") }()
// {
// 	defer func() { fmt.Println("n =", n, "- second deferred print") }()
// 	n++ // n = 1
// }
// n++ // n = 2
// }
// guess output?
// 
