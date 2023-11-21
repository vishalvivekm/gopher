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

func main() {
	fmt.Printf("%+v", getMovie("xyz", 3.5))
}
//