package main

import (
	"flag"
	"fmt"
)

func main() {
	// Declare string and int flags with default values and help messages:
	name := flag.String("name", "User", "Enter your name")
	age := flag.Int("age", 1, "Enter your age")

	// Another way to declare a flag - bind a command-line flag to an existing variable:
	var spacing bool
	flag.BoolVar(&spacing, "spacing", true, "Insert a new line between each message")

	// After declaring all the flags, enable command-line flag parsing:
	flag.Parse()

	fmt.Printf("Hello, %s! ", *name)
	if spacing {
		fmt.Println()
	}
	fmt.Printf("You are %d years old.\n", *age)
}

// ./flags --name vivek --age 10 --spacing=true
// ./flags --name=vivek --age=10 --spacing=false // mendatory to use = with bool
// ./flags -name vivek -age 10 -spacing=false // one or two -
// ./flags -h or --help

/*
Usage of ./flags:
  -age int
        Enter your age (default 1)
  -name string
        Enter your name (default "User")
  -spacing
        Insert a new line between each message (default true)


*/

/* difference between declaring a flag and binding a flag to an existing variable

Declaring: flag.Type(...)
Binding: flag.TypeVar(...)
age := flag.Int("age", 1, "Enter your age")  // declaring an int flag

var age int
age = flag.IntVar("age", 1, "Enter your age") // binding: bind a command-line flag age to the variable age of the int type
*/

/*
os.Args variable:
The os.Args variable is a slice of strings.
All the arguments we pass to os.Args are always treated as string type values.
Its first value (os.Args[0]) is the name of the Go program.
*/
