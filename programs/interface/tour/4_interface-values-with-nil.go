package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T // pointer's zero value is nil
	i = t
	describe(i)// {<nil>, nil}
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*Note that an interface value that holds a nil concrete value is itself non-nil. */
// here in both cases, i's type is : *main.T // fmt.Printf("%T\n", i)
