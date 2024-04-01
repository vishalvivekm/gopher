package main

import (
	"flag"
	"fmt"
)

type stringFlag struct {
	value        string
	isSet        bool
	defaultValue string
}

func (sf *stringFlag) String() string {
	return sf.value
}
func (sf *stringFlag) Set(value string) error {
	sf.value = value
	sf.isSet = true
	return nil
}
func StringFlag(name string, value string, usage string) *stringFlag {
	f := stringFlag{defaultValue: value}
	flag.CommandLine.Var(&f, name, usage)
	return &f
}
func main() {
	name := StringFlag("name", "User", "User name")
	flag.Parse()
	if name.isSet {
		fmt.Println("Flag was set by the user.")
	} else {
		fmt.Println("Flag was not set by the user.")
	}
}
