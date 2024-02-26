package main

import "fmt"

func printBoth(n int) {
  if n == 0 {
  return
}

 fmt.Println(n)
 printBoth(n-1)
 fmt.Println(n)
}
func main() {
 
 printBoth(5)

}
