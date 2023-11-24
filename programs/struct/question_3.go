package main

import "fmt"

type Employee struct {
	eid int
	id  int
}

func main() {
	employees := make([]Employee, 5)
	// fmt.Println(employees) // [{0 0} {0 0} {0 0} {0 0} {0 0}]

	for i := range employees {
		employees[i] = Employee{i, i + 10}
		fmt.Println(employees[i])
	}

}

/*
package main

import "fmt"

type Employee struct {
	eid int
	id int
}
func (e Employee) get_id() int {
	return e.eid + 10;
}
func main() {
	employees := make([]Employee, 5)
	fmt.Println(employees)

	// for i := range employees {
	// 	employees[i] = Employee{i, i+10}
	// 	fmt.Printf("%T ", employees[i])
	// 	fmt.Println(employees[i])
	// }

	for i := range employees {
		employees[i] = Employee{eid: i}
		employees[i].id = employees[i].get_id()
		fmt.Printf("%+v\n", employees[i])
}
}
*/

// error: # command-line-arguments
//./interface.go:31:25: cannot use u (variable of type Undergrad) as Student value in argument to printPercentage: Undergrad does not implement Student (missing method getName)

// every data - type ( struct here) must implement all the functions defined inside the interface

/*
type Student interface {
        getPercentage() int
        getName() string
}

type Undergrad struct {
        name   string
        grades []int
}

func (u Undergrad) getPercentage() int {
        sum := 0
        for _, v := range u.grades {
            sum += v
        }
        return sum / len(u.grades)
}
func (u Undergrad) getName() string {
        return u.name
}

func printData(s Student) {
        fmt.Println(s.getName())
        fmt.Println(s.getPercentage())
}

func main() {
        grades := []int{90, 75, 80}
        u := Undergrad{"Ross", grades}
        printData(u)
}
*/
