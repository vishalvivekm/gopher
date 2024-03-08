//strings : package with strings/string slice manipulation methods, bytes: package with byte slice manipulation methods

package main

import (
	"bytes" 
	"fmt"
	"log"
)

func main() {
	var b bytes.Buffer // a struct implementing both io.Reader and io.Writer 
	_, err := fmt.Fprintln(&b, "Vivek Vishal")

	if err != nil {
		return
	}

	// fmt.Println(b.String())

	buf := make([]byte, 1024)
	if _, err := b.Read(buf); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))

	v := []byte("Vivek Vishal") // typecaste a string into slice of byte
	// c := 'v' character literals are taken as rune type (int32)
	// fmt.Println(c) // 118
	fmt.Printf("%v -> %v\n", v, string(v))
	b.Write([]byte("Golang isn't that hard to learn, is it ?."))
    fmt.Println(b.String())

}

// Buffer implements both io.Reader as well as io.Writer as it has both the methods : Read() and Write()

/**
package Bytes

type Buffer struct {
buf []byte
..
..
}
func (b *Buffer) Write(p []byte) (n int, err error) {
// implementation
}

var I io.Writer
var b bytes.Buffer // doesn't implement io.Writer
var c *bytes.Buffer // implements io.Writer

I = b // wrong
I = c // correct
*/
