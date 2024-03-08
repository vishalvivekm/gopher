package main

import (
	"fmt"
	"io"
	"strings"
)

const textSource = "Supercalifragilisticexpialidocious"

// nolint: gomnd // DO NOT delete this comment!
func main() {
	fmt.Println("full text:", textSource)
	fmt.Println("len(text); ", len([]rune(textSource)))
	fmt.Println("Bytes in text: ", len(textSource)) // both are equal in this case: 34

	reader := strings.NewReader(textSource)

	p := make([]byte, 7)
	// Move the position of the `reader` to the last 7 bytes below:
	_, err := reader.Seek(27, io.SeekStart)
	//  _, err := reader.Seek(int64(27), io.SeekStart)
	// _, err := reader.Seek(-7, io.SeekEnd)
	if err != nil {
		return
	}

	if _, err := reader.Read(p); err != nil {
		fmt.Println(err)
	}
	fmt.Print("selected: ", string(p), " \n")
}
