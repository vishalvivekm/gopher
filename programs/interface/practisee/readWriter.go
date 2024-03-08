package main

import (
	"fmt"
	"io"
	"log"
)

const Size = 1024

type WisdomBox struct {
	quotes string
}

func (wb *WisdomBox) Write(p []byte) (n int, err error) {
	wb.quotes += string(p)
	return len(p), nil
}

func (wb *WisdomBox) Read(p []byte) (n int, err error) {
	// read data to the p slice
	// copy(p, []byte(wb.quotes)) // func copy(dst []Type, src []Type) int
	return copy(p, []byte(wb.quotes)), nil // return copy(p, wb.quotes), nil
	//return len(p), nil
}

/*
func (wb *WisdomBox) Read(p []byte) (n int, err error) {
    read := 0
    for ; read < len(p) && read < len(wb.quotes); read++ {
        p[read] = wb.quotes[read]
    }

	return read, nil
}
// or,

func (wb *WisdomBox) Read(p []byte) (n int, err error) {
	// read data to the p slice
	n = copy(p, wb.quotes)
	if n == 0 {
		return 0, io.EOF
	}
	wb.quotes = wb.quotes[n:]
	return n, nil
}


*/

func main() {
	var box io.ReadWriter = &WisdomBox{}
	fmt.Fprintln(box, "Never underestimate anyone") // func Fprintln(w io.Writer, a ...any) (n int, err error), box is a io.Writer
	fmt.Fprintln(box, "Home is where the heart is")
	fmt.Fprintln(box, "You can be whoever you want to be")

	p := make([]byte, Size)
	n, err := box.Read(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(p[:n-1]))

	// fmt.Println(box)
	/*
		&{Never underestimate anyone
		Home is where the heart is
		You can be whoever you want to be
		}

	*/
}

/*
type ReaderWriter interface {
Reader
Writer
}
*/
