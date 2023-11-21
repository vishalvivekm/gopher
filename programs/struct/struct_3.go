// All primitive/basic types (int and its variants, float and its variants, boolean, string, array, and struct) in Go
//  are passed by value.
package main
import "fmt"
type Movie struct {
	name   string
	rating float32
}

func getMovie(s string, r float32) (m Movie) {
	m = Movie{
			name:   s,
			rating: r,
	}
	return
}

func increaseRating(m *Movie) {
	m.rating += 1.0
}

func main() {
	mov := getMovie("xyz", 2.0)
	increaseRating(mov) // sure? this is correct? nope &mov
	fmt.Printf("%+v", mov)
}
}
// 

/*
24: cannot use mov (variable of type Movie) as *Movie value in argument to increaseRating
fix: increaseRating(&mov)
and then 
output:{name:xyz rating:3}
*/