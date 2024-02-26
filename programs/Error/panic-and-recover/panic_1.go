package main

import "fmt"

func panicHandler() {
    switch recover() {
    case "negative":
        fmt.Println("Number must be positive")
    case "zero":
        fmt.Println("Number must have a non-zero value")
    case "bounds":
        fmt.Println("Number goes out of bounds")
    case "non-prime":
        fmt.Println("Number is non-prime")
    case nil:
        fmt.Println("Number is suitable!")
    }
}

// Tip: recover() only works with defer
func main() {
    defer panicHandler()
    var number int16

    fmt.Scan(&number)

    if number < 0 {
        panic("negative")
    }

    if number == 0 {
        panic("zero")
    }

    if number > 100 { // nolint: gomnd // DO NOT delete this comment!
        panic("bounds")
    }

    var i int16
    for i = 2; i < number; i++ {
        if number%i == 0 {
            panic("non-prime")
        }
    }
}

