
package main

import "fmt"

func main() {
        var a, b string = "foo", "bar"
        if a+b == "foo" {
                fmt.Println("foo")
        } else if a+b == "foobar" {
                fmt.Println("bar")
        } else if a+b == "foobar" {
                fmt.Println("foobar")
        } else {
                fmt.Println("None matched")
        }
        fmt.Println("thank you!")

}
// here only first else-if having the true condition will be executed and then the it will come out of the if-elseif-else block



