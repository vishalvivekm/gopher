// implement the io.Closer interface
/*
type Closer interface {
Close() error
}
*/
package main

import (
	"fmt"
	"io"
)

type FunnyBox struct {
}

func (fb *FunnyBox) Close() error {
	fmt.Println("Bey!")
	return nil
}

// nolint: gosimple // don't delete this comment!
func main() {
	var c io.Closer
	c = &FunnyBox{}
	c.Close()
}
