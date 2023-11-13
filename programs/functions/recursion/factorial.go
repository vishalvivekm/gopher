package main

import "fmt"

func calcFactorial(n int) int {
	if n ==1 {
       return 1 // could not mention 1 here if return type int not specified in func signature
	}
 return n * calcFactorial(n-1)
}

func main() {
fmt.Println(calcFactorial(5))
}
