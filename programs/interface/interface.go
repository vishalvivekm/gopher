// the output

package main
import (
"fmt"
"log"
)

type Student interface {
        getPercentage() int
        getName()
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

func printPercentage(s Student) {
        fmt.Println(s.getPercentage())
}

func main() {
        grades := []int{90, 75, 80}
        u, err := Undergrad{"Ross", grades}
		  printPercentage(u)
}

