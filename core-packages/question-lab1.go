/*
Write a code in main.go file to read the string in a buffer of size 5 at a time and print the output in each iteration.

A Go file is located at /root/code/string directory.

package main

import (
    "strings"
)

func main() {
    reader := strings.NewReader("Let us catch up over a cup of coffee")
    // your code goes here

}




Expected Output:

[76 101 116 32 117] <nil>
[115 32 99 97 116] <nil>
[99 104 32 117 112] <nil>
[32 111 118 101 114] <nil>
[32 97 32 99 117] <nil>
[112 32 111 102 32] <nil>
[99 111 102 102 101] <nil>
[101] <nil>
[] EOF
*/
// solution:
/*
package main

import (
	"strings"
	"fmt"
)

func main() {
	reader := strings.NewReader("Let us catch up over a cup of coffee")
	// your code goes here

	buffer := make([]byte, 5)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(buffer[:n], err)
	}

}

*/
/* explaination:



Add the following Go code in the main.go file inside the directory called string.

package main

import (
    "fmt"
    "strings"
)

func main() {
    reader := strings.NewReader("Let us catch up over a cup of coffee")
    buf := make([]byte, 5)
    for {
        n, err := reader.Read(buf)
        fmt.Println(buf[:n], err)
        if err != nil {
            break
        }
    }

}


Then, run the following go run command as follows: -

go run main.go


    NOTE: - Make sure to run the command in the correct directory.


Explanation:

The code creates a byte slice of length 5, which will be used as a buffer.

The code then enters an infinite loop which repeatedly calls the Read method on the reader object, passing the buffer as an argument.

The Read method reads from the input string and fills the buffer with up to 5 bytes of data.

It returns the number of bytes read, n, and an error, err.

The code then prints the bytes read from the buffer and the error.

If the error is not nil, the loop breaks.

This will continue until the entire input string has been read and the error returned is io.EOF, indicating the end of the input.





*/

/*
Solve the syntax errors in main_test.go file, so go test gives a PASS.


A Go file is located at the /root/code/test directory.

package main

func HelloWorld() string {
    return "hello world"
}



A Go test file is located at the /root/code/test directory.

package main

func TestHelloWorld() {

    if actual != "hello world" {
        t.Errorf("expected 'hello world', got '%s'", actual)
    }
}



    NOTE: - You don't need to run the main.go file.
*/