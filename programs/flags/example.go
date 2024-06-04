package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	reverse := flag.NewFlagSet("reverse", flag.ExitOnError)
	reverseName := reverse.String("name", "User", "Enter name to be reverses")

	repeat := flag.NewFlagSet("repeat", flag.ExitOnError)
	repeatName := repeat.String("name", "user", "Enter name to be repeated")
	repeatCount := repeat.Int("count", 1, "enter no of repetitions")

	if len(os.Args) < 2 {
		log.Fatalln("either repeat or reverse command expected")
	}
	switch os.Args[1] {
	case "reverse":
		reverse.Parse(os.Args[2:])
		runes := []rune(*reverseName)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		fmt.Printf("%s\n", string(runes))
	case "repeat":
		repeat.Parse(os.Args[2:])
		name := *repeatName
		count := *repeatCount
		for i := 1; i <= count; i = i + 1 {
			fmt.Println(name)
		}
	default:
		log.Fatalln("only repeat and reverse commands are supported")

	}

	//repeat.Parse(os.Args[2:])
	//fmt.Println(countName)
	//
	//fmt.Println(*reverseName)
	//fmt.Println(*repeatName)

}

/*
	nameString := "Shashwat"
	name := []rune(nameString)

	for i, j := 0, len(name)-1; i < j; i, j = i+1, j-1 {
		name[i], name[j] = name[j], name[i]
	}
	// can't use nameString instead of name here in for loop
	fmt.Println(nameString)
	fmt.Println(string(name))
*/
/*
package main

import (
	"flag"
	"fmt"
)

func main() {
	reverse := flag.NewFlagSet("reverse", flag.ExitOnError)
	reverseName := reverse.String("reverseName", "User", "Enter name to be reverses")

	repeat := flag.NewFlagSet("repeat", flag.ExitOnError)
	//var s string
	//repeat.StringVar(&s, "s", "user", "enter name ")


	//flag.Parse()
	//fmt.Println(reverseName) // a pointer
	//fmt.Println(*reverseName) // User, a value
	//fmt.Println(s) // user : value
}

*/
