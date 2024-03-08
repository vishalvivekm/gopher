// func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
/*
 method: ReadAt() from io package. ReadAt() implements the io.ReaderAt interface.

Take a look at the ReadAt() method definition below, where off represents the offset (which you may recall from the Seek() method):

func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
*/

package main

import (
	"fmt"
	"log"
	"strings"
)

const quote = "Knowledge is better than money"

// nolint: gomnd // don't delete this comment!
func main() {
	reader := strings.NewReader(quote)

	// Write the missing code below so that the program will output "money"
	buf := make([]byte, 5)
	_, err := reader.ReadAt(buf, 25)
	// checkout the alternate approach below:
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
}

/*
    buf := make([]byte, 5)
	_, err := reader.ReadAt(buf, 25)

	======================================

	offset = 5
	buf := make([]byte, offset)
	_, err := reader.ReadAt(buf, int64(len(quote) - offset ))

	=======================================

    buf := make([]byte, len("money"))
    _, err := reader.ReadAt(buf, int64(len(quote)-len("money")))

*/
