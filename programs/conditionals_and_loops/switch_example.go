package main

import "fmt"

func main() {
	num := 6
	i := num
	switch i {
	case 4:
		fmt.Println("not 6")
		fallthrough
	case 6:
		fmt.Println("yes it's ")
		fallthrough
	case 8:
		fmt.Println("it's 8 and not 6")
		fallthrough
	default:
		fmt.Println("not found")
	}

}

// and so when you use fallthrough it
// it's gonna execute the case after the fallthrough keyword
// output of the above program
// yes it's
// it's 8 and not 6
// not found
