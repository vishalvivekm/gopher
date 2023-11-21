/*
type <name> struct{
	<variable-name> <type>
	... // list of fields
	<variable 
}*/

package main
import "fmt"

type Student struct {
	name string
	rollNo int
	marks []int
	grades map[string]int
}

	func main() {
		// var s Student
		// fmt.Printf("%+v", s)
        // output: {name: rollNo:0 marks:[] grades:map[]}
	
		// st := new(Student)
		// fmt.Printf("%+v"
		// , st)
		// output: &{name: rollNo:0 marks:[] grades:map[]}

		st := Student{
			name: "Vivek Vishal",
			rollNo: 2140124,
			marks: []int{43,45},
			grades: map[string]int{"Chemistry":89, "Maths":0},
		}
			fmt.Printf("%+v", st) 
	}
