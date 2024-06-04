// golang uses implicit break statement for each case, unlike java or c++, So first true case block is executed and then the program exits out of switch block ( unless there's any fallthrough keyword) 

func main() {
var a, b int = 10, 20
switch {
case a+b == 30:
fmt.Println("equal to 30")
case a+b <= 30:
fmt.Println("less than or equal to 30")
default:
fmt.Println("greater than 30")
}
}
//output: equal to 30
