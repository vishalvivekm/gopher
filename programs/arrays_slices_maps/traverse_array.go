package main
import "fmt"

func main() {
        arr := [5]bool{true, true, true, true}
        fmt.Println(cap(arr), " ", len(arr))
		for i := 0; i < len(arr); i++ {
                if arr[i] {
                        fmt.Println(i)
                }
        }
}
// output
// 5 5
// 0
// 1
// 2
// 3
