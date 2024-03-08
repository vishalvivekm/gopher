/*
type Reader interface {
    Read(p []byte) (n int, err error)
}
*/

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func show(r io.Reader, chunkSize int) (string, error) {
	buf := make([]byte, chunkSize)
	n, err := r.Read(buf)
	if err != nil {
		//log.Fatal(err) //
		return "", err

	}
	return fmt.Sprintf("Read %d bytes from '%s'", n, string(buf)), nil
}
func main() {
	chunksize := 1024
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// close file
	defer func(file *os.File) {
		err := file.Close() // Close() returns an error in case the file is already closed
		if err != nil {
			log.Fatal("already closed")
		}
	}(file)

	//	fmt.Println(show(file, chunksize))
	fileText, err := show(file, chunksize)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileText)

	name := strings.NewReader("Vivek Vishal")
	fmt.Println(show(name, chunksize))

	// file; os.File and name; strings.Reader both return Reader which implements io.Reader i.e. they have this method : Read(p []byte) (n int, err error)
}
