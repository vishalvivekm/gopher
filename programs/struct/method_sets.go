// set of methods that are available to a data type.
package main

import "fmt"

type Student struct {
	name   string
	roll   int
	grades []int
}

func (s *Student) displayName() {
	fmt.Println(s.name)
}
func displayName(s *Student) {
	fmt.Println("name from function: ", s.name)
}
func (s *Student) calculatePercentage() float64 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float64(sum*100) / float64(len(s.grades)*100)
}
func main() {
	st := Student{name: "Vvm", grades: []int{40, 40, 50}}
	// st := Student{"Vvm", []int{40, 40, 50}} // wrong
	st.displayName() // calling method related to struct st
	var s *Student = &st
	fmt.Println("name from pointer", s.name)
	// fmt.Println("name from pointer", (*s).name) // see below
	displayName(&st) // calling function displayName, which takes address of the struct
	fmt.Println(st)
	fmt.Printf("%.2f\n", st.calculatePercentage())
}

/*


Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X.
However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

*/
