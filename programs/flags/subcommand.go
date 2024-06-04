package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	repeat := flag.NewFlagSet("repeat", flag.ExitOnError)
	repeatName := repeat.String("name", "User", "Enter the name to be repeated")
	repeatCount := repeat.Int("count", 1, "Enter the number of repetitions")

	if os.Args[1] == "repeat" {
		repeatSubCommands := os.Args[2:]
		if len(repeatSubCommands) == 0 {
			fmt.Println("Repeats name based on given count")
			return
		}
		repeat.Parse(os.Args[2:])

		for i := 1; i <= *repeatCount; i = i + 1 {
			fmt.Println(*repeatName)
		}
	} else {
		log.Fatalln("expected repeat and subcommands")
	}
}
