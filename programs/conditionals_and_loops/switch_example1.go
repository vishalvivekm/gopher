package main

import "fmt"

func main() {
        day := "wednesday"
        switch day {
        case "monday":
                fmt.Println("monday")
        case "tuesday":
                fmt.Println("tuesday")
        case "wednesday":
                fmt.Println("wednesday")
                fallthrough
        case "thursday":
                fmt.Println("thursday")
                fallthrough
        case "friday":
                fmt.Println("friday")
        case "saturday", "sunday":
                fmt.Println("weekend")
        default:
                fmt.Println("default")
        }

}
// output 
// wednesday
// thursday
//friday 
