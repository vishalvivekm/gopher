package main

import "fmt"

func calculate(a int, b int) []float64 {
	sum := float64(a) + float64(b)
	diff := float64(a) - float64(b)
	mult:= float64(a) * float64(b)
	quot := float64(a) / float64(b)

	results :=[]float64{sum, diff, mult, quot}
	/* result := make([]float64, 0)
	return append(results, float64(a+b), float64(a-b), float64(a*b), float64(a/b))

	*/
	return results
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}


/*
package main

import "fmt"

func calculate(a int, b int) (results []float64) {
    results = append(results, float64(a+b), float64(a-b), float64(a*b), float64(a/b))
    return results
}

func main() {
    fmt.Println(calculate(20, 10))
    fmt.Println(calculate(700, 70))
}
*/
