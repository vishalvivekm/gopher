package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i) // (&{Hello}, *main.T)
	i.M()

	i = F(math.Pi)
	describe(i) // (3.141592653589793, main.F)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
// (value, type)
// Calling a method on an interface value executes the method of the same name on its underlying type
