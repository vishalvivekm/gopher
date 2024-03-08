package main

import (
	"fmt"
	"io"
	"strings"
)

/*
' in string package :

type Reader struct {
s string
i int
prevRune
}
*/
func main() {

	r := strings.NewReader("Readfromthisokay")

	p := make([]byte, 5) // 5 - chunksize
	//n, err := r.Read(p)
	//if err != nil {
	//	fmt.Errorf("%v", err)
	//}
	//fmt.Println(n)
	//fmt.Println(string(p)) // Readf

	//for breakOutOfLoop := false; breakOutOfLoop != true; {
	//
	//	_, err := r.Read(p)
	//
	//	switch err {
	//	case nil:
	//		fmt.Println(string(p))
	//	case io.EOF:
	//		fmt.Println("done reading the string")
	//		breakOutOfLoop = true
	//	default:
	//		fmt.Println("error occured", err)
	//	}
	//
	//}

	for {
		n, err := r.Read(p)

		switch err {
		case nil:
			fmt.Println(string(p[:n]))
		case io.EOF:
			fmt.Println(string(p[:]))
			fmt.Println("done reading the string")
			return
		default:
			fmt.Println("error occured", err)
		}

	}

}
