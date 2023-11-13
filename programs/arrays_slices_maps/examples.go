package main

import "fmt"

func main() {

   arr1 := [5]bool{true, true}
        fmt.Println(arr1)
//ans : [ true true false false false]

  arr3 := [10]int{10, 20, 30, 50}
        fmt.Println(arr3)
        fmt.Println(len(arr3))
/*ans: 
[10 20 30 50 0 0 0 0 0 0]
10
*/

 arr4 := [5]string{"a", "b", "c"}
        for index, element := range arr4 {
                fmt.Println(index, "->", element)
        }
/* ans:

0 -> a
1 -> b
2 -> c
3 -> 
4 ->

*/

        arr5 := [5][2]string{{"a"}, {"b"}, {"c"}}
        fmt.Println(arr5)
		fmt.Println(arr5[0][0])
        fmt.Println(arr5[1][1])
        fmt.Println(arr5[2][0])
/*
ans:
[[a ] [b ] [c ] [ ] [ ]]
a

c

*/
}
