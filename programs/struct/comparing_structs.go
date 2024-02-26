package main

import "fmt"

type Person struct {
	name string
	age  rune
}

func main() {

	// multiple ways of initializations of the Person struct
	//Person_1 := new(Person) // the new func allocates memory for all the fields of the struct, sets them to their zero value, and returns a pointer to the struct (*Person).

	//Person_1.name = "Vivek"
	//Person_1.age = 20

	//person_1 := Person{
	//	"vivek",
	//	20,
	//}

	//var person_1 = Person{"vivek", 20}

	//var person_1 = Person{
	//	name: "Vivek",
	//	age: 20,
	//
	//}

	Person1 := Person{"Vivek Vishal", 20}
	Person2 := Person{"Vivek Vishal", 20}
	Person3 := Person{name: "Vivek Vishal"}

	fmt.Println(Person1 == Person2) // true
	fmt.Println(Person1 == Person3) // false, is of Person type only, but age isn't initialized and so isn't the same

	// Also struct type variables are not comparable if they contain fields such as maps or slices that can't be compared using the equality == operator.

}

// resource: https://asankov.dev/blog/2022/01/29/different-ways-to-initialize-go-structs/
