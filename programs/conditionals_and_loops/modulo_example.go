package main

import "fmt"

func main() {
	var x, y int = 100, 9
	{
		x /= y
		fmt.Println(x)
	}
	fmt.Println(x)
	x %= y
	fmt.Println(x)
	fmt.Println(100 % 9)
}

/*
package main

import "fmt"

func main() {
	var x, y int = 100, 9
	{
		x /= y
		fmt.Println(x)
	}
	fmt.Println(x)

}

the block (enclosed in curly braces) does not introduce a new scope in Go
*/
