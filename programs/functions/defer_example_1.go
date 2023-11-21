// cleaning up resources, like open files or database connections is
// important when our program is finished working with these resources, so as to avoid memory leaks and also allow other programs or apps to open and access these resources.
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Create("text.txt") // create and open 'text.txt' in read-and-write mode
    if err != nil {
        log.Fatal(err) // exit the program if there's an unexpected error
    }
    defer file.Close() // close the file before exiting the program ( i.e. when the main() finishes, even if there's any error)

    if _, err := fmt.Fprintln(file, "Hello World!"); err != nil {
        log.Fatal(err)
    }
}