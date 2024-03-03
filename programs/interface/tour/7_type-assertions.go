package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string) // asserts that interface value i holds the concrete type "string" and assignsthe underlying string value to the variable s here.
	// in this case, yes i ( which is an empty interface) holds the concrete type "string" and so the underylying value ("hello") is assigned to s.

	fmt.Println(s) // prints hello

//	r := i.(float64) // as i doesn't hold a float64 value, this triggers panic: panic: interface conversion: interface {} is string, not float64
//	fmt.Println(r)
	
	s, ok := i.(string) 
	fmt.Println(s, ok) // ("hello", true)

	f, ok := i.(float64) // no panic occurs
	fmt.Println(f, ok) // ("", false) since "" is the zero value of string type in go

/*	f = i.(float64) // panic
	fmt.Println(f) */
}

