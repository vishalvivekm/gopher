package main

import "fmt"

func main() {

	fmt.Println("Ahh Help")
	fmt.Println("Ahh Help")
	//     fmt.Printf("%[1]s%[1]s", "Ahh Help\n")

}

/*
package main

import "fmt"

func main() {
    // Write your code here
    fmt.Print("First line\n\nThird line\n\nFifth line")
package main

import (
  "fmt"
  "strings"
)

func customPrint(what []string, postfix string, spacing string) {
    for i, s := range what {
        what[i] = fmt.Sprintf("%s %s", s, postfix)
    }
    fmt.Println(strings.Join(what, spacing))
}

func main() {
    var seq = []string{"First", "Third", "Fifth"}
    const postfix = "line"
    const spacing = "\n\n"
    customPrint(seq, postfix, spacing)
}

}


*/
