package main

import "fmt"

func print(n int) {
	if n == 0 {
			return
	}
	fmt.Print(n)
	print(n - 1)

}

func main() {
	print(5)
}
