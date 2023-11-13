package main

import "fmt"

func main() {
        arr := [5]string{"a", "b", "c", "d", "e"}
        slice := arr[:4]
        fmt.Println(arr)
        fmt.Println(slice)
        slice[1] = "x"
        fmt.Println(arr)
        fmt.Println(slice)

}
//output:
