// set of methods that are available to a data type.
package main

import "fmt"

type Student struct {
	name   string
	grades []int
}

func (s *Student) displayName() {
	fmt.Println(s.name)
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
	st.displayName()
	fmt.Printf("%.2f\n", st.calculatePercentage())
}
