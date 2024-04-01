package main

import (
	"flag"
	"os"
)

func main() {

	repeat := flag.NewFlagSet("repeat", flag.ExitOnError)
	repeatName := repeat.String("name", "User", "Enter the name to be repeated")
	countName := repeat.Int("count", 1, "Enter the number of repetitions")

	if os.Args[1] == "repeat" {
		repeat.Parse(os.Args[2:])

	}
}
