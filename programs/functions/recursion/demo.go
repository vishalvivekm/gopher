package main
import "fmt"

func main() {
var num int
fmt.Scanf("%d", &num)
fmt.Println(Factorial(num))
}
func Factorial(num int) int {
if ( num == 0 ) {
return 1 }
return num * Factorial(num-1) ;
}
