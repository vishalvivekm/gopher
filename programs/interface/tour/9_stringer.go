// inside fmt package
//
//	type Stringer interface {
//		  String() string
//	}
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
type Man struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age) // Sprintf used to return a formatted string.
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	var i fmt.Stringer = z
	//describe(i) // (Zaphod Beeblebrox (9001 years), main.Person)
	fmt.Println(i.String())

	m := Man{"Vivek", 20}
	fmt.Println(m)

}

func describe(i fmt.Stringer) {
	fmt.Printf("(%+v, %T)\n", i, i)
}
