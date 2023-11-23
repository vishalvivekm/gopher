package main

import "fmt"

func main() {
        arr := [10]string{"a", "b", "c"}
        hashmap := make(map[string]int)
        my_slice := arr[:]
        fmt.Println(len(my_slice))
        fmt.Println(cap(my_slice))
        fmt.Println(len(hashmap))
}
// output: ?
package main

import "fmt"

func main() {
        arr := [10]string{"a", "b", "c"}
        hashmap := make(map[string]int)
        my_slice := arr[:]
		fmt.Println(arr)
		fmt.Println(my_slice)
        fmt.Println(len(my_slice))
        fmt.Println(cap(my_slice))
        fmt.Println(len(hashmap))
}
i// 
[a b c       ]
[a b c       ]
10
10
0

